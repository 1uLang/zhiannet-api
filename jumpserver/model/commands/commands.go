package commands

import (
	"encoding/json"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/model/assets"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

//命令记录

const (
	terminal_commands_path = "/api/v1/terminal/commands/"
)

//会话列表
func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = terminal_commands_path
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
	if args.Asset == "" {
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
			if _, isExist := contain[v["asset"].(string)]; isExist {
				resList = append(resList, v)
			}
		}
		return resList, err
	} else {
		return list, err
	}

}
