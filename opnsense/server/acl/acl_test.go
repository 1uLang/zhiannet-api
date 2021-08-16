package acl

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/tidwall/gjson"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//获取列表
func Test_acl_list(t *testing.T) {
	list, err := GetAclList(2)
	r, _ := json.Marshal(list)
	fmt.Println(gjson.ParseBytes(r).Value())
	fmt.Println(err)
}

//获取详情
func Test_acl_info(t *testing.T) {
	res, err := GetAclInfo(&InfoReq{NodeId: 2, ID: "5"})
	r, _ := json.Marshal(res)
	fmt.Println(gjson.ParseBytes(r).Value())
	fmt.Println(err)
}

//添加修改
func Test_acl_save(t *testing.T) {
	res, err := SaveAcl(&SaveAclReq{NodeId: 2, ID: "5",
		Type:         "pass",
		Disabled:     false,
		Interface:    "lan",
		Direction:    "in",
		Ipprotocol:   "inet",
		Protocol:     "tcp",
		Src:          "192.1.1.1",
		Srcmask:      "24",
		Dst:          "112.1.1.1",
		Dstmask:      "24",
		Descr:        "api create 1212",
		SrcBeginPort: "12345",
		SrcEndPort:   "12345",
	})
	r, _ := json.Marshal(res)
	fmt.Println(gjson.ParseBytes(r).Value())
	fmt.Println(err)
}

//删除
func Test_del(t *testing.T) {
	res, err := DelAcl(&DelAclReq{NodeId: 12, Interface: "lan", ID: "8"})
	fmt.Println(res)
	fmt.Println(err)
}

//启动停止
func Test_start(t *testing.T) {
	res, err := StartUpAcl(&StartAclReq{NodeId: 12, Interface: "lan", ID: "9"})
	fmt.Println(res)
	fmt.Println(err)
}
