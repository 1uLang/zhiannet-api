package scans

import (
	"github.com/1uLang/zhiannet-api/hids/util"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/request"
	"time"
)

//扫描目标模板

const (
	scan_templates_url = "/editor/scan/templates"
	scan_url = "/scans"
)

func getScanTemplateUUid() (string, error) {

	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}

	req.Method = "get"
	req.Url += scan_templates_url
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return "", err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return "", err
	}
	var id string
	templates := ret["templates"].([]map[string]interface{})
	for _, template := range templates {
		if template["title"] == "Host Discovery" { //主机漏洞扫描 模板
			id = template["uuid"].(string)
		}
	}
	return id, nil
}

//创建目标
func Create(args *AddReq) (uint64, error) {

	req, err := request.NewRequest()
	if err != nil {
		return 0, err
	}

	req.Method = "post"
	req.Url += scan_url
	req.Params = model.ToMap(*args)

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	id, _ := util.Interface2Uint64(ret["id"])
	//写入数据库
	_, err = AddScans(&NessusScans{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		ScansId:    id,
		CreateTime:  int(time.Now().Unix()),
	})
	return id, nil
}
