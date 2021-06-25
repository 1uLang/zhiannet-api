package reports

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
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
	info, err := List(&reports.ListReq{
		Limit: 1,
		//C:     1,
		UserId: 1,
	})
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println(info)
	}
}

func TestCreate(t *testing.T) {
	res, err := Create(&reports.CreateResp{
		Source: struct {
			IDS  []string `json:"id_list"`
			Type string   `json:"list_type"`
		}{IDS: []string{"33c38a7f-9759-4202-87a5-b0e7cc6b5d0d", "dc7d58eb-6bec-4a75-8418-8c41036b9481"}, Type: "scans"},
		TemplateId: "11111111-1111-1111-1111-111111111112", //快速
	})
	fmt.Println(res)
	fmt.Println(err)
}
