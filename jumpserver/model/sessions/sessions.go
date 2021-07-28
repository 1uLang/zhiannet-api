package sessions

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

//终端会话管理

const (
	terminal_sessions_path = "/api/v1/terminal/sessions/"
)

//会话列表
func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = terminal_sessions_path
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
	assetList, total, err := assets.GetList(&assets.AssetsListReq{
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
		if _, isExist := contain[v["asset_id"].(string)]; isExist {
			resList = append(resList, v)
		}
	}
	return resList, err
	//return list, nil
}

//监控
func Monitor(req *request.Request, id string) error {

	req.Method = "post"
	req.Path = terminal_sessions_path + id + "/replay/"
	ret, err := req.Do()
	if err != nil {
		return err
	}
	//解析返回值
	resp := map[string]interface{}{}
	err = json.Unmarshal(ret, &resp)
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return err
}

//回放
func Replay(req *request.Request, id string) (string, string, error) {

	req.Path = terminal_sessions_path + id + "/replay/"
	token, err := req.GetToken()
	return req.Path, token, err
}

//回放下载
func Download(req *request.Request, id string) (string, string, error) {

	req.Path = terminal_sessions_path + id + "/replay/download/"
	token, err := req.GetToken()
	return req.Path, token, err
}
