package agent

import (
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
func TestDisport(t *testing.T) {
	err := Disport("48C57D8BFC8EE7BEB9ADA36845A6E051", "enable")
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}
