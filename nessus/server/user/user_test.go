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
	err := server.SetUrl("https://156.240.95.239:8834")
	if err != nil {
		panic(err)
	}
	//初始化 nessus 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		Access: "4caa54e1df36556950450ee00d5e6e22b55fdcb81d940e3999c51c743782288c",
		Secret: "c96b82f6507e2249647ef5e50f32642504fbaad37d5cfee28f58c630545f9ebd",
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
