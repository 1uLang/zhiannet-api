package vulnerabilities

import (
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

//List 漏洞列表
//参数：
//	l 显示条数 int
func List(args *ListReq) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Report_vulnerabilities_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}
