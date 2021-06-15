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
	list, err := GetIpsList(&IpsReq{NodeId: 12})
	fmt.Println(list)
	fmt.Println(err)
}

//停用 启动 ips规则
func Test_stop(t *testing.T) {
	res, err := EditIps(&EditIpsReq{NodeId: 12, Sid: 2001223})
	fmt.Println(res)
	fmt.Println(err)
}

//删除 ips规则
func Test_del(t *testing.T) {
	res, err := DelIps(&DelIpsReq{NodeId: 12, Sid: []string{"2000006", "2000007"}})
	fmt.Println(res)
	fmt.Println(err)
}
