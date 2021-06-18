package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/logs"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	LogReq struct {
		NodeId   uint64 `json:"node_id"`
		Keyword  string `json:"keyword"`
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
)

//获取日志列表
func GetLogsList(req *LogReq) (list *logs.LogListResp, err error) {
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
	//设置请求接口必须的cookie
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	return logs.GetLogsList(&logs.LogReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, loginInfo)
}

//清除日志
func ClearLogs(req *NodeReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return logs.ClearLog(loginInfo)
}
