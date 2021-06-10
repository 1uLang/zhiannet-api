package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/request"
)

type (
	NodeReq struct { //节点
		NodeId uint64 `json:"node_id"`
	}
)

//获取登陆的账号信息
func GetLoginInfo(req NodeReq) (logReq *request.LoginReq, err error) {
	var nodeInfo subassemblynode.Subassemblynode
	//获取节点账号信息
	nodeInfo, err = subassemblynode.GetNodeInfoById(req.NodeId)
	if err != nil {
		return
	}
	logReq = &request.LoginReq{
		Name:     nodeInfo.Key,
		Password: nodeInfo.Secret,
		Addr:     nodeInfo.Addr,
		Port:     fmt.Sprintf("%v", nodeInfo.Port),
	}
	return
}
