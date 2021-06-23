package server

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//List 主机列表
func List(args *SearchReq) (list SearchResp, err error) {

	list = SearchResp{}
	if args.PageSize == 0 {
		args.PageSize = 10
	}
	if args.PageNo == 0 {
		args.PageNo = 1
	}

	req, err := request.NewRequest()
	if err != nil {
		return list, err
	}
	req.Method = "get"
	req.Path = _const.ServerList_api_url
	req.Headers["signNonce"] = util.RandomNum(10)
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}

	_, err = model.ParseResp(resp, &list)
	return list, err
}

//Info 主机信息
func Info(serverIp string) (info map[string]interface{}, err error) {
	list, err := List(&SearchReq{ServerIp: serverIp})
	if err != nil {
		return info, err
	}
	if len(list.ServerInfoList) == 0 {
		return info, fmt.Errorf("无该主机信息")
	}
	return list.ServerInfoList[0], nil
}
