package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/model/scans"
	"github.com/1uLang/zhiannet-api/nessus/request"
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
	return request.InitRequestAPIKeys(req)
}

func GetNessus() (resp *model.NessusResp, err error) {
	return model.GetNessusInfo()
}

//检测nessus 配置访问是否异常

func Check() (bool, uint64, error) {
	info, err := GetNessus()
	if err != nil {
		return false, 0, err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false, info.Id, err
	}
	err = SetAPIKeys(&request.APIKeys{info.Access, info.Secret})
	if err != nil {
		return false, info.Id, err
	}
	_, err = scans.List(&scans.ListReq{UserId: 1})
	if err != nil {
		return false, info.Id, err
	}
	return true, info.Id, nil
}
func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("nessus-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, _ := Check()
	if !res {
		conn = 0
		edge_messages.Add(&edge_messages.Edgemessages{
			Level:     "error",
			Subject:   "组件状态异常",
			Body:      "主机漏洞扫描状态不可用",
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
	scans.InitTable()
}
