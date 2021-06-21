package vulnerabilities

import "github.com/1uLang/zhiannet-api/awvs/model/vulnerabilities"

//List 目标列表
func List(req *vulnerabilities.ListReq) (info map[string]interface{}, err error) {
	return vulnerabilities.List(req)
}

//Details 单个漏洞详情
func Details(vul_id string) (info map[string]interface{}, err error) {
	return vulnerabilities.Details(vul_id)
}
