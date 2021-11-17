package access_gateway

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	ListReq struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`

		UserId      int64 `json:"-"`
		AdminUserId int64 `json:"-"`
	}
	ListRes struct {
		GatewayInfo
	}
	GatewayInfo struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		AccountType string `json:"accountType"`
		Created     string `json:"created"`
		Connected   bool   `json:"connected"`
		Message     string `json:"message"`
	}
	CreateReq struct {
		Name        string `json:"name"`
		IP          string `json:"ip"`
		Port        int    `json:"port"`
		Localhost   string `json:"localhost"` // 隧道映射到本地的地址
		AccountType string `json:"accountType"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		PrivateKey  string `json:"privateKey"`
		Passphrase  string `json:"passphrase"`

		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
	UpdateReq struct {
		Id string `json:"id"`
		CreateReq
	}
	AuthorizeReq struct {
		Id string `json:"id"`

		UserId       uint64   `json:"-"`
		AdminUserId  uint64   `json:"-"`
		UserIds      []uint64 `json:"-"`
		AdminUserIds []uint64 `json:"-"`
	}
)

func (this CreateReq) check() error {
	errMsg := ""

	if this.Name == "" {
		errMsg = "网关名称不能为空"
		goto ERR
	}
	if this.IP == "" {
		errMsg = "网关IP不能为空"
		goto ERR
	}
	if this.Port <= 0 {
		errMsg = "网关端口无效"
		goto ERR
	}
	if this.AccountType == "password" {
		if this.Username == "" && this.Password == "" {
			errMsg = "网关授权账号/密码不能为空"
			goto ERR
		}
	} else if this.AccountType == "private-key" {
		if this.Username == "" && this.Password == "" {
			errMsg = "网关授权账号/私钥不能为空"
			goto ERR
		}
	} else {
		errMsg = "网关账户类型错误"
		goto ERR
	}
	this.Localhost = "next-terminal"

ERR:
	if errMsg != "" {
		return fmt.Errorf("参数错误：%s", errMsg)
	}
	return nil
}

func (this UpdateReq) check() error {
	if this.Id == "" {
		return fmt.Errorf("网关id不能为空")
	}
	return this.CreateReq.check()
}

func (this AuthorizeReq) check() error {

	if this.Id == "" {
		return fmt.Errorf("网关id不能为空")
	}
	if this.UserId == 0 && this.AdminUserId == 0 {
		return fmt.Errorf("用户不能为空")
	}

	// 判断当前用户是否用户授权权限
	var count int64
	model := db_model.MysqlConn.Model(&nextTerminalAccessGateway{}).Where("is_delete=?", 0).Where("auth = 0")
	if this.UserId != 0 {
		model = model.Where("user_id = ?", this.UserId)
	}
	if this.AdminUserId != 0 {
		model = model.Where("admin_user_id = ?", this.AdminUserId)
	}
	err := model.Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("无权限")
	}
	return nil
}
