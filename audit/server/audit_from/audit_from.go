package audit_from

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/audit/const"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server"
)

type (
	//列表请求参数
	ReqSearch struct {
		User       *request.UserReq `json:"user" `
		Name       string           `json:"name" `
		AssetsType string           `json:"Assets_type" `
		PageNum    int              `json:"PageNum"`  //当前页码
		PageSize   int              `json:"pageSize"` //每页数
	}
	//列表响应参数
	FromListResp struct {
		Code int `json:"code"`
		Data struct {
			CurrentPage int `json:"currentPage"`
			List        []struct {
				ID         int    `json:"id"`
				UserID     int    `json:"user_id"`
				Name       string `json:"name"`
				Cycle      int    `json:"cycle"`
				SendTime   string `json:"send_time"`
				Format     int    `json:"format"`
				AssetsType int    `json:"assets_type"`
				AssetsID   int    `json:"assets_id"`
				AssetsName string `json:"assets_name"`
				Email      string `json:"email"`
				IsDelete   int    `json:"is_delete"`
				CreateTime int    `json:"create_time"`
				CycleDay   int    `json:"cycle_day"`
				NextTime   int    `json:"next_time"`
			} `json:"list"`
			Total int `json:"total"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	//添加请求参数
	FromReq struct {
		User       *request.UserReq `json:"user" `
		Id         uint64           `json:"id"`
		UserId     uint64           `json:"user_id"`
		Name       string           `json:"name" v:"required#请填写名称"`
		Cycle      int              `json:"cycle" `
		CycleDay   int              `json:"cycle_day" `
		SendTime   string           `json:"send_time" `
		Format     int              `json:"format"  `
		AssetsType int              `json:"assets_type" `
		AssetsId   uint64           `json:"assets_id" `
		Email      string           `json:"email" `
	}

	//删除 请求参数
	DelFromReq struct {
		User *request.UserReq `json:"user" `
		Id   uint64           `json:"id"`
	}
	//获取详情
	GetFromReq struct {
		User *request.UserReq `json:"user" `
		Id   uint64           `json:"id"`
	}

	//获取详情响应参数
	FromResp struct {
		Code int `json:"code"`
		Data struct {
			Info struct {
				ID         int    `json:"id"`
				UserID     int    `json:"user_id"`
				Name       string `json:"name"`
				Cycle      int    `json:"cycle"`
				SendTime   string `json:"send_time"`
				Format     int    `json:"format"`
				AssetsType int    `json:"assets_type"`
				AssetsID   int    `json:"assets_id"`
				Email      string `json:"email"`
				IsDelete   int    `json:"is_delete"`
				CreateTime int    `json:"create_time"`
				CycleDay   int    `json:"cycle_day"`
				NextTime   int    `json:"next_time"`
			} `json:"info"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
)

//列表
func GetAuditFromList(req *ReqSearch) (list *FromListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_FROM_LIST)
	logReq.QueryParams = map[string]string{
		"pageSize":    fmt.Sprintf("%v", req.PageSize),
		"name":        req.Name,
		"assets_type": req.AssetsType,
		"pageNum":     fmt.Sprintf("%v", req.PageNum),
	}
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &list)
	return

}

//添加
func AddFrom(req *FromReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_FROM)
	logReq.QueryParams = req
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return

}

//修改
func EditFrom(req *FromReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EDIT_FROM)
	logReq.QueryParams = req
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//删除
func DelFrom(req *DelFromReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DEL_FROM)
	logReq.QueryParams = map[string]string{
		"id": fmt.Sprintf("%v", req.Id),
	}
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}

//获取详情
func GetFrom(req *GetFromReq) (resp *FromResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_GET_FROM)
	logReq.QueryParams = map[string]string{
		"id": fmt.Sprintf("%v", req.Id),
	}
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
