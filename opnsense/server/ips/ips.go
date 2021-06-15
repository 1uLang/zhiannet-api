package ips

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/ips"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	IpsReq struct {
		NodeId   uint64 `json:"node_id"`
		Keyword  string `json:"keyword"`
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
	EditIpsReq struct {
		NodeId uint64 `json:"node_id"`
		Sid    int64  `json:"sid"`
	}
	DelIpsReq struct {
		NodeId uint64   `json:"node_id"`
		Sid    []string `json:"sid"`
	}
)

//获取日志列表
func GetIpsList(req *IpsReq) (list *ips.IpsListResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	return ips.GetIpsList(&ips.IpsReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//ips规则启动 停止
func EditIps(req *EditIpsReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return ips.EditIps(&ips.EditIpsReq{
		Sid: req.Sid,
	}, loginInfo)
}

//ips规则删除
func DelIps(req *DelIpsReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return ips.DelIps(&ips.DelIpsReq{
		Sid: req.Sid,
	}, loginInfo)
}
