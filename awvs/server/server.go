package server

import (
	"fmt"
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
