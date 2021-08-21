package scans

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/model/targets"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/tidwall/gjson"
)

//List 扫描列表
//参数：
//	l 显示条数 int
func List(args *ListReq) (list map[string]interface{}, err error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url
	args.Limit = 100
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	list, err = model.ParseResp(resp)
	if err != nil {
		return nil, err
	}
	if args.UserId == 0 && args.AdminUserId == 0 {
		return list, err
	}
	//获取当前用户的targets ID
	scanList, total, err := targets.GetList(&targets.AddrListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})

	if total == 0 || err != nil {
		return map[string]interface{}{}, err
	}
	tarMap := map[string]int{}
	for _, v := range scanList {
		tarMap[v.TargetId] = 0
	}
	resList := gjson.ParseBytes(resp)
	list = map[string]interface{}{}
	if resList.Get("scans").Exists() {
		targets := []interface{}{}
		for _, v := range resList.Get("scans").Array() {
			if _, ok := tarMap[v.Get("target_id").String()]; ok {
				targets = append(targets, v.Value())
			}
		}
		list["scans"] = targets
	}
	return list, err
}

//Add 创建扫描
func Add(args *AddReq) (err error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "post"
	req.Url += _const.Scans_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}

//ScanningProfiles 扫描配置文件列表
func ScanningProfiles() (list map[string]interface{}, err error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scanning_profiles_api_url

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//ReportTemplates 报表列表
func ReportTemplates() (list map[string]interface{}, err error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Report_templates_api_url

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//GetScansId 获取目标的扫描id
func GetScansId(target_id string) (map[string]interface{}, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url + "/" + target_id

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Delete 删除扫描
func Delete(scan_id string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "DELETE"
	req.Url += _const.Scans_api_url + "/" + scan_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)

	return err
}

//GetInfo 获取单个扫描状态
func GetInfo(scan_id string) (info map[string]interface{}, err error) {
	if scan_id == "" {
		return nil, fmt.Errorf("扫描id不能为空")
	}
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url + "/" + scan_id

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Statistics 单个扫描概况信息
func Statistics(scan_id, scan_session_id string) (info map[string]interface{}, err error) {
	if scan_id == "" {
		return nil, fmt.Errorf("扫描id不能为空")
	}
	if scan_session_id == "" {
		return nil, fmt.Errorf("扫描会话id不能为空")
	}
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url + "/" + scan_id + "/results/" + scan_session_id + "/statistics"

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Abort 停止扫描
func Abort(scan_id string) error {
	if scan_id == "" {
		return fmt.Errorf("扫描id不能为空")
	}
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "post"
	req.Url += _const.Scans_api_url + "/" + scan_id + "/abort"

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}

func Vulnerabilities(args *VulnerabilitiesReq) (map[string]interface{}, error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return nil, fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url + "/" + args.ScanId + "/results/" + args.ScanSessionId + "/vulnerabilities/" + args.VulId
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}
func VulnerabilitiesList(args *VulnerabilitiesListReq) ([]interface{}, error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return nil, fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Scans_api_url + "/" + args.ScanId + "/results/" + args.ScanSessionId + "/vulnerabilities?l=100"
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	ret,err:= model.ParseResp(resp)
	if err != nil {
		return nil, err
	}
	return ret["vulnerabilities"].([]interface{}),nil
}