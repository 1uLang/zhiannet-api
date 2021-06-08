package host_status

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
)

type (
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
)

//ddos节点列表
func GetDdosNodeList() (list []*subassemblynode.Subassemblynode, err error) {
	list, err = subassemblynode.GetList(&subassemblynode.NodeReq{Type: 1})
	return
}

//主机状态
//func GetHostStatus() {
//	req := &host_status.HostReq{}
//	host_status.HostStatus(req, true)
//}

//主机列表
func GetHostList(req *ddos_host_ip.HostReq) (hostList []*host_status.StatusHost, err error) {
	//先从数据库获取ip列表
	list, err := ddos_host_ip.GetList(req)
	if err != nil {
		return
	}

	apiReq := &host_status.HostReq{}
	for _, v := range list {
		apiReq.Addr = append(apiReq.Addr, v.Addr)
	}
	//获取节点账号信息
	nodeInfo, err := subassemblynode.GetNodeInfoById(req.NodeId)
	if err != nil {
		return
	}
	logReq := &request.LoginReq{
		Name:     nodeInfo.Key,
		Password: nodeInfo.Secret,
		Addr:     nodeInfo.Addr,
		Port:     fmt.Sprintf("%v", nodeInfo.Port),
	}
	hostList, err = host_status.HostList(apiReq, logReq, true)
	return
}

//添加高仿IP
func AddAddr() {

}
