package ips

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	InitDB()
}
func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}

//获取ips列表
func Test_ips_list(t *testing.T) {
	list, err := GetIpsList(&IpsReq{NodeId: 1})
	fmt.Println(list)
	fmt.Println(err)
}

//停用 ips规则
func Test_stop(t *testing.T) {
	res, err := EditIps(&EditIpsReq{NodeId: 1, Sid: 2001223})
	fmt.Println(res)
	fmt.Println(err)
}
