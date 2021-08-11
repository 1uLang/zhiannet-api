package conversation

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}
func Test_ips_alarm_Iface(t *testing.T) {
	list, err := GetList(&ConReq{
		NodeId: 2,
	})
	fmt.Println(list)
	fmt.Println(err)
}
