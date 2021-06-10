package global_status

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

//全局状态
func Test_Global(t *testing.T) {
	InitDB()
	res, err := GetStatusGlobal(&StatusReq{
		NodeId: 1,
	})
	fmt.Println(res)
	fmt.Println(err)
}

//负载

func Test_load(t *testing.T) {
	InitDB()
	res, err := GetLoad(&StatusReq{
		NodeId: 1,
	})
	fmt.Println(res)
	fmt.Println(err)
}

//img
func Test_global_img(t *testing.T) {
	InitDB()

}
