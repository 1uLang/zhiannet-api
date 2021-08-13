package agent

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
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
	list, err := List(&agent.SearchReq{UserName: "LUSIR2"})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestDisport(t *testing.T) {
	err := Disport("0A877B041943F39AB6F9F5BFAC797021", "disable")
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}
func TestDownload(t *testing.T) {
	download, err := Download( "Linux64")
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(download)
}
