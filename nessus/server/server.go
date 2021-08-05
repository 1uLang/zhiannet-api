package server

import (
	"github.com/1uLang/zhiannet-api/hids/model/user"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/request"
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

func GetNessus() (resp *model.NessusResp, err error) {
	return model.GetNessusInfo()
}

func Check()(bool,error)  {
	info,err := GetNessus()
	if err != nil {
		return false,err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false,err
	}
	err = SetAPIKeys(&request.APIKeys{info.Access,info.Secret})
	if err != nil {
		return false,err
	}
	_,err = user.List(&user.SearchReq{})
	if err != nil {
		return false,err
	}
	return true,nil
}