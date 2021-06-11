package host_status

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
	"github.com/1uLang/zhiannet-api/ddos/server"
)

type (
	ShieldReq struct { //屏蔽列表请求参数
		Addr   string `json:"addr"`
		Page   int    `json:"page"`
		NodeId uint64 `json:"node_id"`
	}

	ReleaseShieldReq struct { //释放屏蔽请求参数
		Addr   []string `json:"addr"`
		NodeId uint64   `json:"node_id"`
	}

	LinkReq struct {
		Addr   string `json:"addr"`
		Page   int    `json:"page"`
		NodeId uint64 `json:"node_id"`
	}

	HostGetReq struct {
		NodeId uint64 `json:"node_id"`
		Addr   string `json:"addr"`
	}
	HostSetReq struct {
		NodeId     uint64 `json:"node_id"`
		Addr       string `json:"addr"`
		Ignore     bool   `json:"ignore"`      //忽略所有流量
		ProtectSet int    `json:"protect_set"` //防护参数集
		FilterSet  int    `json:"filter_set"`  //过滤参数集
		SetTcp     int    `json:"set_tcp"`     //tcp端口集
		SetUdp     int    `json:"set_udp"`     //udp端口集
	}
)

//获取登陆的账号信息
//func GetLoginInfo(req NodeReq) (logReq *request.LoginReq, err error) {
//	var nodeInfo subassemblynode.Subassemblynode
//	//获取节点账号信息
//	nodeInfo, err = subassemblynode.GetNodeInfoById(req.NodeId)
//	if err != nil {
//		return
//	}
//	logReq = &request.LoginReq{
//		Name:     nodeInfo.Key,
//		Password: nodeInfo.Secret,
//		Addr:     nodeInfo.Addr,
//		Port:     fmt.Sprintf("%v", nodeInfo.Port),
//	}
//	return
//}

//ddos节点列表
func GetDdosNodeList() (list []*subassemblynode.Subassemblynode, total int64, err error) {
	list, total, err = subassemblynode.GetList(&subassemblynode.NodeReq{Type: 1, State: "1"})
	return
}

//主机状态
//func GetHostStatus() {
//	req := &host_status.HostReq{}
//	host_status.HostStatus(req, true)
//}

//主机列表
func GetHostList(req *ddos_host_ip.HostReq) (hostList []*host_status.StatusHost, total int64, err error) {
	//先从数据库获取ip列表
	var list []*ddos_host_ip.DdosHostIp
	list, total, err = ddos_host_ip.GetList(req)
	if err != nil || total == 0 {
		return
	}

	apiReq := &host_status.HostReq{}
	for _, v := range list {
		apiReq.Addr = append(apiReq.Addr, v.Addr)
	}
	//获取节点信息
	logReq, err := server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil {
		return
	}
	hostList, err = host_status.HostList(apiReq, logReq, true)
	return
}

//添加高仿IP
func AddAddr(req *ddos_host_ip.AddHost) (id uint64, err error) {
	//添加ip到数据库
	id, err = ddos_host_ip.Add(req)
	if err != nil {
		return
	}
	//调取接口添加
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	err = host_status.AddAddr(req.Addr, logReq, true)
	if err != nil {
		//api添加失败 删除已添加的数据
		ddos_host_ip.DeleteByIds([]uint64{id})
	}
	return
}

//ip详情-屏蔽列表
func GetHostShieldList(req *ShieldReq) (list *host_status.StatusFblink, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	if req.Page <= 1 {
		req.Page = 1
	}
	list, err = host_status.HostShieldList(&host_status.ShieldListReq{Page: req.Page, Addr: req.Addr}, logReq, true)
	return
}

//释放屏蔽列表
func ReleaseShield(req *ReleaseShieldReq) (err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	_, err = host_status.ReleaseShield(&host_status.ReleaseShieldReq{Addr: req.Addr}, logReq, true)
	return

}

//链接监控列表
func GetLinkList(req *LinkReq) (list *host_status.StatusLink, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	list, err = host_status.LinkList(&host_status.LinkListReq{Addr: req.Addr, Page: req.Page}, logReq, true)
	return

}

//获取主机信息
func GetHostInfo(req *HostGetReq) (res *host_status.StatusHostResp, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = host_status.GetHostInfo(&host_status.HostReq{
		Addr: []string{req.Addr},
	}, logReq, true)
	return
}

//主机设置  b
func SetHost(req *HostSetReq) (res *host_status.Success, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	res, err = host_status.SetHost(&host_status.HostSetReq{
		Addr:       req.Addr,
		Ignore:     req.Ignore,
		ProtectSet: req.ProtectSet,
		FilterSet:  req.FilterSet,
		SetTcp:     req.SetTcp,
		SetUdp:     req.SetUdp,
	}, logReq, true)
	return

}
