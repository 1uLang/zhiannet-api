package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/model/dashboard"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

/*
	awvs web漏洞扫描api对接 server 层
*/

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
func Check() (bool, error) {

	info,err := GetWebScan()
	if err != nil {
		return false,err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false,err
	}
	err = SetAPIKeys(&request.APIKeys{XAuth: info.Key})
	if err != nil {
		return false,err
	}
	_,err = dashboard.MeStats()
	if err != nil {
		return false,err
	}
	return true,nil
}