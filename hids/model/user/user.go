package user

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
	"github.com/1uLang/zhiannet-api/nessus/model"
	"strconv"
	"time"
)

func Add() (uint64, error) {

	//参数判断
	//ok, err := args.Check()
	//if err != nil || !ok {
	//	return 0, fmt.Errorf("参数错误：%v", err)
	//}

	req, err := request.NewRequest()
	if err != nil {
		return 0, err
	}

	req.Method = "post"
	req.Url += _const.Org_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = map[string]interface{}{
		"orgOrDept":   "test12312",
		"userCount":   100,
		"serverCount": 100,
		"wafCount":    1,
		"ygCount":     1,
		"expireTime":  time.Now().Add(1 * time.Hour).Format("2006-01-02 15:04:05"),
		"description": "test",
	}

	resp, err := req.Do()
	if err != nil {
		return 0, err
	}

	ret, err := model.ParseResp(resp)
	if err != nil {
		return 0, err
	}
	fmt.Println(ret)
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", ret["orgId"]), 10, 64)
	return id, nil
}
