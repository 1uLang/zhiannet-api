package host_server

import "github.com/1uLang/zhiannet-api/zstack/request/host"

func HostList(req *host.HostListReq) (resp *host.HostListResp, err error) {

	return host.HostList(req)
}

//
