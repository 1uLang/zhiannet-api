package model

// LoginReq 登陆请求参数
type LoginReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}
