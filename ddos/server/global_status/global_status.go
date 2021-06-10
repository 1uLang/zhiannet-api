package global_status

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/global_status"
	"github.com/1uLang/zhiannet-api/ddos/server"
	"net/http"
)

type (
	StatusReq struct {
		NodeId uint64 `json:"node_id"`
	}
)

//获取全局统计
func GetStatusGlobal(req *StatusReq) (res *global_status.StatusGlobal, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = global_status.GetStatusGlobal(logReq, true)
	return res, err
}

//获取负载信息
func GetLoad(req *StatusReq) (res *global_status.StatusHealth, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = global_status.GetLoad(logReq, true)
	return res, err
}

func GlobalImg(w http.ResponseWriter, r *http.Request) {
	model.InitMysqlLink()
	cache.InitClient()
	req := StatusReq{
		NodeId: 1,
	}
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err := server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	global_status.GlobalImg(logReq, true)
	return
}
