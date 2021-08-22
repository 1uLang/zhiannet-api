package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	"github.com/1uLang/zhiannet-api/hids/model/user"
	"github.com/1uLang/zhiannet-api/hids/request"
	"time"
)

/*
	nessus 主机扫描api对接 server 层
*/
type CheckRequest struct{}

// SetUrl 初始化 Nessus APIKeys
func SetUrl(url string) error {
	return request.InitServerUrl(url)
}

// SetAPIKeys 初始化 Nessus APIKeys
func SetAPIKeys(req *request.APIKeys) error {
	model.HidsUserNameAPI = "dengbao"
	return request.InitRequestAPIKeys(req)
}

//获取主机防护系统 节点信息
func GetHideInfo() (resp *model.HidsResp, err error) {
	return model.GetHidsInfo()
}

//检测hids 配置是否正常
func Check() (bool, uint64, int, error) {

	info, err := GetHideInfo()
	if err != nil {
		return false, 0, 0, err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	err = SetAPIKeys(&request.APIKeys{info.AppId, info.Secret})
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	_, err = user.List(&user.SearchReq{})
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	return true, info.Id, info.ConnState, nil
}

func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hids-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, oldConn, _ := Check()
	if id > 0 {
		if !res {
			conn = 0
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "主机防护状态不可用",
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})
		}
		if conn != oldConn {
			subassemblynode.UpdateConnState(id, conn)
			if conn == 1 {
				edge_messages.Add(&edge_messages.Edgemessages{
					Level:     "success",
					Subject:   "组件状态恢复正常",
					Body:      "主机防护恢复可用状态",
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

	agent.InitTable()
}
