package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/model/dashboard"
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"time"
)

/*
	awvs web漏洞扫描api对接 server 层
*/
type (
	CheckRequest struct{}
)

// SetUrl 初始化 awvs url or ip:port
func SetUrl(url string) error {
	return request.InitServerUrl(url)
}

// SetAPIKeys 初始化 awvs APIKeys
func SetAPIKeys(req *request.APIKeys) error {
	ok, err := req.Check()
	if err != nil {
		return err
	}
	if ok {
		return request.InitRequestXAuth(req)
	} else {
		return fmt.Errorf("参数错误")
	}
}

func GetWebScan() (resp *model.WebScanResp, err error) {
	return model.GetWebScanInfo()
}

//检测awvs 是否配置正常
func Check() (bool, uint64, error) {

	info, err := GetWebScan()
	if err != nil {
		return false, 0, err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false, info.Id, err
	}
	err = SetAPIKeys(&request.APIKeys{XAuth: info.Key})
	if err != nil {
		return false, info.Id, err
	}
	_, err = dashboard.MeStats()
	if err != nil {
		return false, info.Id, err
	}
	return true, info.Id, nil
}

func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("awvs-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, _ := Check()
	if !res {
		conn = 0
		edge_messages.Add(&edge_messages.Edgemessages{
			Level:     "error",
			Subject:   "组件状态异常",
			Body:      "Web漏洞扫描状态不可用",
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
	targets.InitTable()
	reports.InitTable()
}
