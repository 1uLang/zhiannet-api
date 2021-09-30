package channels_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"testing"
)

func init() {
	model.InitMysqlLink()
}

func Test_node_num(t *testing.T) {
	list, total, e := GetList(&channels.ChannelReq{
		PageNum:  1,
		PageSize: 100,
	})
	fmt.Println(list, total, e)
}

func Test_add(t *testing.T) {
	list, e := Add(&channels.Channels{
		Name: "张三俱乐部",
		User: "老张",
	})
	fmt.Println(list, e)
}

func Test_list(t *testing.T) {
	list, _, e := GetList(&channels.ChannelReq{
		PageNum:  1,
		PageSize: 99,
		Status:   "1",
	})
	fmt.Println(list, e)
}
