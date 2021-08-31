package logs_statistics_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/iwind/TeaGo/logs"
	"testing"
)

func init() {
	model.InitMysqlLink()
}

func Test_log_statistice(t *testing.T) {
	res, err := GetWafStatistics([]int64{1}, 0)

	logs.Println(res, err)
	for _, v := range res {
		fmt.Println(v)
	}
}
