package risk

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
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

func TestRiskList(t *testing.T) {

	list, err := SystemRiskList(&risk.SearchReq{
		//UserName: "LUSIR2",
		Level:        2, //高危
		ProcessState: 1, //未处理
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestVirusList(t *testing.T) {

	list, err := VirusList(&risk.RiskSearchReq{
		//UserName:    "LUSIR2",
		//IsProcessed: false, //待处理
	})
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}

func TestSystemDistributed(t *testing.T) {

	req := &risk.SearchReq{UserName: "luobing", MacCode: "48C57D8BFC8EE7BEB9ADA36845A6E051"}
	list, err := SystemDistributed(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestWeakList(t *testing.T) {

	req := &risk.SearchReq{UserName: "luobing", MacCode: "48C57D8BFC8EE7BEB9ADA36845A6E051"}
	list, err := WeakList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestDangerAccountList(t *testing.T) {

	req := &risk.SearchReq{UserName: "luobing", MacCode: "48C57D8BFC8EE7BEB9ADA36845A6E051"}
	list, err := DangerAccountList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
