package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/utils"
)

type (
	NodeReq struct { //节点
		NodeId uint64 `json:"node_id"`
	}
)

//获取登陆的账号信息
func GetLoginInfo(req NodeReq) (logReq *request.ApiKey, err error) {
	var nodeInfo *subassemblynode.Subassemblynode
	//获取节点账号信息
	nodeInfo, err = subassemblynode.GetNodeInfoById(req.NodeId)
	if err != nil {
		return
	}
	logReq = &request.ApiKey{
		Username: nodeInfo.Key,
		Password: nodeInfo.Secret,
		Addr:     nodeInfo.Addr,
		Port:     fmt.Sprintf("%v", nodeInfo.Port),
		IsSsl:    nodeInfo.IsSsl == 1,
	}
	logReq.Addr = utils.CheckHttpUrl(logReq.Addr, nodeInfo.IsSsl == 1)
	return
}

//云防火墙 节点列表
func GetOpnsenseNodeList() (list []*subassemblynode.Subassemblynode, total int64, err error) {
	list, total, err = subassemblynode.GetList(&subassemblynode.NodeReq{Type: 2, State: "1"})
	return
}
