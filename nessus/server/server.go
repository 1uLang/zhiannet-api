package server

import (
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
