package host_status

import (
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
	"github.com/1uLang/zhiannet-api/ddos/server"
	"strconv"
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
	HostListResp struct { //主机列表响应参数
		HostId       uint64  `json:"host_id"`
		Addr         string  `json:"addr"`
		BandwidthIn  float64 `json:"bandwidth_in"`  //带宽in input_bps
		BandwidthOut float64 `json:"bandwidth_out"` //带宽out output_bps
		RateSyn      float64 `json:"rate_syn"`      //频率 syn
		RateAck      float64 `json:"rate_ack"`      //频率 ack
		RateUdp      float64 `json:"rate_udp"`      //频率 udp
		RateIcmp     float64 `json:"rate_icmp"`     //频率 icmp
		RateFrag     float64 `json:"rate_frag"`     //频率 frag
		RateNonip    float64 `json:"rate_nonip"`    //频率 nonip
		RateNewTcp   float64 `json:"rate_new_tcp"`  //频率 new_tcp
		RateNewUdp   float64 `json:"rate_new_udp"`  //频率 new_udp
		TcpConnIn    float64 `json:"tcp_conn_in"`   //tcp in 连接数
		TcpConnOut   float64 `json:"tcp_conn_out"`  //tcp out 连接数
		UdpConn      float64 `json:"udp_conn"`      //udp  连接数
		UserName     string  `json:"user_name"`     //用户名
		Remark       string  `json:"remark"`        //备注
		NodeId       uint64  `json:"node_id"`
	}
)

//获取登陆的账号信息
//func GetLoginInfo(req NodeReq) (logReq *audit_db.LoginReq, err error) {
//	var nodeInfo subassemblynode.Subassemblynode
//	//获取节点账号信息
//	nodeInfo, err = subassemblynode.GetNodeInfoById(req.NodeId)
//	if err != nil {
//		return
//	}
//	logReq = &audit_db.LoginReq{
//		Name:     nodeInfo.Key,
//		Password: nodeInfo.Secret,
//		Addr:     nodeInfo.Addr,
//		Port:     fmt.Sprintf("%v", nodeInfo.Port),
//	}
//	return
//}
//ddos节点信息
func GetDDoSNodeInfo(id uint64) (*subassemblynode.Subassemblynode, error) {
	return subassemblynode.GetNodeInfoById(id)
}

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
func GetHostList(req *ddos_host_ip.HostReq) (lists []*HostListResp, total int64, err error) {
	//先从数据库获取ip列表
	var list []*ddos_host_ip.DdosHostIp
	list, total, err = ddos_host_ip.GetList(req)
	if err != nil || total == 0 {
		return
	}
	hostMap := make(map[string]uint64, len(list))
	itemsMap := make(map[string]*ddos_host_ip.DdosHostIp, len(list))
	//userMap := make(map[string]uint64, 0)
	apiReq := &host_status.HostReq{}
	for _, v := range list {
		//对应IP地址
		apiReq.Addr = append(apiReq.Addr, v.Addr)
		hostMap[v.Addr] = v.Id
		itemsMap[v.Addr] = v
		//对应用户ID
		//userMap[v.Addr] = v.UserId
	}
	//获取节点信息
	logReq, err := server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil {
		return
	}
	hostList, err := host_status.HostList(apiReq, logReq, true)
	if err != nil || hostList == nil {
		return
	}
	lists = make([]*HostListResp, len(hostList))
	//all := &HostListResp{
	//	Addr: "all",
	//}
	//使用uid 获取用户信息
	//userList, _, err := GetUserInfoByUid(userMap)
	//if err != nil {
	//	return
	//}

	for k, v := range hostList { //所有ip的数据
		l := &HostListResp{
			Addr:   v.Netaddr,
			Remark: itemsMap[v.Netaddr].Remark,
			NodeId: itemsMap[v.Netaddr].NodeId,
			//UserName: "",
		} //获取用户信息
		//if userId, ok := userMap[v.Netaddr]; ok {
		//	if username, ok := userList[userId]; ok {
		//		l.UserName = username.Username
		//	}
		//}
		if id, ok := hostMap[v.Netaddr]; ok { //ddos_host_id表的ID
			l.HostId = id
		}
		if len(v.Host) > 0 {
			for _, y := range v.Host {
				if y.Address == l.Addr { //当前IP的数据

					l.BandwidthIn, _ = strconv.ParseFloat(y.InputBps, 64)
					l.BandwidthOut, _ = strconv.ParseFloat(y.OutputBps, 64)
					l.RateSyn, _ = strconv.ParseFloat(y.SynRate, 64)
					l.RateAck, _ = strconv.ParseFloat(y.AckRate, 64)
					l.RateUdp, _ = strconv.ParseFloat(y.UdpRate, 64)
					l.RateIcmp, _ = strconv.ParseFloat(y.IcmpRate, 64)
					l.RateFrag, _ = strconv.ParseFloat(y.FragRate, 64)
					l.RateNonip, _ = strconv.ParseFloat(y.NonipRate, 64)
					l.RateNewTcp, _ = strconv.ParseFloat(y.NewTcpRate, 64)
					l.RateNewUdp, _ = strconv.ParseFloat(y.NewUdpRate, 64)
					l.TcpConnIn, _ = strconv.ParseFloat(y.TcpConnIn, 64)
					l.TcpConnOut, _ = strconv.ParseFloat(y.TcpConnOut, 64)
					l.UdpConn, _ = strconv.ParseFloat(y.UdpConn, 64)

					//相关host数据合计
					//all.BandwidthIn += l.BandwidthIn
					//all.BandwidthOut += l.BandwidthOut
					//all.RateSyn += l.RateSyn
					//all.RateAck += l.RateAck
					//all.RateUdp += l.RateUdp
					//all.RateIcmp += l.RateIcmp
					//all.RateFrag += l.RateFrag
					//all.RateNonip += l.RateNonip
					//all.RateNewTcp += l.RateNewTcp
					//all.RateNewUdp += l.RateNewUdp
					//all.TcpConnIn += l.TcpConnIn
					//all.TcpConnOut += l.TcpConnOut
					//all.UdpConn += l.UdpConn
				}
			}
		}
		lists[k] = l
	}
	return
}

//通过userid  获取用户信息
func GetUserInfoByUid(req map[string]uint64) (list map[uint64]*edge_users.EdgeUsers, total int64, err error) {
	if len(req) == 0 {
		return
	}
	uids := make([]uint64, 0)
	for _, v := range req {
		uids = append(uids, v)
	}
	return edge_users.GetListByUid(uids)
}

func UpdateAddr(req *ddos_host_ip.UpdateHost) (err error) {

	o, err := ddos_host_ip.Info(req.Id)
	if err != nil {
		return err
	}
	if o.NodeId != req.NodeId || o.Addr != req.Addr {
		//先删除 后添加
		err = DeleteAddr([]uint64{req.Id})
		if err != nil {
			return err
		}
		_, err = AddAddr(&req.AddHost)
		return err
	} else {
		_, err = ddos_host_ip.Edit(&req.AddHost, req.Id)
		return err
	}
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
	//fmt.Println("logReq==", logReq)
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

//删除高防IP
func DeleteAddr(host_ids []uint64) error {
	return ddos_host_ip.DeleteByIds(host_ids)
}

//ip详情-屏蔽列表
func GetHostShieldList(req *ShieldReq) (list *host_status.StatusFblink, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	//fmt.Println("logReq==", logReq)
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
	//fmt.Println("logReq==", logReq)
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
	//fmt.Println("logReq==", logReq)
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
	//fmt.Println("logReq==", logReq)
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
	//fmt.Println("logReq==", logReq)
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
