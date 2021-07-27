package user

import (
	"encoding/json"
	_const "github.com/1uLang/zhiannet-api/audit/const"
	"github.com/1uLang/zhiannet-api/audit/request"

	"fmt"
)

type (
	//添加请求参数
	AddUserReq struct {
		User        *request.UserReq `json:"user" `
		Email       string           `json:"email"`
		IsAdmin     int              `json:"is_admin"`
		NickName    string           `json:"nickName"`
		Opt         int              `json:"opt" `
		Password    string           `json:"password"`
		Phonenumber string           `json:"phonenumber"`
		RoleIds     []uint64         `json:"roleIds"`
		Sex         int              `json:"sex"`
		Status      int              `json:"status"`
		UserName    string           `json:"userName"`
	}

	//添加用户响应参数
	AddUserResp struct {
		Code int `json:"code"`
		Data struct {
			Id int `json:"id"`
		} `json:"data"`
		Msg string `json:"msg"`
	}

	RoleListResp struct {
		Code int `json:"code"`
		Data struct {
			RoleList []struct {
				ID         int    `json:"id"`
				Status     int    `json:"status"`
				CreateTime int    `json:"create_time"`
				UpdateTime int    `json:"update_time"`
				ListOrder  int    `json:"list_order"`
				Name       string `json:"name"`
				Remark     string `json:"remark"`
				DataScope  int    `json:"data_scope"`
				UID        int    `json:"uid"`
				CustomerID int
			} `json:"roleList"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
)

//添加用户
func AddUser(req *AddUserReq) (resp *AddUserResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.User)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_USER)
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

//获取角色
func GetRole(req *request.UserReq) (resp *RoleListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req)
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.AUDIT_ADD_USER)
	logReq.QueryParams = req
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return

}
