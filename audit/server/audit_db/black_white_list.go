package audit_db

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/black_white_list"
	"github.com/1uLang/zhiannet-api/ddos/server"
)

type (
	BWReq struct {
		NodeId uint64 `json:"node_id"`
		Addr   string `json:"addr"`
		Page   int    `json:"page"`
	}
	EditBWReq struct {
		NodeId uint64   `json:"node_id"`
		Addr   []string `json:"addr"`
		White  bool     `json:"white"`
	}
)

//黑白名单列表
func GetBWList(req *BWReq) (list *black_white_list.StatusBwlist, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	list, err = black_white_list.BWList(&black_white_list.BWListReq{Addr: req.Addr, Page: req.Page}, logReq, true)
	return

}

//添加黑白名单
func AddBW(req *EditBWReq) (res *black_white_list.Success, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = black_white_list.AddBW(&black_white_list.EditBWReq{Addr: req.Addr, White: req.White}, logReq, true)
	return

}

//删除黑白名单
func DeleteBW(req *EditBWReq) (res *black_white_list.Success, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = black_white_list.DeleteBW(&black_white_list.EditBWReq{Addr: req.Addr}, logReq, true)
	return

}
