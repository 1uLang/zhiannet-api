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
func Download( osType string) (string, error) {

	if osType != "Windows" && osType != "Linux32" && osType != "Linux64" {
		return "", fmt.Errorf("操作系统类型参数错误")
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Ageent_download_api_url, model.HidsUserNameAPI, osType)
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
func Install( osType string) (string, error) {
	if osType != "Windows" && osType != "Linux" {
		return "", fmt.Errorf("操作系统类型参数错误")
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}
	req.Method = "get"
	req.Path = fmt.Sprintf(_const.Ageent_install_api_url, model.HidsUserNameAPI, osType)
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
	agentList := make([]map[string]interface{},0)
	agents,total ,err := GetList(&ListReq{UserId: args.UserId,AdminUserId: args.AdminUserId})
	if err != nil || total == 0{
		return list,err
	}
	contain := map[string]int{}
	for k,v := range agents{
		contain[v.IP] = k
		agentList = append(agentList, map[string]interface{}{"serverIp":v.IP})
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.Ageent_list_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	args.UserName = model.HidsUserNameAPI
	args.PageSize = 10
	args.PageNo = 1
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}

	_, err = model.ParseResp(resp, &list)

	for _,item := range list.List{
		if idx,isExist := contain[item["serverIp"].(string)];isExist{
			agentList[idx] = item
		}
	}
	list.List = agentList
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
	req.Method = "post"
	req.Path = fmt.Sprintf(_const.Agent_dispose_api_uil, macCode, opt)
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = map[string]interface{}{
		"opt":     opt,
		"macCode": macCode,
	}

	resp, err := req.Do()
	fmt.Println(string(resp))
	return err
}

//Create 新增agent ip 用户输入
func Create(args *CreateReq) (err error) {

	if args.AgentIp == "" || args.UserId == 0 && args.AdminUserId == 0 {
		return fmt.Errorf("参数错误")
	}
	return addAgent(&hidsAgents{IP: args.AgentIp, UserId: args.UserId, AdminUserId: args.AdminUserId})
}
func Update(args *UpdateReq) (err error) {
	if args.Id == 0 || args.AgentIp == "" || args.UserId == 0 && args.AdminUserId == 0 {
		return fmt.Errorf("参数错误")
	}
	return updateAgent(&hidsAgents{Id: args.Id, IP: args.AgentIp, UserId: args.UserId, AdminUserId: args.AdminUserId})
}
func Delete(args *DeleteReq)error  {
	if args.Id == 0 {
		return fmt.Errorf("参数错误")
	}
	return deleteAgent(args.Id)
}
