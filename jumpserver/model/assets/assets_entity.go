package assets

import (
	"fmt"
	"strings"
)

type ListReq struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	NodeId      string `json:"node_id,omitempty"`
	Display     int    `json:"display"`
	UserId      uint64 `json:"user_id,omitempty"`
	AdminUserId uint64 `json:"admin_user_id,omitempty"`
}

type CreateReq struct {
	HostName  string   `json:"hostname"`   //主机名
	IP        string   `json:"ip"`         //IP
	Platform  string   `json:"platform"`   //系统平台
	PublicIp  string   `json:"public_ip"`  //公网ip
	Protocols []string `json:"protocols"`  //协议组
	AdminUser string   `json:"admin_user"` //管理用户
	Nodes     []string `json:"nodes"`      //节点
	Active    bool     `json:"is_active"`  //激活
	Comment   string   `json:"comment"`    //备注

	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}

type UpdateReq struct {
	ID string `json:"id"`
	CreateReq
}
type AuthorizeReq struct {
	Asset string
	Emails []string
}
type DelAuthorizeReq struct {
	Asset string
}

type AuthorizeListReq struct {
	Asset string
}

var protocolsMaps = map[string]bool{
	"ssh":    false,
	"vnc":    false,
	"rdp":    false,
	"telnet": false,
}
var platformMaps = map[string]bool{
	"Linux":       false,
	"Unix":        false,
	"MacOS":       false,
	"BSD":         false,
	"Windows":     false,
	"Windows2016": false,
	"Other":       false,
}

func (this *CreateReq) check() error {

	if this.HostName == "" {
		return fmt.Errorf("请输入主机名")
	}
	if this.IP == "" {
		return fmt.Errorf("请输入主机ip或域名")
	}
	if this.Platform == "" {
		return fmt.Errorf("请选择主机系统平台")
	} else {
		if _, isExist := platformMaps[this.Platform]; !isExist {
			return fmt.Errorf("主机系统平台类型错误")
		}
	}
	for _, v := range this.Protocols {
		protocols := strings.Split(v, "/")
		if len(protocols) != 2 {
			return fmt.Errorf("请设置主机协议组")
		} else {
			if _, isExist := protocolsMaps[protocols[0]]; !isExist {
				return fmt.Errorf("主机协议类型错误")
			}
		}
	}

	return nil
}

func (this *UpdateReq) check() error {

	if this.ID == "" {
		return fmt.Errorf("请输入资产id")
	}
	return this.CreateReq.check()
}
