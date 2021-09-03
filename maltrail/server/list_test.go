package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

func Test_list(t *testing.T) {
	list, err := GetList(&ListReq{
		Date: "2021-09-01",
	})
	if len(list) > 0 {
		for _, v := range list {
			fmt.Println(v.Trail, v.Info, v.Reference)
		}
	}
	fmt.Println(err)
}
