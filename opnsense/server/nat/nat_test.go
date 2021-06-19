package nat

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
func Test_nat_list(t *testing.T) {
	list, err := GetNat1To1List(&ListReq{NodeId: 12})
	fmt.Println(list[0])
	fmt.Println(list[1])
	fmt.Println(err)
}

func Test_nat_info(t *testing.T) {
	res, err := GetNat1To1Info(&InfoReq{NodeId: 12, Id: "3"})
	fmt.Println(res)
	fmt.Println(err)
}

//测试新增
func Test_nat_add(t *testing.T) {
	res, err := SaveNat1To1(&SaveNat1To1Req{
		NodeId:    12,
		ID:        "3",
		Interface: "lan",
		Type:      "nat",
		External:  "1.1.1.1/24",
		Src:       "1.2.3.4",
		Srcmask:   "24",
		Dst:       "12.1.1.1",
		Dstmask:   "24",
		Descr:     "api ",
	})
	fmt.Println(res)
	fmt.Println(err)
}

//启动停止
func Test_start_up(t *testing.T) {
	res, err := StartUpNat1To1(&StartNat1To1Req{NodeId: 12, Id: "3"})
	fmt.Println(res)
	fmt.Println(err)
}

//删除
func Test_del(t *testing.T) {
	res, err := DelNat1To1(&DelNat1To1Req{NodeId: 12, Id: "2"})
	fmt.Println(res)
	fmt.Println(err)
}

//应用
func Test_apply(t *testing.T) {
	res, err := ApplyNat1To1(12)
	fmt.Println(res)
	fmt.Println(err)
}
