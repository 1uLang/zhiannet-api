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

	req := &risk.DetailReq{
		MacCode: "48C57D8BFC8EE7BEB9ADA36845A6E051",
	}

	//req.Req.IsProcessed = true
	//req.Req.State = 1
	list, err := VirusDetailList(req)

	//UserName:    "luobing",
	//IsProcessed: false, //待处理

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
func TestDangerAccountList(t *testing.T) {

	req := &risk.SearchReq{UserName: "luobing"}
	list, err := DangerAccountList(req)
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
func TestConfigDefectDetail(t *testing.T) {

	list, err := ConfigDefectDetail("48C57D8BFC8EE7BEB9ADA36845A6E051", "415", false)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestDangerAccountDetailList(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	req.Req.UserName = "luobing"
	req.Req.ProcessState = 2
	list, err := DangerAccountDetailList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestAbnormalProcessDetail(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	req.Req.UserName = "luobing"
	req.Req.ProcessState = 2
	list, err := AbnormalProcessDetail("BEEBC76C2D8D6C2C9F587A52EF5ACFEF", "395", false)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestWebShellDetailList(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "BEEBC76C2D8D6C2C9F587A52EF5ACFEF"
	req.Req.UserName = "cysct56"
	//req.Req.State = 1
	//req.Req.IsProcessed = true
	list, err := WebShellDetailList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestWebShellList(t *testing.T) {

	req := &risk.RiskSearchReq{}
	req.UserName = "luobing"
	//req.Req.IsProcessed = true
	list, err := WebShellList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestWeakDetailList(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	req.Req.UserName = "luobing"
	req.Req.ProcessState = 2
	list, err := WeakDetailList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}
func TestConfigDefectDetailList(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	req.Req.UserName = "luobing"
	req.Req.ProcessState = 2
	list, err := ConfigDefectDetailList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}

func TestConfigDefectDetailList2(t *testing.T) {

	req := &risk.DetailReq{}
	req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	req.Req.UserName = "luobing"
	req.Req.ProcessState = 1
	list, err := ConfigDefectDetailList(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}

func TestSystemRiskDetail(t *testing.T) {

	list, err := SystemRiskDetail("48C57D8BFC8EE7BEB9ADA36845A6E051", "13255", false)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(list)
}

func TestProcessRisk(t *testing.T) {

	req := &risk.ProcessReq{Opt: "ignore"}
	req.Req.RiskIds = []int{13258}
	req.Req.ItemIds = []string{"4284860"}
	req.Req.MacCode = "48C57D8BFC8EE7BEB9ADA36845A6E051"
	err := ProcessRisk(req)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}
}
