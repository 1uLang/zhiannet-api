package global_status

import (
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/global_status"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	GlobalReq struct {
		NodeId uint64 `json:"node_id"`
	}
)

//全局状态
func GetGlobalStatus(req *GlobalReq) (res *global_status.GlobalStatus, err error) {
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
	return global_status.GetGlobal(loginInfo)
}
