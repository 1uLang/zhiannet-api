package targets

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/tidwall/gjson"
	"time"
)

//List 	目标列表
func List(args *ListReq) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "GET"
	req.Url += _const.Targets_api_url
	args.Limit = 999
	args.C = 0
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
	targetList, total, err := GetList(&AddrListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return map[string]interface{}{}, err
	}
	tarMap := map[string]int{}
	for _, v := range targetList {
		tarMap[v.TargetId] = 0
	}
	resList := gjson.ParseBytes(resp)
	list = map[string]interface{}{}
	if resList.Get("targets").Exists() {
		targets := []interface{}{}
		for _, v := range resList.Get("targets").Array() {
			if _, ok := tarMap[v.Get("target_id").String()]; ok {
				targets = append(targets, v.Value())
			}
		}
		list["targets"] = targets
	}

	return list, err
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
	target_id = fmt.Sprintf("%v", ret["target_id"])
	AddAddr(&WebscanAddr{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		TargetId:    target_id,
		CreateTime:  int(time.Now().Unix()),
	})
	return target_id, nil
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
	ret, err := model.ParseResp(resp)
	fmt.Println(ret)
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
