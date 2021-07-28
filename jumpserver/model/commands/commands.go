package commands

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

//命令记录

const (
	terminal_commands_path = "/api/v1/terminal/commands/"
)

func getCommandStorageId(req *request.Request, args *ListReq) (string, error) {

	req.Method = "get"
	req.Path = "/api/v1/terminal/command-storages/tree/"
	ret, err := req.Do()
	if err != nil {
		return "", err
	}
	//解析返回值
	list := make([]map[string]interface{}, 0)
	err = json.Unmarshal(ret, &list)
	if err != nil {
		return "", err
	}
	if len(list) > 0 {
		return list[0]["id"].(string), nil
	}
	return "", fmt.Errorf("堡垒机部分参数为初始化")
}

//会话列表
func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	var err error
	req.Params = model.ToMap(args)
	req.Params["command_storage_id"], err = getCommandStorageId(req, args)
	if err != nil {
		return nil, err
	}
	req.Method = "get"
	req.Path = terminal_commands_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}

	//解析返回值
	list := make([]map[string]interface{}, 0)

	_ = json.Unmarshal(ret, &list)

	return list, err

}
