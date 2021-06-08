package user

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/user"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
	"testing"
)

func init() {
	//初始化 nessus 服务器地址
	err := server.SetUrl("https://user.cloudhids.net/")
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
func TestAdd(t *testing.T) {
	orgId, err := Add(&user.AddReq{
		UserName: "LUSIR2",
		Password: "2323",
		Role:     3,
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(orgId)
}
func TestList(t *testing.T) {

	list, err := List(&user.SearchReq{
		UserName: "LUSIR2",
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
