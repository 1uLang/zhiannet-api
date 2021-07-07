package scans

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/scans"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/awvs/server"
	"testing"
)

func init() {
	//初始化 awvs 服务器地址
	err := server.SetUrl("https://scan-web.zhiannet.com")
	if err != nil {
		panic(err)
	}
	//初始化 awvs 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		XAuth: "1986ad8c0a5b3df4d7028d5f3c06e936c429ffb1149e2491b84fe51cc63a6b26a",
	})
	if err != nil {
		panic(err)
	}
}

func TestList(t *testing.T) {
	info, err := List(&scans.ListReq{
		Limit: 10,
		C:     1,
		//UserId: 1,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println(info)
	}
}

func TestVulnerabilities(t *testing.T) {
	info, err := Vulnerabilities(&scans.VulnerabilitiesReq{
		ScanId:        "86a50d76-5ef4-40e6-a88c-44086d1f9a7e",
		ScanSessionId: "8158bb38-c1a5-4f41-b956-7fc5fb984c40",
		VulId:         "2612438631037535816",
	})

	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println(info)
	}
}
