package scans

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
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
	req.Params = map[string]interface{}{
		"l": args.Limit,
		"c": args.C,
	}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Add 创建扫描
func Add(args *AddReq) (target_id string, err error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return "", fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}

	req.Method = "post"
	req.Url = _const.Awvs_server + _const.Scans_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	info, err := model.ParseResp(resp)
	if err != nil {
		return "", err
	}
	return info["target_id"].(string), nil
}

//ScanningProfiles 扫描配置文件列表
func ScanningProfiles() (list map[string]interface{}, err error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url = _const.Awvs_server + _const.Scanning_profiles_api_url

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
	req.Url = _const.Awvs_server + _const.Report_templates_api_url

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
	req.Url = _const.Awvs_server + _const.Scans_api_url + "/" + target_id

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
	req.Url = _const.Awvs_server + _const.Scans_api_url + "/" + scan_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	ret, err := model.ParseResp(resp)
	fmt.Println(ret)
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
	req.Url = _const.Awvs_server + _const.Scans_api_url + "/" + scan_id

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
	req.Url = _const.Awvs_server + _const.Scans_api_url + "/" + scan_id + "/results/" + scan_session_id + "/statistics"

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}
