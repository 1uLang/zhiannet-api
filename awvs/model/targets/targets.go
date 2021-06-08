package targets

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

//List 	目标列表
func List(args *ListReq) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "GET"
	req.Url = _const.Awvs_server + _const.Targets_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Add 添加目标
func Add(args *AddReq) (target_id string, err error) {

	ok, err := args.Check()
	if err != nil || !ok {
		return "", fmt.Errorf("参数错误：%v", err)
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}

	req.Method = "POST"
	req.Url = _const.Awvs_server + _const.Targets_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	ret, err := model.ParseResp(resp)
	if err != nil {
		return "", err
	}
	return ret["target_id"].(string), nil
}

//Delete 删除目标
func Delete(target_id string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "DELETE"
	req.Url = _const.Awvs_server + _const.Targets_api_url + "/" + target_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	ret, err := model.ParseResp(resp)
	fmt.Println(ret)
	return err
}

//Update 修改目标信息
func Update(target_id string, args *UpdateReq) error {

	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "PATCH"
	req.Url = _const.Awvs_server + _const.Targets_api_url + "/" + target_id
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}

//SetLogin 目标登录设置
func SetLogin(target_id string, args *SetLoginReq) error {

	if target_id == "" {
		return fmt.Errorf("目标id不能为空")
	}
	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "PATCH"
	req.Url = _const.Awvs_server + _const.Targets_api_url + "/" + target_id + "/configuration"
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}
