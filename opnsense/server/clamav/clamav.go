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
