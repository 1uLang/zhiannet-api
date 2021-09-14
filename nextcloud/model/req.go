package model

// LoginReq 登陆请求参数
type LoginReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// CreateUserReq 创建用户请求参数
type CreateUserReq struct {
	Userid      string        `json:"userid"`
	Password    string        `json:"password"`
	DisplayName string        `json:"displayName"`
	Email       string        `json:"email"`
	Groups      []string      `json:"groups"`
	Subadmin    []interface{} `json:"subadmin"`
	Quota       string        `json:"quota"`
	Language    string        `json:"language"`
}

// ChangeUPwdReq 修改用户密码参数
type ChangeUPwdReq struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
