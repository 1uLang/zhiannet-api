package logs_statistics_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/iwind/TeaGo/logs"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
}

func Test_log_statistice(t *testing.T) {
	res, err := GetWafStatistics([]int64{1}, 0, 1)

	logs.Println(res, err)
	for _, v := range res {
		fmt.Println(v)
	}
}

func Test_time(t *testing.T) {
	Str := "Thu Jul 15 10:39:10 2021"
	stime, err := time.ParseInLocation("Mon Jan _2 15:04:05 2006", Str, time.Local)
	fmt.Println(stime, err)
}
