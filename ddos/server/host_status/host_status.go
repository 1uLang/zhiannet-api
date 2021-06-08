package host_status

import "github.com/1uLang/zhiannet-api/ddos/request/host_status"

//主机状态
func GetHostStatus() {
	req := &host_status.HostReq{}
	host_status.HostStatus(req, true)
}

//主机列表
func GetHostList(req *host_status.HostReq) {
	req = &host_status.HostReq{}
	host_status.HostList(req, true)
}
