package attack_check_server

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_db_nodes"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/server/edge_ssl_policies_server"
	"github.com/1uLang/zhiannet-api/common/util"
	"time"
)

const ipListVersion = "IP_LIST_VERSION"
const timeout = 5 * 60 //s 定时任务间隔时间
var tab_prefixx = "edgeHTTPAccessLogs_"

type httpLog struct {
	ID               uint64 `gorm:"column:id" json:"id" form:"id"`
	Status           int    `gorm:"column:status" json:"status" form:"status"`
	FirewallPolicyId uint64 `gorm:"column:firewallPolicyId" json:"firewallPolicyId" form:"firewallPolicyId"`
	RemoteAddr       string `gorm:"column:remoteAddr" json:"remoteAddr" form:"remoteAddr"`
}

type version struct {
	Id      uint64 `gorm:"column:id" json:"id" form:"id"`
	Key     string `gorm:"column:key" json:"key" form:"key"`
	Version uint64 `gorm:"column:version" json:"version" form:"version"`
}

type waf struct{}

//version 版本号
func increaseVersion() (version, error) {
	//ipListVersion
	valus := version{}
	err := model.MysqlConn.Debug().Table("edgeSysLockers").Where("`key`=?", ipListVersion).Find(&valus).Error
	return valus, err
}
func increaseAdd(ver version) error {
	//ipListVersion
	return model.MysqlConn.Debug().Table("edgeSysLockers").Save(&ver).Error
}

var firewallPolicyMaps = map[uint64]map[string]interface{}{}

func (this *httpLog) CheckAttach() (uint64, bool, error) {

	if code, isExist := firewallPolicyMaps[this.FirewallPolicyId]; isExist {
		return code["id"].(uint64), code["code"].(int) == this.Status, nil
	}
	policy := struct {
		Inbound      []byte `gorm:"column:inbound" json:"inbound" form:"inbound"`
		BlockOptions []byte `gorm:"column:blockOptions" json:"blockOptions" form:"blockOptions"`
	}{}
	opt := struct {
		Url        string `json:"url"`
		Body       string `json:"body"`
		Timeout    int    `json:"timeout"`
		StatusCode int    `json:"statusCode"`
	}{}
	black := struct {
		BlackListRef struct {
			IsOn   bool   `json:"isOn"`
			ListId uint64 `json:"listId"`
		} `json:"blackListRef"`
	}{}
	err := model.MysqlConn.Table("edgeHTTPFirewallPolicies").Where("id=?", this.FirewallPolicyId).Find(&policy).Error
	if err != nil {
		return 0, false, err
	}
	err = json.Unmarshal(policy.BlockOptions, &opt)
	if err != nil {
		return 0, false, err
	}
	err = json.Unmarshal(policy.Inbound, &black)
	if err != nil {
		return 0, false, err
	}
	firewallPolicyMaps[this.FirewallPolicyId] = map[string]interface{}{"id": black.BlackListRef.ListId, "code": opt.StatusCode}
	return black.BlackListRef.ListId, this.Status == opt.StatusCode, nil
}

type httpsProtocolConfig struct {
	SslPolicyRef struct {
		IsOn        bool   `json:"isOn"`
		SslPolicyId uint64 `json:"sslPolicyId"`
	} `json:"sslPolicyRef"`
}

func checkAndUpdateHttpsConfig(config []byte) {
	ref := httpsProtocolConfig{}
	_ = json.Unmarshal(config, &ref)
	if ref.SslPolicyRef.SslPolicyId > 0 {
		err := edge_ssl_policies_server.CheckAndUpdate(ref.SslPolicyRef.SslPolicyId, []string{"TLS 1.1", "TLS 1.0"}, "TLS 1.2")
		if err != nil {
			fmt.Println("update https ssl info error:", err)
		}
	}
}

//waf 入侵检测
func (waf) WAFAttackCheck() error {

	//读取db nodes
	db, err := edge_db_nodes.NewConn()
	if err != nil {
		return err
	}
	logs := make([]httpLog, 0)
	//edgeHTTPFirewallPolicies 判断其拦截状态码
	tabname := fmt.Sprintf("%s%s", tab_prefixx, time.Now().Format("20060102"))
	err = db.Table(tabname).
		Where("createdAt>=?", time.Now().Unix()-timeout).
		Where("status!=200").Find(&logs).Error
	if err != nil {
		return err
	}
	//读取当天日志，查询非200整除放回状态码 -> 访问源ip
	ipWafs := map[string]httpLog{}
	for _, v := range logs {
		ipWafs[fmt.Sprintf("%v_%v", v.Status, v.RemoteAddr)] = v
	}
	for _, v := range ipWafs {
		listId, ok, err := v.CheckAttach()
		if err != nil {
			return err
		}
		//入侵拦截
		//找到 所属waf策略[firewallPolicyId] -> 策略自动添加该ip黑名单
		if ok {
			err = addBlackIP(listId, v.RemoteAddr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func addBlackIP(id uint64, addr string) (err error) {

	ipitem := struct {
		Id         uint64 `gorm:"column:id" json:"id" form:"id"`
		ListId     uint64 `gorm:"column:listId" json:"listId" form:"listId"`
		IPFrom     string `gorm:"column:ipFrom" json:"ipFrom" form:"ipFrom"`
		IPTo       string `gorm:"column:ipTo" json:"ipTo" form:"ipTo"`
		ExpiredAt  int64  `gorm:"column:expiredAt" json:"expiredAt" form:"expiredAt"`
		CreateAt   int64  `gorm:"column:createdAt" json:"createdAt" form:"createdAt"`
		IpFromLong uint64 `gorm:"column:ipFromLong" json:"ipFromLong" form:"ipFromLong"`
		Reason     string `gorm:"column:reason" json:"reason" form:"reason"`
		Type       string `gorm:"column:type" json:"type" form:"type"`
		EventLevel string `gorm:"column:eventLevel" json:"eventLevel" form:"eventLevel"`
		Version    uint64 `gorm:"column:version" json:"version" form:"version"`
		State      uint8  `gorm:"column:state" json:"state" form:"state"`
	}{}

	//判断是否已加入至黑名单
	var count int64
	err = model.MysqlConn.Table("edgeIPItems").Where("ipFrom=?", addr).Where("state=1").Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		return nil
	}

	if util.IsIPv4(addr) {
		ipitem.Type = "ipv4"
	} else if util.IsIPv6(addr) {
		ipitem.Type = "ipv6"
	}
	ipitem.IPFrom = addr
	ipitem.IpFromLong = util.IP2Long(addr)
	ipitem.EventLevel = "critical"
	ipitem.Reason = "入侵检测自动加入黑名单"
	ipitem.ListId = id
	//启用
	ipitem.State = 1
	ver, err := increaseVersion()
	if err != nil {
		return err
	}
	ver.Version++
	ipitem.Version = ver.Version
	ipitem.CreateAt = time.Now().Unix()

	//写入数据库
	err = model.MysqlConn.Table("edgeIPItems").Create(&ipitem).Error
	if err != nil {
		return err
	}
	_ = increaseAdd(ver)

	edge_messages.Add(&edge_messages.Edgemessages{
		Level:     "error",
		Subject:   "入侵检测",
		Body:      fmt.Sprintf("WAF入侵检测[%s],已自动添加至WAF IP黑名单", addr),
		Type:      "IntrusionDetection",
		Params:    "{}",
		Createdat: uint64(time.Now().Unix()),
		Day:       time.Now().Format("20060102"),
		Hash:      "",
		Role:      "admin",
	})
	return nil
}
