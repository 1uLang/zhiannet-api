package server

import (
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/user"
	"github.com/1uLang/zhiannet-api/hids/request"
)

/*
	nessus 主机扫描api对接 server 层
*/

// SetUrl 初始化 Nessus APIKeys
func SetUrl(url string) error {
	return request.InitServerUrl(url)
}

// SetAPIKeys 初始化 Nessus APIKeys
func SetAPIKeys(req *request.APIKeys) error {
	return request.InitRequestAPIKeys(req)
}

//获取主机防护系统 节点信息
func GetHideInfo() (resp *model.HidsResp, err error) {
	return model.GetHidsInfo()
}

//检测hids 配置是否正常
func Check() (bool,error) {

	info,err := GetHideInfo()
	if err != nil {
		return false,err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false,err
	}
	err = SetAPIKeys(&request.APIKeys{info.AppId,info.Secret})
	if err != nil {
		return false,err
	}
	_,err = user.List(&user.SearchReq{})
	if err != nil {
		return false,err
	}
	return true,nil
}