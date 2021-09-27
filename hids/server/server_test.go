package server

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

//获取主机防护系统节点信息
func Test_get_hids(t *testing.T) {
	info, err := GetHideInfo()
	fmt.Println(info)
	fmt.Println(err)
}

func Test_StatisticsRequest(t *testing.T) {
	n := &StatisticsRequest{}
	n.Run()
}
