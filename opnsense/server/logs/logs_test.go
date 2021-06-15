package logs

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

//获取日志列表
func Test_logs_list(t *testing.T) {
	list, err := GetLogsList(&LogReq{NodeId: 1})
	fmt.Println(list)
	fmt.Println(err)
}

//清除日志
func Test_logs_clear(t *testing.T) {
	res, err := ClearLogs(&NodeReq{NodeId: 1})
	fmt.Println(res)
	fmt.Println(err)
}
