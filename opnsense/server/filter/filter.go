package filter

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/1uLang/zhiannet-api/opnsense/request/filter"
	"github.com/1uLang/zhiannet-api/opnsense/server"
)

type (
	FilterReq struct {
		NodeId   uint64 `json:"node_id"`
		Keyword  string `json:"keyword"`
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
	NodeReq struct {
		NodeId uint64 `json:"node_id"`
	}
	UuidReq struct {
		NodeId uint64 `json:"node_id"`
		Uuid   string `json:"uuid"`
	}
	SaveReq struct {
		NodeId uint64      `json:"node_id"`
		Uuid   string      `json:"uuid"`
		Add    bool        `json:"add"`
		Rule   filter.Rule `json:"rule"`
	}
)

//获取日志列表
func GetFilterList(req *FilterReq) (list *filter.FilterListResp, err error) {
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
	return filter.GetFilterList(&filter.FilterReq{
		Current:      fmt.Sprintf("%v", req.PageNum),
		RowCount:     fmt.Sprintf("%v", req.PageSize),
		SearchPhrase: req.Keyword,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//获取详情
func GetFilterInfo(req *UuidReq) (res *filter.FilterInfoResp, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return filter.GetFilterInfo(&filter.UuidReq{
		Uuid: req.Uuid,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//启动停止
func EnableFilter(req *UuidReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return filter.EnableFilter(&filter.UuidReq{
		Uuid: req.Uuid,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//删除规则
func DelFilter(req *UuidReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return filter.DelFilter(&filter.UuidReq{
		Uuid: req.Uuid,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//添加规则
func AddFilter(req *SaveReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return filter.SaveFilter(&filter.SaveReq{
		Add:  true,
		Rule: req.Rule,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}

//编辑规则
func EditFilter(req *SaveReq) (res bool, err error) {
	var loginInfo *request.ApiKey
	loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	if err != nil || loginInfo == nil {
		return res, err
	}
	return filter.SaveFilter(&filter.SaveReq{
		Uuid: req.Uuid,
		Add:  false,
		Rule: req.Rule,
	}, &request.ApiKey{
		Username: loginInfo.Username,
		Password: loginInfo.Password,
		Port:     loginInfo.Port,
		Addr:     loginInfo.Addr,
	})
}
