package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/ddos/request/global_status"
	"testing"
)

func Test_ApiLogin(t *testing.T) {
	cook, err := Login()
	fmt.Println(cook, err)
}

//全局统计  GetStatusGlobal
func Test_status_global(t *testing.T) {
	global_status.GetStatusGlobal()
}

//负载信息
func Test_load(t *testing.T) {
	global_status.GetLoad()
}

////主机列表
//func Test_host(t *testing.T){
//	HostStatusApi()
//}
