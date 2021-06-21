package monitor_list

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
func Test_list(t *testing.T) {
	res, total, err := GetList(&ListReq{
		MonitorType: 1,
		PageNum:     0,
		PageSize:    10,
	})
	fmt.Println(res[0])
	fmt.Println(total)
	fmt.Println(err)
}

func Test_del(t *testing.T) {
	Del(&DelReq{1})
}
