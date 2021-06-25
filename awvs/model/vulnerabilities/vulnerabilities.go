package vulnerabilities

import (
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/tidwall/gjson"
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
	list, err = model.ParseResp(resp)
	if err != nil {
		return map[string]interface{}{}, err
	}
	if args.UserId == 0 && args.AdminUserId == 0 {
		return list, err
	}
	//获取当前用户的targets ID
	vList, total, err := targets.GetList(&targets.AddrListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})

	if total == 0 || err != nil {
		return map[string]interface{}{}, err
	}
	tarMap := map[string]int{}
	for _, v := range vList {
		tarMap[v.TargetId] = 0
	}
	resList := gjson.ParseBytes(resp)
	list = map[string]interface{}{}
	if resList.Get("vulnerabilities").Exists() {
		targets := []gjson.Result{}
		for _, v := range resList.Get("vulnerabilities").Array() {
			if _, ok := tarMap[v.Get("target_id").String()]; ok {
				targets = append(targets, v)
			}
		}
		list["vulnerabilities"] = targets
	}
	return list, err
}
func Details(vul_id string) (info map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Report_vulnerabilities_api_url + "/" + vul_id

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}
