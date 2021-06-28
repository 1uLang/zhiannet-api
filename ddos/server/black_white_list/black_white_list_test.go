package black_white_list

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}

//黑白名单列表
func Test_bw_list(t *testing.T) {
	InitDB()
	list, err := GetBWList(&BWReq{NodeId: 6, Addr: "8.210.100.191"})
	fmt.Println(list)
	fmt.Println(err)
}

//添加黑白名单
func Test_add_bw(t *testing.T) {
	InitDB()
	list, err := AddBW(&EditBWReq{NodeId: 1, Addr: []string{"1.62.2.37"}, White: false})
	fmt.Println(list)
	fmt.Println(err)
}

//删除黑白名单
func Test_del_bw(t *testing.T) {
	InitDB()
	list, err := DeleteBW(&EditBWReq{NodeId: 1, Addr: []string{"1.62.2.37"}})
	fmt.Println(list)
	fmt.Println(err)
}
