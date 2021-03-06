package admin_users

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
	"strings"
)

//资产管理 - 管理用户

const (
	assets_admin_users_path      = "/api/v1/assets/admin-users/"
	assets_admin_users_list_path = "/api/v1/assets/system-users-assets-relations/"
)

func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = assets_admin_users_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	list := make([]map[string]interface{}, 0)

	err = json.Unmarshal(ret, &list)
	if err != nil {
		return nil, err
	}
	users, total, err := GetList(&UserListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return nil, err
	}
	contain := map[string]bool{}
	for _, v := range users {
		contain[v.AdUser] = true
	}
	resList := make([]map[string]interface{}, 0)
	for _, v := range list {
		if _, isExist := contain[v["id"].(string)]; isExist {
			resList = append(resList, v)
		}
	}
	return resList, err
}
func Create(req *request.Request, args *CreateReq) (map[string]interface{}, error) {
	err := args.check()
	if err != nil {
		return nil, fmt.Errorf("参数错误：%v", err)
	}

	req.Method = "post"
	req.Params = model.ToMap(args)
	req.Path = assets_admin_users_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)
	//fmt.Println(info)
	err = json.Unmarshal(ret, &info)
	if err != nil {
		return nil, err
	}
	_, err = AddAdminUser(&AdminUser{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		AdUser:      info["id"].(string),
	})
	//fmt.Println(info, args.UserId, args.AdminUserId)
	return info, err
}
func Delete(req *request.Request, id string) error {

	req.Method = "delete"
	req.Path = assets_admin_users_path + id + "/"
	ret, err := req.Do()
	if err != nil {
		return err
	}
	if len(ret) > 0 {
		//解析返回值
		info := make(map[string]interface{}, 0)
		err = json.Unmarshal(ret, &info)
		return fmt.Errorf(info["detail"].(string))
	} else {
		err = DeleteByAdminUserIds([]string{id})
		return err
	}
}

//Update 修改资产
func Update(req *request.Request, args *UpdateReq) (map[string]interface{}, error) {

	if args.ID == "" {
		return nil, fmt.Errorf("请输入资产id")
	}
	err := args.check()
	if err != nil {
		return nil, err
	}
	req.Method = "put"
	req.Params = model.ToMap(args)
	req.Path = assets_admin_users_path + args.ID + "/"
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)

	err = json.Unmarshal(ret, &info)
	return info, err
}
func AssetsList(req *request.Request, args *AssetsListReq) ([]map[string]interface{}, error) {

	err := args.check()
	if err != nil {
		return nil, err
	}
	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = assets_admin_users_list_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	retInfo := make(map[string]interface{}, 0)
	err = json.Unmarshal(ret, &retInfo)
	if err != nil {
		return nil, err
	}
	ass, total, err := assets.GetList(&assets.AssetsListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return nil, err
	}
	contain := map[string]bool{}
	for _, v := range ass {
		contain[v.AssetsId] = true
	}
	resList := make([]map[string]interface{}, 0)
	for _, v := range retInfo["results"].([]interface{}) {

		result := v.(map[string]interface{})

		if _, isExist := contain[result["asset"].(string)]; isExist {

			tmp := strings.TrimSuffix(result["asset_display"].(string), ")")
			tmps := strings.SplitN(tmp, "(", 2)
			result["asset_name"] = tmps[0]
			result["asset_ip"] = tmps[1]
			resList = append(resList, result)
		}
	}
	return resList, err
}
