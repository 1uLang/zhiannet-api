package server

import (
	"fmt"
	hids_server "github.com/1uLang/zhiannet-api/hids/model/server"
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

	list, err := List(&hids_server.SearchReq{
		UserName: "LUSIR2",
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestInfo(t *testing.T) {
	info, err := Info("154.91.39.82")

	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(info)
}
