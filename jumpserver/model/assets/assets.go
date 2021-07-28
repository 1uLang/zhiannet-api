package assets

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	users_model "github.com/1uLang/zhiannet-api/jumpserver/model/users"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
	"reflect"
	"time"
)

//资产管理

const (
	assets_assets_path = "/api/v1/assets/assets/"
)

func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	nodeid, err := nodeId(req)
	if err != nil {
		return nil, err
	}
	args.NodeId = nodeid
	args.Display = 1

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

func nodeId(req *request.Request) (string, error) {

	req.Method = "get"
	req.Params = nil
	//req.Params = map[string]interface{}{
	//	"assets":0,
	//}

	req.Path = "/api/v1/assets/nodes/children/tree/"
	ret, err := req.Do()
	if err != nil {
		return "", err
	}
	//解析返回值
	info := make([]map[string]interface{}, 0)
	err = json.Unmarshal(ret, &info)
	if err != nil {
		return "", err
	}
	if len(info) == 0 {
		return "", fmt.Errorf("无节点信息")
	}
	for _, node := range info {
		if node["pId"].(string) == "" {
			return node["meta"].(map[string]interface{})["node"].(map[string]interface{})["id"].(string), nil
		}
	}
	return "", nil
}

//Create 创建资产
func Create(req *request.Request, args *CreateReq) (map[string]interface{}, error) {

	nodeid, err := nodeId(req)
	if err != nil {
		return nil, err
	}
	args.Nodes = []string{nodeid}

	err = args.check()
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
	fmt.Println(info)
	if reflect.TypeOf(info["hostname"]).String() != "string" { //报错
		return nil, fmt.Errorf("该主机名已存在")
	}
	//写入数据库
	_, err = AddAsset(&Asset{
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

//Authorize 资产授权
func Authorize(req *request.Request, args *AuthorizeReq) error {

	//列出所有用户
	all, err := users_model.List(req, &users_model.ListReq{})
	if err != nil {
		return err
	}

	allUserMaps := map[string]string{}
	for _, user := range all {
		allUserMaps[user["email"].(string)] = user["username"].(string)
	}

	for _, email := range args.Emails {
		if len(email) == 0 {
			continue
		}
		username, isExist := allUserMaps[email]
		if !isExist {
			return fmt.Errorf("邮箱%v未匹配到用户，请重新输入", email)
		}
		id,err := users_model.GetUserIdByUsername(&users_model.Edgeusers{
			Username: username,
		})
		if err != nil {
			return fmt.Errorf("邮箱%v匹配用户，失败：%v", err)
		}
		//写入数据库
		_, err = AddAsset(&Asset{
			UserId:      id,
			AdminUserId: 0,
			AssetsId:    args.Asset,
			CreateTime:  int(time.Now().Unix()),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func Link(req *request.Request, id string) (string,string,error ){

	req.Path = "/luna/?login_to=" + id
	token ,err := req.GetToken()
	return req.Path,token,err
}