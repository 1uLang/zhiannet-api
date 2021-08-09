package ips

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/ips"
	"github.com/1uLang/zhiannet-api/opnsense/server"
	"github.com/tidwall/gjson"
)

type (
	IpsReq struct {
		NodeId   uint64 `json:"node_id"`
		Keyword  string `json:"keyword"`
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
	EditIpsReq struct {
		NodeId uint64 `json:"node_id"`
		Sid    int64  `json:"sid"`
	}
	DelIpsReq struct {
		NodeId uint64   `json:"node_id"`
		Sid    []string `json:"sid"`
	}
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
	EditActionReq struct {
		NodeId uint64 `json:"node_id"`
		Sid    int64  `json:"sid"`
		Action string `json:"action"`
	}
	IpsAlarmReq struct {
		IpsReq
		FileId string `json:"fileid"`
	}
)

//获取日志列表
func GetIpsList(req *IpsReq) (list *ips.IpsListResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	return ips.GetIpsList(&ips.IpsReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, loginInfo)
}

//ips规则启动 停止
func EditIps(req *EditIpsReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.EditIps(&ips.EditIpsReq{
		Sid: req.Sid,
	}, loginInfo)
}

//ips规则删除
func DelIps(req *DelIpsReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.DelIps(&ips.DelIpsReq{
		Sid: req.Sid,
	}, loginInfo)
}

//应用 使规则生效
func ApplyIps(req *NodeReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.ApplyIps(loginInfo)
}

//drop 或者 alert
func EditAction(req *EditActionReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return res, err
	}
	return ips.EditActionIps(&ips.EditActionIpsReq{
		Sid:    req.Sid,
		Action: req.Action,
	}, loginInfo)
}

//ips-警报
func GetIpsAlarmList(req *IpsAlarmReq) (list *ips.IpsAlarmListResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}
	req2 := &ips.IpsAlarmReq{
		IpsReq: ips.IpsReq{
			Current:      fmt.Sprintf("%v", req.PageNum),
			RowCount:     fmt.Sprintf("%v", req.PageSize),
			SearchPhrase: req.Keyword,
		},
		FileId: req.FileId,
	}
	list, err = ips.GetIpsAlarmList(req2, loginInfo)
	if list != nil && len(list.Rows) > 0 {
		//获取接口名称
		IFace, err := GetIpsAlarmIface(&NodeReq{
			NodeId: req.NodeId,
		})
		if err != nil || IFace == "" {
			IFace = `{"em0":"wan","em1":"lan","lo0":"Loopback"}`
		}
		IFaceName := gjson.Parse(IFace)
		for k, v := range list.Rows {
			list.Rows[k].InIface = IFaceName.Get(v.InIface).String()
		}
	}

	return list, err
}

//报警 时间下拉列表
func GetIpsAlarmTime(req *NodeReq) (list []*ips.IpsAlarmTimeResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return list, err
	}
	//设置请求接口必须的cookie 和 x-csrftoken
	err = request.SetCookie(loginInfo)
	if err != nil {
		return list, err
	}

	return ips.GetIpsAlarmTime(loginInfo)
}

//获取接口名称
func GetIpsAlarmIface(req *NodeReq) (res string, err error) {
	var resp interface{}
	resp, err = cache.CheckCache(
		fmt.Sprintf("opnsense_GetIpsAlarmIface_2_%v", req.NodeId),
		func() (interface{}, error) {
			var res string
			var loginInfo *request.ApiKey
			loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
			if err != nil || loginInfo == nil {
				return res, err
			}
			//设置请求接口必须的cookie 和 x-csrftoken
			err = request.SetCookie(loginInfo)
			if err != nil {
				return res, err
			}
			res, err = ips.GetIpsAlarmIface(loginInfo)
			return res, err
		}, 20, true,
	)
	return fmt.Sprintf("%s", resp), err
}
