package vulnerabilities

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/vulnerabilities"
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
	info, err := List(&vulnerabilities.ListReq{
		Limit: 20,
		Query: "target_id:b51265ab-9d1c-43c1-8196-a8e36549352f",
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println(info)
	}
}
