package email

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/audit/const"
	"github.com/1uLang/zhiannet-api/audit/request"
)

type (
	//获取邮箱配置请求参数
	EmailInfoReq struct {
		User *request.UserReq `json:"user" `
	}
	EmailInfoRes struct {
		Code int    `json:"code"`
		Data Data   `json:"data"`
		Msg  string `json:"msg"`
	}

	Info struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		CustomerStatus int    `json:"customer_status"`
		CreateTime     int    `json:"create_time"`
		Status         int    `json:"status"`
		Host           string `json:"host"`
		Port           int    `json:"port"`
		Username       string `json:"username"`
		Password       string `json:"password"`
	}
	Data struct {
		Info Info `json:"info"`
	}

	//保存配置参数
	SetEmailReq struct {
		User     *request.UserReq `json:"user" `
		Id       uint64           `json:"id"`
		Host     string           `json:"host"  `
		Port     uint             `json:"port"  `
		Username string           `json:"username"  `
		Password string           `json:"password"  `
		To       string           `json:"to"`
	}

	EmailRes struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
)

func GetEmail(req *EmailInfoReq) (resp *EmailInfoRes, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EMAIL_INFO)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return

}

func SetEmail(req *SetEmailReq) (resp *EmailRes, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EMAIL_EDIT)
	logReq.QueryParams = map[string]string{
		//"uid":      fmt.Sprintf("%v", req.Uid),
		"host":     fmt.Sprintf("%v", req.Host),
		"username": fmt.Sprintf("%v", req.Username),
		"password": fmt.Sprintf("%v", req.Password),
		"port":     fmt.Sprintf("%v", req.Port),
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

func CheckEmail(req *SetEmailReq) (resp *EmailRes, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_EMAIL_CHECK)
	logReq.QueryParams = map[string]string{
		//"uid":      fmt.Sprintf("%v", req.Uid),
		"host":     fmt.Sprintf("%v", req.Host),
		"username": fmt.Sprintf("%v", req.Username),
		"password": fmt.Sprintf("%v", req.Password),
		"port":     fmt.Sprintf("%v", req.Port),
		"to":       fmt.Sprintf("%v", req.To),
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
