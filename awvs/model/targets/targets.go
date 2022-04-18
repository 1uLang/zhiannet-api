package targets

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/tidwall/gjson"
	"time"
)

func Search(args *ListReq) (list []interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "GET"
	req.Url += _const.Targets_api_url
	args.Limit = 100
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	respInfo, err := model.ParseResp(resp)
	if err != nil {
		return nil, err
	}

	if targets, ok := respInfo["targets"]; ok {
		return targets.([]interface{}), nil
	} else {
		return nil, nil
	}

}

//List 	目标列表

func List(args *ListReq) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "GET"
	req.Url += _const.Targets_api_url
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
	//获取数据库 当前用户的扫描用户
	//targetList, total, err := GetList(&AddrListReq{
	//	UserId:      args.UserId,
	//	AdminUserId: args.AdminUserId,
	//	PageSize:    999,
	//	PageNum:     1,
	//})
	//if total == 0 || err != nil {
	//	return map[string]interface{}{}, err
	//}
	tarMap := map[string]uint64{}
	//for _, v := range targetList {
	//	tarMap[v.TargetId] = v.Id
	//}
	resList := gjson.ParseBytes(resp)
	list = map[string]interface{}{}
	targets := []interface{}{}
	if resList.Get("targets").Exists() {
		for _, v := range resList.Get("targets").Array() {
			if dbid, ok := tarMap[v.Get("target_id").String()]; ok {
				item := v.Value().(map[string]interface{})
				item["id"] = dbid
				targets = append(targets, item)
			}
		}
	}
doNext:
	if resList.Get("pagination").Exists() {
		if resList.Get("pagination").Get("count").Int() > int64(args.Limit+args.C) {
			req.Url += _const.Targets_api_url
			args.Limit = 100
			args.C += args.Limit
			req.Params = model.ToMap(args)

			resp, err := req.Do()
			if err != nil {
				return nil, err
			}
			resList = gjson.ParseBytes(resp)
			list = map[string]interface{}{}
			if resList.Get("targets").Exists() {
				for _, v := range resList.Get("targets").Array() {
					if dbid, ok := tarMap[v.Get("target_id").String()]; ok {
						item := v.Value().(map[string]interface{})
						item["id"] = dbid
						targets = append(targets, item)
					}
				}
			}
			goto doNext
		}
	}
	list["targets"] = targets

	return list, err
}

//Add 添加目标
func Add(args *AddReq) (target_id string, err error) {

	if args.C { //检测 该addr 是否已添加 是着 直接返回
		list, err := List(&ListReq{Limit: 100, Query: "text_search:*" + args.Address + ";", UserId: args.UserId, AdminUserId: args.AdminUserId})
		if err != nil {
			fmt.Println(err)
		} else {
			if len(list) > 0 {
				return list["targets"].([]interface{})[0].(map[string]interface{})["target_id"].(string), nil
			}
		}
	}

	ok, err := args.Check()
	if err != nil || !ok {
		return "", fmt.Errorf("参数错误：%v", err)
	}
	req, err := request.NewRequest()
	if err != nil {
		return "", err
	}

	req.Method = "POST"
	req.Url += _const.Targets_api_url
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	ret, err := model.ParseResp(resp)
	if err != nil {
		return "", err
	}
	//TODO 暂不处理入库失败
	//入库
	config, _ := json.Marshal(GetConfigResp{Username: args.Username, Password: args.Password})
	target_id = fmt.Sprintf("%v", ret["target_id"])
	//设置 登录设置
	_ = siteLogin(target_id, args.Username, args.Password)
	_, _ = AddAddr(&WebscanAddr{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		TargetId:    target_id,
		Config:      config,
		CreateTime:  int(time.Now().Unix()),
	})
	return target_id, nil
}

//登录设置
func siteLogin(target_id, username, password string) error {

	kind := "automatic"
	if username == "" || password == "" || target_id == "" {
		kind = "none"
	}
	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "PATCH"
	req.Url += "/api/v1/targets/" + target_id + "/configuration"
	req.Params = map[string]interface{}{
		"login": map[string]interface{}{
			"credentials": map[string]interface{}{
				"enabled":  true,
				"password": password,
				"username": username,
			},
			"kind": kind,
		},
		"sensor": false,
		"ssh_credentials": map[string]interface{}{
			"kind": "none",
		},
	}

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

//Delete 删除目标
func Delete(target_id string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "DELETE"
	req.Url += _const.Targets_api_url + "/" + target_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)

	if err == nil {
		//删除数据库中的数据
		DeleteByTargetIds([]string{target_id})
	}
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
	req.Url += _const.Targets_api_url + "/" + target_id
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
	req.Url += _const.Targets_api_url + "/" + target_id + "/configuration"
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return err
	}
	_, err = model.ParseResp(resp)
	return err
}
