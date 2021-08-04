package edge_logs_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_logs"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
}

func TestGetLogList(t *testing.T) {
	list, total, err := GetLogList(&edge_logs.UserLogReq{
		UserId: 1,
	})
	fmt.Println(list[0])
	fmt.Println(total)
	fmt.Println(err)
}

func Test_time(t *testing.T) {
	str := ""
	stime, _ := time.ParseInLocation("2006-01-02", str, time.Local)
	fmt.Println(stime.Unix())
}
