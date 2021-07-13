package user

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
	"time"
)

//Add 新增用户 先新增机构，在机构下面再创建用户
func Add(args *AddReq) (uint64, error) {

	//先判断用户账号是否已存在
	list, err := List(&SearchReq{UserName: args.UserName})
	if err != nil {
		return 0, fmt.Errorf("获取数据失败：%v", err)
	}
	if list.TotalData > 0 {
		fmt.Println("该账号已存在")
		return 0, nil
	}

	//参数判断
	ok, err := args.Check()
	if err != nil || !ok {
		return 0, fmt.Errorf("参数错误：%v", err)
	}

	req, err := request.NewRequest2()
	if err != nil {
		return 0, err
	}

	req.Method = "post"
	req.Path = _const.Org_api_url
	req.Params = map[string]interface{}{}
	req.Params["signNonce"] = util.RandomNum(10)
	req.Params["data"] = map[string]interface{}{
		"orgOrDept":   args.UserName,
		"userCount":   1,
		"serverCount": 1,
		"wafCount":    1,
		"ygCount":     1,
		//TODO:后续 设置设置改机构的具体到期时间
		"expireTime": time.Now().Add(1 * time.Hour).Format("2006-01-02 15:04:05"),
	}

	resp, err := req.Do2()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	//id, _ := strconv.ParseUint(fmt.Sprintf("%v", ret["orgId"]), 10, 64)

	if ret["returnCode"].(string) != "1"{
		return 0,fmt.Errorf(ret["returnMsg"].(string))
	}

	//ret[orgId] 新增机构id
	fmt.Println(ret)
	args.OrgId, _ = util.Interface2Int(ret["data"].(map[string]interface{})["orgId"])
	fmt.Println(ret["data"],args.OrgId)
	//新增用户
	req.Method = "post"
	req.Path = _const.AddUser_api_url
	req.Params = map[string]interface{}{}
	req.Params["signNonce"] = util.RandomNum(10)
	req.Params["data"] = model.ToMap(args)
	resp, err = req.Do2()
	if err != nil {
		return 0, err
	}

	ret, err = model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	id, _ := util.Interface2Uint64(ret["userId"])
	return id, nil
}

//List 用户列表
func List(args *SearchReq) (list SearchResp, err error) {

	list = SearchResp{}
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
	req.Path = _const.UserList_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}
	_, err = model.ParseResp(resp, &list)
	return list, err
}
