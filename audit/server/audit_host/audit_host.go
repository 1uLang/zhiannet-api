package audit_host

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/audit/const"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server"
	"github.com/1uLang/zhiannet-api/utils"
	"time"
)

type (
	//列表请求参数
	ReqSearch struct {
		User     *request.UserReq `json:"user" `
		Status   string           `json:"status" `
		Name     string           `json:"name" `
		Ip       string           `json:"ip" `
		System   string           `json:"system" `
		PageNum  int              `json:"PageNum"`  //当前页码
		PageSize int              `json:"pageSize"` //每页数
	}
	//列表响应参数
	HostListResp struct {
		Code int `json:"code"`
		Data struct {
			CurrentPage int `json:"currentPage"`
			List        []struct {
				ID         int    `json:"id"`
				UID        int    `json:"uid"`
				Name       string `json:"name"`
				IP         string `json:"ip"`
				System     int    `json:"system"`
				TimeLong   int    `json:"time_long"`
				Status     int    `json:"status"`
				CreateTime int    `json:"create_time"`
				Audits     string `json:"audits"`
				AuditID    string `json:"audit_id"`
				Username   string `json:"username"`
			} `json:"list"`
			Total         int            `json:"total"`
			StatusOptions server.Options `json:"statusOptions"` //数据库状态
			TypeOptions   server.Options `json:"typeOptions"`   //数据库类型
		} `json:"data"`
		Msg string `json:"msg"`
	}

	//添加请求参数
	HostReq struct {
		User     *request.UserReq `json:"user" `
		Uid      uint64           `json:"uid"`
		Name     string           `json:"name"`
		IP       string           `json:"ip" `
		System   uint             `json:"system"`
		Status   uint             `json:"status"`
		TimeLong int              `json:"time_long"`
	}

	//修改请求参数
	HostEditReq struct {
		User   *request.UserReq `json:"user" `
		Id     uint64           `json:"id"`
		Name   string           `json:"name" `
		Status uint             `json:"status" `
	}
	//删除 请求参数
	DelHostReq struct {
		User *request.UserReq `json:"user" `
		Id   uint64           `json:"id"`
	}
	//日志请求参数
	HostLogReq struct {
		UserId    *request.UserReq `json:"user_id" `
		StartTime time.Time        `json:"startTime"`
		EndTime   time.Time        `json:"endTime"`
		Message   string           `json:"message"`
		Page      int64            `json:"pageNum"`
		Size      int64            `json:"pageSize"`
		//ServerHost     string    `json:"serverHost"`
		//StatisticsType string    `json:"statisticsType"`
		AuditId  []string `json:"auditId"`
		TimeType string   `json:"timeType"`
		Sort     string   `json:"sort"`
		//DateHistogram  int64     `json:"date_histogram"`

		ScrollId string `json:"scroll_id"` //连续深度分页 scrollID
		Export   bool   `json:"export"`    //导出文件

	}
	//日志列表响应参数
	HostLogResp struct {
		Code int `json:"code"`
		Data struct {
			Page     int           `json:"page"`
			Log      []interface{} `json:"log"`
			Total    int           `json:"total"`
			Filename string        `json:"filename"` //导出的文件地址

		} `json:"data"`
		Msg string `json:"msg"`
	}
)

//host列表
func GetAuditHostList(req *ReqSearch) (list *HostListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_HOST_LIST)
	logReq.QueryParams = map[string]string{
		"pageSize": fmt.Sprintf("%v", req.PageSize),
		"status":   req.Status,
		"name":     req.Name,
		"ip":       req.Ip,
		"system":   req.System,
		"pageNum":  fmt.Sprintf("%v", req.PageNum),
	}
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &list)
	return

}

//添加主机
func AddHost(req *HostReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_HOST)
	logReq.QueryParams = map[string]string{
		"uid":      fmt.Sprintf("%v", req.Uid),
		"name":     fmt.Sprintf("%v", req.Name),
		"ip":       fmt.Sprintf("%v", req.IP),
		"system":   fmt.Sprintf("%v", req.System),
		"status":   fmt.Sprintf("%v", req.Status),
		"timelong": fmt.Sprintf("%v", req.TimeLong),
		"audits":   "[]",
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

//修改主机
func EditHost(req *HostEditReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EDIT_HOST)
	logReq.QueryParams = map[string]string{
		"id":     fmt.Sprintf("%v", req.Id),
		"name":   fmt.Sprintf("%v", req.Name),
		"status": fmt.Sprintf("%v", req.Status),
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

//删除
func DelHost(req *DelHostReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DEL_HOST)
	logReq.QueryParams = map[string]string{
		"ids": fmt.Sprintf("%v", req.Id),
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

//通过邮箱授权
func AuthHost(req *server.AuthReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_AUTH_EMAIL)
	logReq.QueryParams = map[string]interface{}{
		"emails": req.Email,
		"type":   "1", //0数据库 1主机 2应用
		"value":  fmt.Sprintf("%v", req.Id),
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

//获取邮箱授权列表
func GetAuthEmail(req *server.AuthReq) (resp *server.AuthEmailResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_AUTH_EMAIL_LIST)
	logReq.QueryParams = map[string]string{
		"type":  "1", //0数据库 1主机 2应用
		"value": fmt.Sprintf("%v", req.Id),
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

//日志列表
func GetHostLog(req *HostLogReq) (resp *HostLogResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.UserId)
	if err != nil {
		return
	}
	host := logReq.Addr
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_HOST_LOG_LIST)
	logReq.QueryParams = req
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	if req.Export && resp.Code == 0 {

		resp.Data.Filename = fmt.Sprintf("%v%v", utils.CheckHttpUrl(host, logReq.IsSsl), resp.Data.Filename)
	}
	return
}
