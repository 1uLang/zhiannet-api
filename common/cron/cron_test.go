package cron

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/cron/logs"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_db_nodes"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}
func Test_cron(t *testing.T) {
	//PortPing()
	InitCron()

	time.Sleep(time.Minute * 30)
}

func Test_conn_db(t *testing.T) {

	res, err := edge_db_nodes.NewConn()
	fmt.Println(res, err)
}

func Test_statistice_waf(t *testing.T) {

	sta := new(logs.StatisticsWAFLogs)
	sta.Run()
}

func Test_statistice_ddos(t *testing.T) {

	sta := new(logs.StatisticsDDOSLogs)
	sta.Run()
}

func Test_statistice_nfw(t *testing.T) {
	sta := new(logs.StatisticsNFWLogs)
	sta.Run()
}
