package assets

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	users_model "github.com/1uLang/zhiannet-api/jumpserver/model/users"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
	"reflect"
	"strings"
	"time"
)

//资产管理

const (
	assets_assets_path      = "/api/v1/assets/assets/"
	assets_permissions_path = "/api/v1/perms/asset-permissions/"
	//更新硬件信息 - 测试链接性
	//assets_tasks_path = "/api/v1/assets/assets/0bf4e168-d3b4-443d-a2a8-c135f451ece5/tasks/"
	//测试链接性
	//assets_check_link_path = "/ui/#/ops/celery/task/ecc601e9-152b-4f45-8af4-bff1c1935581/log"
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

	nodeid, err := nodeId(req)
	if err != nil {
		return nil, err
	}
	args.Nodes = []string{nodeid}

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

type authorizeReq struct {
	Actions     []string `json:"actions"`
	Assets      []string `json:"assets"`
	DateExpired string   `json:"date_expired"`
	DateStart   string   `json:"date_start"`
	IsActive    bool     `json:"is_active"`
	Name        string   `json:"name"`
	Nodes       []string `json:"nodes"`
	SystemUsers []string `json:"system_users"`
	Users       []string `json:"users"`
}

func systemUsers(req *request.Request) ([]string, error) {

	req.Path = "/api/v1/assets/system-users/"
	req.Method = "get"
	req.Params = map[string]interface{}{
		"offset":  0,
		"limit":   99,
		"display": 1,
		"draw":    1,
	}
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	retInfo := map[string]interface{}{}
	err = json.Unmarshal(resp, &retInfo)
	if err != nil {
		return nil, err
	}
	ids := []string{}
	for _, item := range retInfo["results"].([]interface{}) {
		ids = append(ids, item.(map[string]interface{})["id"].(string))
	}
	fmt.Println(ids)
	if len(ids) == 0 {
		return nil, fmt.Errorf("未初始化管理用户。")
	}
	return ids, nil
}
func authorize(req *request.Request, asset, touser string, node, sysUsers []string) error {

	users, err := users_model.List(req, &users_model.ListReq{Name: touser})
	if err != nil {
		return err
	}
	fmt.Println(users)
	if len(users) == 0 {

		return fmt.Errorf("资产授权失败 %s 该用户不存在", touser)
	}

	req.Path = assets_permissions_path
	req.Method = "post"
	args := &authorizeReq{}
	args.Assets = []string{asset}
	args.Actions = []string{"all", "connect", "updownload", "upload_file", "download_file"}
	args.IsActive = true
	args.Name = asset + "-" + touser
	layout := "2006-01-02T15:04:05"
	args.DateStart = time.Now().Format(layout) + ".481Z"
	args.DateExpired = time.Now().AddDate(10, 0, 0).Format(layout) + ".481Z"
	args.Nodes = node
	args.SystemUsers = sysUsers

	args.Users = []string{users[0]["id"].(string)}
	fmt.Println(args)
	ret, err := req.Do()
	if err != nil {
		return err
	}

	fmt.Println("do resp : ",string(ret))
	return nil
}

//Authorize 资产授权
func Authorize(req *request.Request, args *AuthorizeReq) error {

	node, err := nodeId(req)
	if err != nil {
		return err
	}
	sysUsers, err := systemUsers(req)
	if err != nil {
		return err
	}

	//列出所有用户
	all, err := users_model.List(req, &users_model.ListReq{})
	if err != nil {
		return err
	}

	allUserMaps := map[string]string{}
	for _, user := range all {
		allUserMaps[user["email"].(string)] = user["username"].(string)
	}
	fmt.Println(args.Emails)
	for _, email := range args.Emails {
		fmt.Println(email)
		if len(email) == 0 {
			continue
		}
		username, isExist := allUserMaps[email]
		if !isExist {
			return fmt.Errorf("邮箱%v未匹配到用户，请重新输入", email)
		}
		fmt.Println(username)
		ent, err := users_model.GetUserInfoByUsername(&users_model.Edgeusers{
			Username: username,
		})
		if err != nil {
			return fmt.Errorf("邮箱%v匹配用户，失败：%v", email, err)
		}
		fmt.Println(ent)
		//授权
		err = authorize(req, args.Asset, username, []string{node}, sysUsers)
		if err != nil {
			return err
		}
		//写入数据库
		_, err = AddAsset(&Asset{
			UserId:      ent.Id,
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

func DelAuthorize(req *request.Request, args *DelAuthorizeReq) error {

	req.Method = "delete"
	req.Path = assets_permissions_path + args.Asset + "/"
	req.Params = nil
	_, err := req.Do()
	return err
}
func AuthorizeList(req *request.Request, args *AuthorizeListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Path = assets_permissions_path
	req.Params = map[string]interface{}{
		"asset_id": args.Asset,
	}
	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}
	ret := make([]map[string]interface{}, 0)
	for _, v := range result {
		name := v["name"].(string)
		if !strings.Contains(name, "-") {
			continue
		}
		tmps := strings.Split(name, "-")
		if len(tmps) != 2 {
			continue
		}
		username := tmps[1]
		ent, err := users_model.GetUserInfoByUsername(&users_model.Edgeusers{Username: username})
		if err != nil {
			fmt.Println("查询用户名：", username, " 失败：", err)
			return nil, err
		}
		ret = append(ret, map[string]interface{}{
			"username": username,
			"name":     ent.Name,
			"email":    ent.Email,
			"id":       v["id"],
		})
	}
	return ret, err
}
func Link(req *request.Request, id string) (string, string, error) {

	req.Path = "/luna/?login_to=" + id
	token, err := req.GetToken()
	return req.Path, token, err
}

//Info 资产详情
func Info(req *request.Request, id string) (map[string]interface{}, error) {

	req.Method = "get"

	req.Path = assets_assets_path + id + "/"
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

	return info, err
}

//更新硬件信息
func task(req *request.Request, id, method string) (string, error) {

	req.Method = "post"
	req.Path = assets_assets_path + id + "/tasks/"
	req.Params = map[string]interface{}{
		"action": method,
	}
	ret, err := req.Do()
	if err != nil {
		return "", err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)
	err = json.Unmarshal(ret, &info)
	if err != nil {
		return "", err
	}
	fmt.Println(info)
	return "/ui/#/ops/celery/task/" + info["task"].(string) + "/log", nil
}
func Refresh(req *request.Request, id string) (string, error) {

	return task(req, id, "refresh")
}

//测试资产可连接性
func CheckLink(req *request.Request, id string) (string, error) {

	return task(req, id, "test")
}
