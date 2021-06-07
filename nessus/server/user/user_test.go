package user

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/nessus/model/user"
	"github.com/1uLang/zhiannet-api/nessus/request"
	"github.com/1uLang/zhiannet-api/nessus/server"
	"testing"
)

func init() {
	//初始化 nessus 服务器地址
	err := server.SetUrl("https://182.150.0.108:8834/")
	if err != nil {
		panic(err)
	}
	//初始化 nessus 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		Access: "d15f22095fe5fb6f1a81a574fa13163cc1c4a0b596de244453588aaf8a057129",
		Secret: "6ad76937e2e9c39a5e9463b67f43423fbfafffd8095c834b9f70a43aad9da591",
	})
	if err != nil {
		panic(err)
	}
}

func TestAddUser(t *testing.T) {

	err := Add(&user.AddReq{
		Username:    "test12332",
		Password:    "test",
		Permissions: "128",
		Name:        "test",
		Email:       "2439716@qq.com",
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {

	err := Delete("12")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {

	err := Update(&user.UpdateReq{
		ID:          "13",
		Permissions: "32",
		Name:        "test",
		Email:       "2439716@qq.com",
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestChangePasswordUser(t *testing.T) {

	err := ChangePassword(&user.ChangePasswordReq{
		ID:              "13",
		CurrentPassword: "test",
		Password:        "test2",
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
func TestEnabled(t *testing.T) {
	err := Enabled(&user.EnableReq{
		ID:      "13",
		Enabled: true,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestAPIKeys(t *testing.T) {
	accessKey, secretKey, err := APIKeys("13")
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println("get user api keys : accessKey[", accessKey, "],secretKey[", secretKey, "]")
	}
}
