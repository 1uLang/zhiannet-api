package global_status

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//全局状态
func Test_Global(t *testing.T) {
	GetGlobalStatus(&GlobalReq{
		NodeId: 12,
	})
	//fmt.Println(res)
	//fmt.Println(err)
}
