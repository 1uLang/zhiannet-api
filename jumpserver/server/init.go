package server

import (
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type Request struct {
	Users     users       //用户接口
	Assets    assets      //资产接口
	AdminUser admin_users //管理用户接口
}

var req *Request

//NewServerRequest 初始化 服务器url 当前用户的username,password
func NewServerRequest(url, username, password string) (*Request, error) {

	_ = request.InitServerUrl(url)
	err := request.InitToken(username, password)
	if err != nil {
		return nil, err
	}
	req := &Request{}
	req.Users.req, err = request.NewRequest()
	if err != nil {
		return nil, err
	}
	req.Assets.req = req.Users.req
	req.AdminUser.req = req.Users.req
	return req, err
}
func GetFortCloud() (resp *model.JumpserverResp, err error) {
	return model.GetJumpserverInfo()
}
