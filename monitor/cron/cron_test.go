package cron

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
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
