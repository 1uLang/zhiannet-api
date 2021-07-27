package server

import (
	"fmt"
	assets_model "github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	"testing"
)

func init() {
	InitMysql()
	var err error
	//info, err := GetFortCloud()
	//if err != nil {
	//	panic(err)
	//}
	req, err = NewServerRequest("http://182.150.0.106:8080", "lusir", "dengbao-lusir")
	//req, err = NewServerRequest(info.Addr, info.Username, info.Password)
	if err != nil {
		panic(err)
	}
}
func TestGetFortCloud(t *testing.T) {
	var err error
	info, err := GetFortCloud()
	if err != nil {
		panic(err)
	}
	req, err = NewServerRequest(info.Addr, info.Username, info.Password)
	fmt.Println(info)
	if err != nil {
		panic(err)
	}
}

func TestAssetsList(t *testing.T) {
	args := &assets_model.ListReq{}
	list, err := req.Assets.List(args)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	fmt.Println(list)
}
func TestCreate(t *testing.T) {
	args := &assets_model.CreateReq{}
	args.HostName = "182.150.0.100"
	args.IP = "182.150.0.100"
	args.Platform = "Linux"
	args.Protocols = []string{"ssh/22"}
	args.Active = true
	args.AdminUser = "e6957684-7e25-4c1a-af60-5329e4398032"
	args.Comment = "182.150.0.100"
	args.PublicIp = "182.150.0.100"
	info, err := req.Assets.Create(args)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	fmt.Println(info)
}
func TestUpdate(t *testing.T) {
	args := &assets_model.UpdateReq{}
	args.ID = "e8dd287a-f4e0-4804-8bfe-6df5ef9678fe"
	args.HostName = "1"
	args.IP = "1"
	args.Platform = "Linux"
	args.Protocols = []string{"ssh/22"}
	args.Active = true
	//req.AdminUser = "1ef6fc7d-ca98-4fa1-8a21-800ae58c48ef"
	args.Comment = "创建主机测试"
	args.PublicIp = "1"
	info, err := req.Assets.Update(args)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	fmt.Println(info)
}
func TestDelete(t *testing.T) {
	err := req.Assets.Delete("5eb846ce-7343-4ad7-9ed6-d637ddf7cd4f")
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
}
