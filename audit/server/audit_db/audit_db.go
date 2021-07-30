package audit_db

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
		Type     string           `json:"type" `
		PageNum  int              `json:"PageNum"`  //当前页码
		PageSize int              `json:"pageSize"` //每页数
	}
	//列表响应参数
	DbListResp struct {
		Code int `json:"code"`
		Data struct {
			CurrentPage int `json:"currentPage"`
			List        []struct {
				ID         int    `json:"id"`
				UID        int    `json:"uid"`
				Name       string `json:"name"`
				Type       int    `json:"type"`
				Version    string `json:"version"`
				IP         string `json:"ip"`
				Port       string `json:"port"`
				System     int    `json:"system"`
				TimeLong   int    `json:"time_long"`
				Status     int    `json:"status"`
				CreateTime int    `json:"create_time"`
				AuditID    string `json:"audit_id"`
				Username   string `json:"username"`
			} `json:"list"`
			Total         int            `json:"total"`
			StatusOptions server.Options `json:"statusOptions"` //数据库状态
			TypeOptions   server.Options `json:"typeOptions"`   //数据库类型
		} `json:"data"`
		Msg string `json:"msg"`
	}

	//添加数据库请求参数
	DBReq struct {
		User *request.UserReq `json:"user" `
		//Uid      uint64           `json:"uid"`
		Type     uint   `json:"type"`
		Name     string `json:"name"`
		Version  string `json:"version"`
		IP       string `json:"ip" `
		Port     string `json:"port" `
		System   uint   `json:"system"`
		Status   uint   `json:"status"`
		TimeLong int    `json:"time_long"`
	}

	//修改数据库请求参数
	DBEditReq struct {
		User   *request.UserReq `json:"user" `
		Id     uint64           `json:"id"`
		Name   string           `json:"name" `
		Status uint             `json:"status" `
	}
	//删除 请求参数
	DelDbReq struct {
		User *request.UserReq `json:"user" `
		Id   uint64           `json:"id"`
	}
	DbLogReq struct {
		UserId    *request.UserReq `json:"user_id" `
		StartTime time.Time        `json:"startTime"`
		EndTime   time.Time        `json:"endTime"`
		Message   string           `json:"message"`
		Page      int64            `json:"pageNum"`
		Size      int64            `json:"pageSize"`
		//ServerHost     string           `json:"serverHost"`
		ClientHost string `json:"clientHost"`
		//SqlType        string           `json:"sqlType"`
		User   string `json:"user"`
		DbName string `json:"dbName"`
		//StatisticsType string           `json:"statisticsType"`
		TimeType string   `json:"timeType" `
		Risk     string   `json:"risk"` //风险语句 true?false
		AuditId  []string `json:"auditId"`
		Sort     string   `json:"sort"`

		ScrollId string `json:"scroll_id"` //连续深度分页 scrollID
		Export   bool   `json:"export"`    //导出文件
	}

	DbLogResp struct {
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

//数据库列表
func GetAuditBdList(req *ReqSearch) (list *DbListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DB_LIST)
	logReq.QueryParams = map[string]string{
		"pageSize": fmt.Sprintf("%v", req.PageSize),
		"status":   req.Status,
		"name":     req.Name,
		"ip":       req.Ip,
		"type":     req.Type,
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

//添加数据库
func AddDb(req *DBReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_DB)
	logReq.QueryParams = map[string]string{
		//"uid":      fmt.Sprintf("%v", req.Uid),
		"type":     fmt.Sprintf("%v", req.Type),
		"name":     fmt.Sprintf("%v", req.Name),
		"version":  fmt.Sprintf("%v", req.Version),
		"ip":       fmt.Sprintf("%v", req.IP),
		"port":     fmt.Sprintf("%v", req.Port),
		"system":   fmt.Sprintf("%v", req.System),
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

//修改数据库
func EditDb(req *DBEditReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EDIT_DB)
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
func DelDb(req *DelDbReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DEL_DB)
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
func AuthDb(req *server.AuthReq) (resp *server.Resp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_AUTH_EMAIL)
	logReq.QueryParams = map[string]interface{}{
		"emails": req.Email,
		"type":   "0", //0数据库 1主机 2应用
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
		"type":  "0", //0数据库 1主机 2应用
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
func GetDbLog(req *DbLogReq) (resp *DbLogResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.UserId)
	if err != nil {
		return
	}
	host := logReq.Addr
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_DB_LOG_LIST)
	//logReq.QueryParams = map[string]interface{}{
	//	"startTime":  req.StartTime,
	//	"endTime":    req.EndTime,
	//	"message":    req.Message,
	//	"pageNum":    req.Page,
	//	"pageSize":   req.Size,
	//	"clientHost": req.ClientHost,
	//	"user":       req.User,
	//	"dbName":     req.DbName,
	//	"timeType":   req.TimeType,
	//	"risk":       req.Risk,
	//	"auditId":    req.AuditId,
	//	"sort":       "desc",
	//	"scroll_id":  req.ScrollId,
	//}
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
