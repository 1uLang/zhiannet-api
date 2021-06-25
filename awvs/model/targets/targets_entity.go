package targets

import "fmt"

type ListReq struct {
	Limit       int    `json:"l,omitempty"` //限制条数
	C           int    `json:"c,omitempty"` //偏移量
	Query       string `json:"q,omitempty"` //筛选器
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}
type UpdateReq struct {
	Criticality int    `json:"criticality"` //危险程度，30,20,10,0 默认 10
	Description string `json:"description"` //描述
}
type AddReq struct {
	Address     string `json:"address"` //地址
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
	UpdateReq
}

func (this *AddReq) Check() (bool, error) {

	if this.Address == "" {
		return false, fmt.Errorf("目标地址不能为空")
	}
	return this.UpdateReq.Check()
}
func (this *UpdateReq) Check() (bool, error) {

	if this.Criticality != 30 && this.Criticality != 20 && this.Criticality != 10 && this.Criticality != 0 {
		return false, fmt.Errorf("目标危险程度无效")
	}
	return true, nil
}

type SetLoginReq struct {
	Login struct { //站点预设登录
		Kind        string `json:"kind"` //启用:automatic; 不启用:none(默认); 使用登录序列:sequence
		Credentials struct {
			Enabled  bool   `json:"enabled"`  //是否生效
			Username string `json:"username"` //账号
			Password string `json:"password"` //密码
		} `json:"credentials"` //登录凭证
		SshCredentials struct {
			Kind string `json:"kind"` //启用:automatic; 不启用:none(默认); 使用登录序列:sequence
		}
	} `json:"login"`
	Sensor bool `json:"sensor"` //传感器
}

func (this *SetLoginReq) Check() (bool, error) {
	if this.Login.Kind == "" || (this.Login.Kind != "automatic" && this.Login.Kind != "none") {
		return false, fmt.Errorf("是否使用登录序列号错误")
	}
	if this.Login.Credentials.Enabled && (this.Login.Credentials.Username == "" || this.Login.Credentials.Password == "") {
		return false, fmt.Errorf("登录账号密码错误")
	}
	return true, nil
}
