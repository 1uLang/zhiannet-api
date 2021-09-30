package edge_node_clusters

import "github.com/1uLang/zhiannet-api/common/model"

type (
	EdgeNodeClusters struct {
		Id                   uint64 `gorm:"column:id" json:"id" form:"id"`                                                       //ID
		AdminId              uint64 `gorm:"column:adminId" json:"adminId" form:"adminId"`                                        //管理员ID
		Userid               uint64 `gorm:"column:userId" json:"userId" form:"userId"`                                           //用户ID
		Name                 string `gorm:"column:name" json:"name" form:"name"`                                                 //名称
		UseAllApiNodes       int    `gorm:"column:useAllAPINodes" json:"useAllAPINodes" form:"useAllAPINodes"`                   //是否使用所有API节点
		ApiNodes             string `gorm:"column:apiNodes" json:"apiNodes" form:"apiNodes"`                                     //使用的API节点
		InstallDir           string `gorm:"column:installDir" json:"installDir" form:"installDir"`                               //安装目录
		Order                int    `gorm:"column:order" json:"order" form:"order"`                                              //排序
		CreatedAt            uint64 `gorm:"column:createdAt" json:"createdAt" form:"createdAt"`                                  //创建时间
		GrantId              int    `gorm:"column:grantId" json:"grantId" form:"grantId"`                                        //默认认证方式
		State                int    `gorm:"column:state" json:"state" form:"state"`                                              //状态
		AutoRegister         int    `gorm:"column:autoRegister" json:"autoRegister" form:"autoRegister"`                         //是否开启自动注册
		UniqueId             string `gorm:"column:uniqueId" json:"uniqueId" form:"uniqueId"`                                     //唯一ID
		Secret               string `gorm:"column:secret" json:"secret" form:"secret"`                                           //密钥
		HealthCheck          string `gorm:"column:healthCheck" json:"healthCheck" form:"healthCheck"`                            //健康检查
		DnsName              string `gorm:"column:dnsName" json:"dnsName" form:"dnsName"`                                        //DNS名称
		DnsDomainId          int    `gorm:"column:dnsDomainId" json:"dnsDomainId" form:"dnsDomainId"`                            //域名ID
		Dns                  string `gorm:"column:dns" json:"dns" form:"dns"`                                                    //DNS配置
		Toa                  string `gorm:"column:toa" json:"toa" form:"toa"`                                                    //TOA配置
		Cachepolicyid        int    `gorm:"column:cachePolicyId" json:"cachePolicyId" form:"cachePolicyId"`                      //缓存策略ID
		HttpFirewallPolicyId int    `gorm:"column:httpFirewallPolicyId" json:"httpFirewallPolicyId" form:"httpFirewallPolicyId"` //WAF策略ID
		AccessLog            string `gorm:"column:accessLog" json:"accessLog" form:"accessLog"`                                  //访问日志设置
		SystemServices       string `gorm:"column:systemServices" json:"systemServices" form:"systemServices"`                   //系统服务设置
		IsOn                 int    `gorm:"column:isOn" json:"isOn" form:"isOn"`                                                 //是否启用
	}
	ListReq struct {
		Username string   `json:"user_name"`
		PageNum  int      `json:"page_num"`
		PageSize int      `json:"page_size"`
		Ids      []uint64 `json:"ids"`
	}
)

func GetList(req *ListReq) (list []*EdgeNodeClusters, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("edgeNodeClusters")
	if req != nil {
		if req.Username != "" {
			model = model.Where("name LIKE ? ", "%"+req.Username+"%")
		}
		if len(req.Ids) > 0 {
			model = model.Where("id in(?)", req.Ids)
		}
	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	return
}
