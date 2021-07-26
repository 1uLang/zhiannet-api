package admin_users

import (
	"fmt"
)

type ListReq struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}

type CreateReq struct {
	Name       string `json:"name"`        //名称
	UserName   string `json:"username"`    //用户名
	Password   string `json:"password"`    //密码
	PrivateKey string `json:"private_key"` //SSH密钥
	Comment    string `json:"comment"`     //备注

	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}

type UpdateReq struct {
	ID string `json:"id"`
	CreateReq
}

func (this *CreateReq) check() error {

	if this.Name == "" {
		return fmt.Errorf("请输入管理用户名称")
	}
	if this.UserName == "" {
		return fmt.Errorf("请输入管理用户账号")
	}
	return nil
}

func (this *UpdateReq) check() error {
	if this.ID == "" {
		return fmt.Errorf("请输入管理用户id")
	}
	return this.CreateReq.check()
}
