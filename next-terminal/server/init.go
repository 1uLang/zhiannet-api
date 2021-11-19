package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	access_gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
	"time"
)

type Request struct {
	Assets  asset   //资产接口
	Session session //会话管理
	Cert    cert    //授权凭证
	GateWay gateway //接入网关
}

//NewServerRequest 初始化 服务器url 当前用户的username,password
func NewServerRequest(url, username, password string) (*Request, error) {

	var err error
	req := &Request{}
	req.Assets.req, err = request.NewRequest(url)
	if err != nil {
		return nil, err
	}
	req.Session.req = req.Assets.req
	req.Cert.req = req.Assets.req
	req.GateWay.req = req.Assets.req
	return req, err
}
func GetFortCloud() (resp *model.NextTerminalResp, err error) {
	return model.GetNextTerminalInfo()
}
func Check() (bool, uint64, int, error) {
	info, err := GetFortCloud()
	if err != nil {
		return false, 0, 0, err
	}
	req, err := NewServerRequest(info.Addr, info.Username, info.Password)
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	_, _, err = req.Assets.List(&asset_model.ListReq{UserId: 1})
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	//_, err = user.List(&user.SearchReq{})
	//if err != nil {
	//	return false, info.Id, err
	//}
	return true, info.Id, info.ConnState, nil
}

type CheckRequest struct{}

func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("next-terminal-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, oldConn, _ := Check()
	//fmt.Println(res, id, err)

	if id > 0 {
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
		if oldConn != conn {
			subassemblynode.UpdateConnState(id, conn)
			if conn == 1 {
				edge_messages.Add(&edge_messages.Edgemessages{
					Level:     "success",
					Subject:   "组件状态恢复正常",
					Body:      "堡垒机恢复可用状态",
					Type:      "AdminAssembly",
					Params:    "{}",
					Createdat: uint64(time.Now().Unix()),
					Day:       time.Now().Format("20060102"),
					Hash:      "",
					Role:      "admin",
				})
			}
		}

	}
}

func InitTable() {
	//cert_model.InitTable()
	//asset_model.InitTable()
	access_gateway_model.InitTable()
}
