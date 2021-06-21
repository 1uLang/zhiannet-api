package examine

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/examine"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
	"testing"
)

func init() {
	//初始化 nessus 服务器地址
	err := server.SetUrl("https://user.cloudhids.net")
	if err != nil {
		panic(err)
	}
	//初始化 nessus 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		AppId:  "39rkz",
		Secret: "tkvgpvjuht2625mo",
	})
	if err != nil {
		panic(err)
	}
}
func TestList(t *testing.T) {
	list, err := List(&examine.SearchReq{UserName: "luobing", Type: -1, State: -1, Score: -1})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestDetails(t *testing.T) {
	info, err := Details("48C57D8BFC8EE7BEB9ADA36845A6E051")
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(info)
}
