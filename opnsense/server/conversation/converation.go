package conversation

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	conversation_req "github.com/1uLang/zhiannet-api/opnsense/request/conversation"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	ConReq struct {
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
func GetList(req *ConReq) (list *conversation_req.ListResp, err error) {
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
	return conversation_req.GetList(&conversation_req.ConReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, loginInfo)
}
