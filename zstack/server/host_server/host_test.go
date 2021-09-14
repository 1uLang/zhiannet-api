package host_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/zstack/request/host"
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//云主机
func Test_host(t *testing.T) {
	//res, err := AllHostList(&host.HostListReq{})
	//logs.Println(res)
	//logs.Println(err)
	//
	//for k, v := range res.Inventories {
	//	fmt.Println(k, v.State)
	//}

	ch := &CheckHost{}
	ch.Run()

}

//list
func Test_host_list(t *testing.T) {
	res, err := HostList(&host.HostListReq{})
	logs.Println(res)
	logs.Println(err)

	for k, v := range res.Inventories {
		fmt.Println(k, v.State)
	}

}

func Test_host_spec(t *testing.T) {
	res, err := UpdateSpec(&host.UpdateSpecReq{
		SpecUUid: "8d7866ddb343409a914b10504743e637",
		HostUUid: "60a06990818e4b1494ca734062f65dcf",
	})
	logs.Println(res)
	logs.Println(err)
}
