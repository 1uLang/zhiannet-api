package users

import (
	"fmt"
)

type ListReq struct {
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}
type CreateReq struct {
	Email              string   `json:"email"`                //邮箱
	mfaLevel           int      `json:"mfa_level"`            //双因子认证 0
	Name               string   `json:"name"`                 //用户名称
	needUpdatePassword bool     `json:"need_update_password"` //是否需要重置密码 fasle
	OrgRoles           []string `json:"org_roles"`            //用户组 User
	Password           string   `json:"password"`             //密码
	passwordStrategy   string   `json:"password_strategy"`    //
	role               string   `json:"role"`
	source             string   `json:"source"` //用户来源 local
	Username           string   `json:"username"`
	DateExpired        string   `json:"date_expired"`
}
type Update struct {
	ID string `json:"id"`
	CreateReq
}

func (this *CreateReq) check() error {

	this.source = "local"
	this.role = "admin"
	this.passwordStrategy = "custom"
	this.needUpdatePassword = false
	this.mfaLevel = 0

	if this.Name == "" {
		return fmt.Errorf("请输入名称")
	}
	if this.Username == "" {
		return fmt.Errorf("请输入用户名")
	}
	if this.Email == "" {
		return fmt.Errorf("请输入邮箱")
	}
	if this.Password == "" {
		return fmt.Errorf("请输入密码")
	}
	return nil
}
