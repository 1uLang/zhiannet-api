package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	asset_model "github.com/1uLang/zhiannet-api/next-terminal/model/asset"
	session_model "github.com/1uLang/zhiannet-api/next-terminal/model/session"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
	InitTable()
}
func Test_conn(t *testing.T) {
	ls := new(CheckRequest)
	ls.Run()
}
func TestGateway_List(t *testing.T) {
	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	//req, err := NewServerRequest("http://192.168.137.8:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	err = req.GateWay.Create(&gateway_model.CreateReq{
		Name:        "test",
		IP:          "127.0.0.1",
		Port:        22,
		AccountType: "password",
		Username:    "root",
		Password:    "123456",
	})
	list, total, err := req.GateWay.List(&gateway_model.ListReq{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list, total)
}
func TestGateway_Delete(t *testing.T) {
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	err = req.GateWay.Delete("4bd020d5-829c-41fc-b1e4-421d6b0b481b")
	if err != nil {
		t.Fatal(err)
	}
}
func TestGateway_Update(t *testing.T) {
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	args := gateway_model.UpdateReq{
		Id: "b0f5d5fc-458e-4c2e-a888-9de3d3ccd28f",
	}
	args.Name = "1231231"
	args.IP = "127.0.0.1"
	args.Port = 22
	args.AccountType = "password"
	args.Username = "root"
	args.Password = "123456"
	err = req.GateWay.Update(&args)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGateway_Authorize(t *testing.T) {
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	args := gateway_model.AuthorizeReq{
		Id:      "b0f5d5fc-458e-4c2e-a888-9de3d3ccd28f",
		UserId:  1,
		UserIds: []uint64{2, 3},
	}
	err = req.GateWay.Authorize(&args)

	if err != nil {
		t.Fatal(err)
	}
}

func TestGateway_ReConnect(t *testing.T) {
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	err = req.GateWay.ReConnect("b0f5d5fc-458e-4c2e-a888-9de3d3ccd28f")

	if err != nil {
		t.Fatal(err)
	}
}
func TestGateway_AuthorizeUserList(t *testing.T) {
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	usrs, err := req.GateWay.AuthorizeUserList("b0f5d5fc-458e-4c2e-a888-9de3d3ccd28f")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(usrs, err)
}

func TestGateway_Info(t *testing.T) {
	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	info, err := req.GateWay.Info("cf4cc148-c528-43ef-a67c-0359c7ff5532")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info, err)
}

func TestSession_Replay(t *testing.T) {
	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	info, err := req.Session.Replay(&session_model.ReplayReq{Id: "39f28274-aed6-41c7-8bb3-120b02eba0e2"})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info, err)
}

func TestAsset_Create(t *testing.T) {

	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	err = req.Assets.Create(&asset_model.CreateReq{
		AccountType: "custom",
		IP:          "1.1.1.1",
		Name:        "test",
		Password:    "1231231",
		Port:        22,
		Protocol:    "ssh",
		Username:    "root",
		UserId:      1,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestAsset_List(t *testing.T) {

	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	list, _, err := req.Assets.List(&asset_model.ListReq{
		UserId: 1,
	})

	if err != nil {
		t.Fatal(err)
	}
	for _, v := range list {
		fmt.Println(v)
	}
}
func TestAsset_Connect(t *testing.T) {

	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	url, err := req.Assets.Connect(&asset_model.ConnectReq{
		Id: "166129b2-85f2-42bd-8131-c705dc7274f2",
	})

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}

func TestAsset_Update(t *testing.T) {

	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	args := &asset_model.UpdateReq{
		Id: "daf8004c-46bc-4bce-963b-183234c8b013",
	}
	args.AccountType = "custom"
	args.IP = "1.1.1.1"
	args.Name = "test"
	args.Password = "1231231"
	args.Port = 22
	args.Protocol = "ssh"
	args.Username = "root"
	args.UserId = 1
	err = req.Assets.Update(args)

	if err != nil {
		t.Fatal(err)
	}
}

func TestAsset_Authorize(t *testing.T) {

	req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	args := &asset_model.AuthorizeReq{
		AssetId: "7fb127d7-bc27-4ac9-b145-4b9e45679caa",
		UserId:  1,
		UserIds: []uint64{2, 3},
	}
	err = req.Assets.Authorize(args)

	if err != nil {
		t.Fatal(err)
	}
}
