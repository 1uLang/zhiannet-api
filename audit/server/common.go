package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/utils"
)

type (
	UserReq struct { //节点
		UserId      uint64 `json:"user_id"`
		AdminUserId uint64 `json:"admin_user_id"`
	}
)

//获取登陆的账号信息
func GetLoginInfo(req UserReq) (logReq *request.LoginReq, err error) {
	var nodeInfo *subassemblynode.Subassemblynode
	nodeInfo, err = GetAuditInfo()
	logReq = &request.LoginReq{
		Name:     nodeInfo.Key,
		Password: nodeInfo.Secret,
		Addr:     nodeInfo.Addr,
		Port:     fmt.Sprintf("%v", nodeInfo.Port),
		IsSsl:    nodeInfo.IsSsl == 1,
	}
	logReq.Addr = utils.CheckHttpUrl(logReq.Addr, nodeInfo.IsSsl == 1)

	//等保平台 超级管理员
	if req.AdminUserId == 1 {

	}
	return
}

//获取审计系统节点信息
func GetAuditInfo() (nodeInfo *subassemblynode.Subassemblynode, err error) {
	//获取节点账号信息
	nodeInfos, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		State:    "1",
		Type:     6,
		PageSize: 1,
	})
	if err != nil {
		return
	}
	if len(nodeInfos) == 0 {
		return
	}
	nodeInfo = nodeInfos[0]
	return
}
