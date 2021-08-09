package server

import (
	_const "github.com/1uLang/zhiannet-api/hids/const"
	"github.com/1uLang/zhiannet-api/hids/model"
	"github.com/1uLang/zhiannet-api/hids/model/agent"
	"github.com/1uLang/zhiannet-api/hids/request"
	"github.com/1uLang/zhiannet-api/hids/util"
)

//List 主机列表
func List(args *SearchReq) (list SearchResp, err error) {
	agentList := make([]map[string]interface{}, 0)
	agents, total, err := agent.GetList(&agent.ListReq{UserId: args.UserId, AdminUserId: args.AdminUserId})
	if err != nil || total == 0 {
		return list, err
	}
	contain := map[string]int{}
	for k, v := range agents {
		contain[v.IP] = k
	}

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
	args.UserName = model.HidsUserNameAPI
	req.Params = model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return list, err
	}

	_, err = model.ParseResp(resp, &list)

	for _, item := range list.ServerInfoList {
		if _, isExist := contain[item["serverIp"].(string)]; !isExist {
			continue
		}
		agentList = append(agentList, item)
	}
	list.TotalData = len(agentList)
	list.ServerInfoList = agentList
	return list, err
}

//Info 主机信息
func Info(serverIp string) (info map[string]interface{}, err error) {
	list, err := List(&SearchReq{ServerIp: serverIp, UserName: model.HidsUserNameAPI})
	if err != nil {
		return info, err
	}
	if len(list.ServerInfoList) == 0 {
		return nil, nil
	}
	return list.ServerInfoList[0], nil
}
