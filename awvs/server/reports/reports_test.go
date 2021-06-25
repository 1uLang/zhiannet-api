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

func TestCreate(t *testing.T) {
	req := &reports.CreateResp{
		Source: struct {
			IDS  []string `json:"id_list"`
			Type string   `json:"list_type"`
		}{IDS: []string{"d15b7205-fb69-49d0-8743-e03440d1828f"}, Type: "scans"},
		TemplateId: "11111111-1111-1111-1111-111111111112", //快速
	}
	info, err := Create(req)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		fmt.Println(info)
	}
}
