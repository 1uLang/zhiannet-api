package attack_check_server

import (
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/server"
	"testing"
)

func TestHidsAttackCheck(t *testing.T)  {
	InitTestDB()
	//初始化 hidss 服务器地址
	err := server.SetUrl("https://hids.zhiannet.com")
	if err != nil {
		t.Fatal(err)
	}
	//初始化 hdis 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		AppId:  "39rkz",
		Secret: "tkvgpvjuht2625mo",
	})
	ips,err := hids{}.AttackCheck()
	if err != nil {
		t.Fatal(err)
	}

	err = ddos{}.AddBlackIP(ips)
	if err != nil {
		t.Fatal(err)
	}
}
