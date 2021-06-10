package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/logs"
	"github.com/1uLang/zhiannet-api/ddos/server"
)

type (
	LinkLogReq struct {
		Addr   string `json:"addr"`
		Level  int    `json:"level"`
		NodeId uint64 `json:"node_id"`
	}
)

func GetLinkLogList(req *LinkLogReq) (list *logs.LogsReportLink, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	if req.Level == 0 {
		req.Level = 1
	}
	list, err = logs.LinkLogList(&logs.LinkLogReq{
		Addr:  req.Addr,
		Level: req.Level,
	},
		logReq, true)
	return

}
