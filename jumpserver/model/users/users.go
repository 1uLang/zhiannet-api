package users

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/jumpserver/model"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
	"time"
)

const (
	users_path = "/api/v1/users/users/"
)

//users 用户层

//List 用户列表
func List(req *request.Request, args *ListReq) ([]map[string]interface{}, error) {

	req.Method = "get"
	req.Params = model.ToMap(args)
	req.Path = users_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	list := make([]map[string]interface{}, 0)

	err = json.Unmarshal(ret, &list)
	return list, err
}
func Create(req *request.Request, args *CreateReq) (map[string]interface{}, error) {

	if err := args.check(); err != nil {
		return nil, fmt.Errorf("参数错误：%v", err)
	}

	req.Method = "post"

	//设置过期时间
	if args.DateExpired == "" {
		//60年
		args.DateExpired = time.Now().AddDate(60, 0, 0).Format("2006-01-02 15:04:05")
	}
	req.Params = model.ToMap(args)
	req.Path = users_path
	ret, err := req.Do()
	if err != nil {
		return nil, err
	}
	//解析返回值
	info := make(map[string]interface{}, 0)

	err = json.Unmarshal(ret, &info)
	return info, err
}