package attack_message_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/hids/request"
	hids_server "github.com/1uLang/zhiannet-api/hids/server"
	hids_agent_server "github.com/1uLang/zhiannet-api/hids/server/agent"
	hids_risk_server "github.com/1uLang/zhiannet-api/hids/server/risk"
	"time"
)

type hids struct{}

//检测所有agent ip 是否异常
func (hids) AttackCheck() error {
	info, err := hids_server.GetHideInfo()
	if err != nil {
		return err
	}
	err = hids_server.SetUrl(info.Addr)
	if err != nil {
		return err
	}
	err = hids_server.SetAPIKeys(&request.APIKeys{info.AppId, info.Secret})
	if err != nil {
		return err
	}

	resp, err := hids_agent_server.ListAll()
	if err != nil {
		return err
	}
	var ips []string
	ipsInfo := map[string]struct {
		users  []uint64
		admins []uint64
	}{}
	for _, item := range resp.List {
		serverIp := item["serverIp"].(string)

		//查询该ip所属用户列表
		users, admins, err := getAgentIpUsers(serverIp)
		if err != nil {
			return err
		}
		if len(users) > 0 || len(admins) > 0 {
			ips = append(ips, serverIp)
			ipsInfo[serverIp] = struct {
				users  []uint64
				admins []uint64
			}{users: users, admins: admins}
		}
	}
	riskIps, err := hids_risk_server.CheckRiskAttack(ips)
	if err != nil {
		return err
	}
	invadeIps, err := hids_risk_server.CheckInvadeAttack(ips)
	if err != nil {
		return err
	}
	for _, ip := range riskIps {
		_ = checkRiskAndCreateMessage(ip, ipsInfo[ip].users, ipsInfo[ip].admins, "漏洞风险")
	}
	for _, ip := range invadeIps {
		_ = checkRiskAndCreateMessage(ip, ipsInfo[ip].users, ipsInfo[ip].admins, "入侵威胁")
	}
	return nil
}
func checkRiskAndCreateMessage(ip string, users, admins []uint64, msg string) error {

	for _, user := range users {
		edge_messages.Add(&edge_messages.Edgemessages{
			Level:     "warning",
			Subject:   "主机异常",
			Body:      fmt.Sprintf("您的主机[%s]存在%s，请前往主机防护下相应菜单查阅详情。", ip, msg),
			Type:      "HostRisk",
			Params:    "{}",
			Userid:    uint(user),
			Createdat: uint64(time.Now().Unix()),
			Day:       time.Now().Format("20060102"),
			Hash:      "",
			Role:      "user",
		})
	}

	for _, admin := range admins {
		edge_messages.Add(&edge_messages.Edgemessages{
			Level:     "warning",
			Subject:   "主机异常",
			Body:      fmt.Sprintf("您的主机[%s]存在'%s'，请前往主机防护下相应菜单查阅详情。", ip, msg),
			Type:      "HostRisk",
			Params:    "{}",
			Adminid:   uint(admin),
			Createdat: uint64(time.Now().Unix()),
			Day:       time.Now().Format("20060102"),
			Hash:      "",
			Role:      "admin",
		})
	}

	return nil
}

//获取改agent ip 所属用户列表
func getAgentIpUsers(ip string) (users []uint64, admins []uint64, err error) {
	return hids_agent_server.GetUserListByAgentIP(ip)
}
