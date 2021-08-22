package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	cert_model "github.com/1uLang/zhiannet-api/next-terminal/model/cert"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
	"time"
)

type Request struct {
	Assets  asset   //资产接口
	Session session //会话管理
	Cert    cert    //授权凭证
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
func Check() (bool, uint64, error) {
	info, err := GetFortCloud()
	if err != nil {
		return false, 0, err
	}
	req, err := NewServerRequest(info.Addr, info.Username, info.Password)
	if err != nil {
		return false, info.Id, err
	}
	_, _, err = req.Assets.List(&asset_model.ListReq{UserId: 1})
	if err != nil {
		return false, info.Id, err
	}
	//_, err = user.List(&user.SearchReq{})
	//if err != nil {
	//	return false, info.Id, err
	//}
	return true, info.Id, nil
}

type CheckRequest struct{}

func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("next-terminal-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, _ := Check()
	//fmt.Println(res, id, err)
	if !res {
		conn = 0
		edge_messages.Add(&edge_messages.Edgemessages{
			Level:     "error",
			Subject:   "组件状态异常",
			Body:      "堡垒机状态不可用",
			Type:      "AdminAssembly",
			Params:    "{}",
			Createdat: uint64(time.Now().Unix()),
			Day:       time.Now().Format("20060102"),
			Hash:      "",
			Role:      "admin",
		})
	}
	if id > 0 {
		subassemblynode.UpdateConnState(id, conn)
	}
}

func InitTable() {
	cert_model.InitTable()
	asset_model.InitTable()
}
