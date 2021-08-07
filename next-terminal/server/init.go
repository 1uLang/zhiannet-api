package server

import (
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

type Request struct {
	Assets  asset   //资产接口
	Session session //会话管理
	Cert 	cert 	//授权凭证
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
	req.Assets.req, err = request.NewRequest()
	if err != nil {
		return nil, err
	}
	req.Session.req = req.Assets.req
	req.Cert.req = req.Assets.req
	return req, err
}
func GetFortCloud() (resp *model.NextTerminalResp, err error) {
	return model.GetNextTerminalInfo()
}
