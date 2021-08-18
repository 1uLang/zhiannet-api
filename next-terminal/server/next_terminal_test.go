package server

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}
func Test_conn(t *testing.T) {
	ls := new(CheckRequest)
	ls.Run()
}
