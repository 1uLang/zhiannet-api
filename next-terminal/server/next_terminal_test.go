package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
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
func TestAsset_List(t *testing.T) {
	req, err := NewServerRequest("http://156.249.24.77:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	list, total, err := req.Assets.List(&asset_model.ListReq{UserId: 1})
	fmt.Println(list, total, err)
}
