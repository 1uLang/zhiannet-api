package host_status

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"testing"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}
func Test_ddos_list(t *testing.T) {
	model.InitMysqlLink()
	list, err := GetDdosNodeList()
	fmt.Println("data ====", list[0])
	fmt.Println(list, err)
}

//主机状态
func Test_host_status(t *testing.T) {
	//GetHostStatus()
}

//主机列表
func Test_host_list(t *testing.T) {
	InitDB()
	list, err := GetHostList(&ddos_host_ip.HostReq{NodeId: 1})
	fmt.Println(list[0])
	fmt.Println(err)
}
