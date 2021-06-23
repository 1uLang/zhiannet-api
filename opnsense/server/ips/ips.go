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
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
	EditActionReq struct {
		NodeId uint64 `json:"node_id"`
		Sid    int64  `json:"sid"`
		Action string `json:"action"`
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
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	return ips.GetIpsList(&ips.IpsReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, loginInfo)
}

//ips规则启动 停止
func EditIps(req *EditIpsReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
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
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.DelIps(&ips.DelIpsReq{
		Sid: req.Sid,
	}, loginInfo)
}

//应用 使规则生效
func ApplyIps(req *NodeReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.ApplyIps(loginInfo)
}

//drop 或者 alert
func EditAction(req *EditActionReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.EditActionIps(&ips.EditActionIpsReq{
		Sid:    req.Sid,
		Action: req.Action,
	}, loginInfo)
}
