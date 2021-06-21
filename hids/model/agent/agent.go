package agent

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//agent 管理

//Download agent下载
func Download(username, osType string) (string, error) {

	if osType != "Windows" && osType != "Linux32" && osType != "Linux64" {
		return "", fmt.Errorf("操作系统类型参数错误")
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Ageent_download_api_url, username, osType)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	ret := map[string]interface{}{}
	_, err = model.ParseResp(resp, &ret)
	return ret["agentDownAddress"].(string), err
}

//Install 安装
func Install(username, osType string) (string, error) {
	if osType != "Windows" && osType != "Linux" {
		return "", fmt.Errorf("操作系统类型参数错误")
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Ageent_install_api_url, username, osType)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = nil

	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	ret := map[string]interface{}{}
	_, err = model.ParseResp(resp, &ret)
	return ret["agentInstallCmd"].(string), err
}

//List agent安装主机列表
func List(args *SearchReq) (list SearchResp, err error) {

	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.Ageent_list_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}

//Disport 处理
func Disport(macCode, opt string) error {

	if opt != "disable" && opt != "delete" && opt != "enable" {
		return fmt.Errorf("操作参数错误，该操作参数无效。")
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Agent_dispose_api_uil, opt, macCode)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = nil

	_, err = req.Do()
	return err
}
