package baseline

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/baseline"
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

	list, err := Detail(&baseline.DetailReq{
		MacCode:  "48C57D8BFC8EE7BEB9ADA36845A6E051",
		PageSize: 161,
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestDetail(t *testing.T) {

	list, err := Detail(&baseline.DetailReq{
		MacCode:  "48C57D8BFC8EE7BEB9ADA36845A6E051",
		PageSize: 1,
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}

func TestTemplateList(t *testing.T) {

	list, err := TemplateList(&baseline.TemplateSearchReq{
		UserName: "luobing",
		PageSize: 100,
		PageNo:   1,
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
