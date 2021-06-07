package user

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/nessus/const"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"github.com/1uLang/zhiannet-api/nessus/request"
	"strconv"
)

/*
Add 新增用户
	参数：
	permissions : 权限 int 16/24/32/40/64   标准32
	32 	具有此角色的用户可以创建扫描、策略和用户目标组。
	128	具有此角色的用户具有与标准用户相同的权限，但还可以管理用户、组、代理、资产数据导出、漏洞数据导出、排除、系统目标组、用户目标组、访问组和扫描程序。
*/
func Add(args *AddReq) (uint64, error) {

	//参数判断
	ok, err := args.Check()
	if err != nil || !ok {
		return 0, fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return 0, err
	}

	req.Method = "post"
	req.Url += _const.Users_api_url
	req.Params = model.ToMap(*args)

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", ret["id"]), 10, 64)
	return id, nil
}

/*
Delete 删除用户
参数：
	id 用户id
*/
func Delete(id string) error {
	//参数判断
	if id == "" {
		return fmt.Errorf("参数错误：用户id不能为空")
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	req.Method = "delete"
	req.Url += _const.Users_api_url + "/" + id
	req.Params = nil
	resp, err := req.Do()
	if err != nil {
		panic(err)
	}
	_, err = model.ParseResp(resp)
	return err
}

/*
Update 更新用户信息
参数：
	id 用户id
	permissions	权限	string(int)
	name		名称	string
	email		邮箱	string
	enabled		账号启动/禁用	bool
*/
func Update(args *UpdateReq) error {

	//参数判断
	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "put"
	req.Url += _const.Users_api_url + "/" + args.ID
	req.Params = map[string]interface{}{
		"permissions": args.Permissions,
		"name":        args.Name,
		"email":       args.Email,
		"enabled":     args.Enabled,
	}

	resp, err := req.Do()
	if err != nil {
		panic(err)
	}

	_, err = model.ParseResp(resp)
	return err
}

/*
ChangePassword 修改密码用户
参数：
	id 用户id
	current_password 当前密码
	password 修改密码
*/
func ChangePassword(args *ChangePasswordReq) error {

	//参数判断
	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "put"
	req.Url += _const.Users_api_url + "/" + args.ID + "/chpasswd"
	req.Params = map[string]interface{}{
		"current_password": args.CurrentPassword,
		"password":         args.Password,
	}
	resp, err := req.Do()
	if err != nil {
		panic(err)
	}

	_, err = model.ParseResp(resp)
	return err
}

/*
Enable 禁用启用用户
参数：
	id 用户id
	enabled 启用、禁用 TRUE、FALSE
当前账号无权限
*/
func Enable(args *EnableReq) error {

	ok, err := args.Check()
	if err != nil || !ok {
		return fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "put"
	req.Url += _const.Users_api_url + "/" + args.ID + "/enabled"
	req.Params = map[string]interface{}{
		"enabled": args.Enabled,
	}
	resp, err := req.Do()
	if err != nil {
		panic(err)
	}

	_, err = model.ParseResp(resp)
	return err
}

/*
APIKeys 过去用户APIKeys
参数：
	id 用户id
*/
func APIKeys(id string) (accessKey, secretKey string, err error) {

	//参数检测
	if id == "" {
		return "", "", fmt.Errorf("参数错误：用户id不能为空")
	}

	req, err := request.NewRequest()
	if err != nil {
		return "", "", err
	}

	req.Method = "put"
	req.Url += _const.Users_api_url + "/" + id + "/keys"

	resp, err := req.Do()
	if err != nil {
		panic(err)
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return "", "", err
	}

	return ret["accessKey"].(string), ret["secretKey"].(string), nil
}
