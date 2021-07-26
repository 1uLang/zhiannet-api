package assets

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
	"time"
)

//资产管理

const (
	assets_assets_path = "/api/v1/assets/assets/"
)

func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = assets_assets_path
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
	assetList, total, err := GetList(&AssetsListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return []map[string]interface{}{}, err
	}
	contain := map[string]int{}
	for _, v := range assetList {
		contain[v.AssetsId] = 0
	}
	resList := make([]map[string]interface{}, 0)
	for _, v := range list {
		if _, isExist := contain[v["id"].(string)]; isExist {
			resList = append(resList, v)
		}
	}
	return resList, err
}

//Create 创建资产
func Create(req *request.Request, args *CreateReq) (map[string]interface{}, error) {

	err := args.check()
	if err != nil {
		return nil, fmt.Errorf("参数错误：%v", err)
	}
	req.Method = "post"
	req.Params = model.ToMap(args)

	req.Path = assets_assets_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)

	err = json.Unmarshal(ret, &info)
	if err != nil {
		return nil, err
	}
	//写入数据库
	AddAsset(&Asset{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		AssetsId:    info["id"].(string),
		CreateTime:  int(time.Now().Unix()),
	})
	return info, err
}
func Delete(req *request.Request, id string) error {

	if id == "" {
		return fmt.Errorf("资产id不能为空")
	}

	req.Method = "delete"
	req.Path = assets_assets_path + id + "/"
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

		//删除数据库中的数据
		DeleteByAssetIds([]string{id})

		return nil
	}
}

//Update 修改资产
func Update(req *request.Request, args *UpdateReq) (map[string]interface{}, error) {

	err := args.check()
	if err != nil {
		return nil, err
	}
	req.Method = "put"
	req.Params = model.ToMap(args)
	req.Path = assets_assets_path + args.ID + "/"
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)

	err = json.Unmarshal(ret, &info)
	return info, err
}
