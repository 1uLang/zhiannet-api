package clamav

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
func Test_clamav(t *testing.T) {
	list, err := GetClamAV(&NodeReq{
		NodeId: 2,
	})
	fmt.Println(list)
	fmt.Println(err)
}

func Test_clamav_log(t *testing.T) {
	list, err := GetLog(&LogReq{
		NodeId: 2,
	})
	fmt.Println(list)
	fmt.Println(err)
}
