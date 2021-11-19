package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
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
	//req, err := NewServerRequest("http://156.249.24.77:7799", "admin", "admin")
	req, err := NewServerRequest("http://192.168.137.8:8088", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	//err = req.GateWay.Create(&gateway_model.CreateReq{
	//	Name:        "test",
	//	IP:          "127.0.0.1",
	//	Port:        22,
	//	AccountType: "password",
	//	Username:    "root",
	//	Password:    "123456",
	//})
	//fmt.Println(err)
	list, total, err := req.GateWay.List(&gateway_model.ListReq{UserId: 3})
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
