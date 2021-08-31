package clamav

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/clamav"

	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
	LogReq struct {
		NodeId       uint64 `json:"node_id"`
		Current      int    `json:"current"`       //页数
		RowCount     int    `json:"row_count"`     //每页条数
		SearchPhrase string `json:"search_phrase"` //关键词
	}
)

func GetClamAV(req *NodeReq) (list *clamav.ClamAVResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}

	//设置请求接口必须的cookie
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	return clamav.GetClamAV(loginInfo)
}

func GetLog(req *LogReq) (list *clamav.LogResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}
	//设置请求接口必须的cookie
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}

	//fmt.Println(loginInfo)
	//return
	if req.Current <= 0 {
		req.Current = 1
	}
	if req.RowCount <= 0 {
		req.RowCount = 20
	}

	return clamav.Log(&clamav.LogReq{
		SearchPhrase: req.SearchPhrase,
		Current:      req.Current,
		RowCount:     req.RowCount,
	}, loginInfo)
}
