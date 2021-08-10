package audit_app

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
		AppType  string           `json:"appType" `
		PageNum  int              `json:"PageNum"`  //当前页码
		PageSize int              `json:"pageSize"` //每页数
	}
	//列表响应参数
	AppListResp struct {
		Code int `json:"code"`
		Data struct {
			CurrentPage int `json:"currentPage"`
			List        []struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				IP         string `json:"ip"`
				AppType    int    `json:"app_type"`
				Status     int    `json:"status"`
				User       int    `json:"user"`
				CreateTime int    `json:"create_time"`
				TimeLong   int    `json:"time_long"`
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
	AppReq struct {
		User *request.UserReq `json:"user" `
		//Uid      uint64           `json:"uid"`
		Name     string `json:"name"`
		IP       string `json:"ip" `
		AppType  uint   `json:"app_type"`
		Status   uint   `json:"status"`
		TimeLong int    `json:"timelong"`
	}

	//修改请求参数
	AppEditReq struct {
		User   *request.UserReq `json:"user" `
		Id     uint64           `json:"id"`
		Name   string           `json:"name" `
		Status uint             `json:"status" `
	}
	//删除 请求参数
	DelAppReq struct {
		User *request.UserReq `json:"user" `
		Id   uint64           `json:"id"`
	}

	//日志请求参数
	AppLogReq struct {
		UserId    *request.UserReq `json:"user_id" `
		StartTime time.Time        `p:"startTime"`
		EndTime   time.Time        `p:"endTime"`
		Message   string           `p:"message"`
		Page      int64            `p:"pageNum"`
		Size      int64            `p:"pageSize"`
		//ServerHost     string    `p:"serverHost"`
		//StatisticsType string    `p:"statisticsType"`
		TimeType string   `p:"timeType"`
		AuditId  []string `p:"auditId"`
		Sort     string   `p:"sort"`
		//DateHistogram  int64     `p:"date_histogram"`

		Export   bool   `json:"export"`    //导出文件
		ScrollId string `json:"scroll_id"` //连续深度分页 scrollID

	}
	//日志列表响应参数
	AppLogResp struct {
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

//应用列表
func GetAuditAppList(req *ReqSearch) (list *AppListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_APP_LIST)
	logReq.QueryParams = map[string]string{
		"pageSize": fmt.Sprintf("%v", req.PageSize),
		"status":   req.Status,
		"name":     req.Name,
		"ip":       req.Ip,
		"appType":  req.AppType,
		"pageNum":  fmt.Sprintf("%v", req.PageNum),
	}
	logReq.ReqType = "post"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &list)
	return

}

//添加应用
func AddApp(req *AppReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_APP)
	logReq.QueryParams = map[string]string{
		//"uid":      fmt.Sprintf("%v", req.Uid),
		"name":     fmt.Sprintf("%v", req.Name),
		"ip":       fmt.Sprintf("%v", req.IP),
		"appType":  fmt.Sprintf("%v", req.AppType),
		"status":   fmt.Sprintf("%v", req.Status),
		"timelong": fmt.Sprintf("%v", req.TimeLong),
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

//修改应用
func EditApp(req *AppEditReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EDIT_APP)
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
func DelApp(req *DelAppReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DEL_APP)
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
func AuthApp(req *server.AuthReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_AUTH_EMAIL)
	logReq.QueryParams = map[string]interface{}{
		"emails": req.Email,
		"type":   "2", //0数据库 1主机 2应用
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
	logReq.QueryParams = map[string]interface{}{
		"emails": req.Email,
		"ids":    req.Ids,
		"type":   "2", //0数据库 1主机 2应用
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

//日志列表
func GetAppLog(req *AppLogReq) (resp *AppLogResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.UserId)
	if err != nil {
		return
	}
	host := logReq.Addr
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_APP_LOG_LIST)
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
