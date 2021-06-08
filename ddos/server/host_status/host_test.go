package host_status

import (
	"github.com/1uLang/zhiannet-api/ddos/request/host_status"
	"testing"
)

//主机状态
func Test_host_status(t *testing.T) {
	GetHostStatus()
}

//主机列表
func Test_host_list(t *testing.T) {
	GetHostList(&host_status.HostReq{})
}
