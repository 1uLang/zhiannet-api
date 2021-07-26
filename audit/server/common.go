package server

import "github.com/1uLang/zhiannet-api/audit/request"

type (
	//列表下拉参数
	Options struct {
		DictName string `json:"dict_name"`
		Remark   string `json:"remark"`
		Values   []struct {
			IsDefault int    `json:"isDefault"`
			Key       string `json:"key"`
			Remark    string `json:"remark"`
			Value     string `json:"value"`
		} `json:"values"`
	}

	//修改 添加 公共响应参数
	Resp struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}

	//授权 请求参数
	AuthReq struct {
		User  *request.UserReq `json:"user" `
		Email []string         `json:"email"`
		Id    uint64           `json:"id"`
	}

	//修改权限响应参数
	AuthEmailResp struct {
		Code int `json:"code"`
		Data struct {
			Email []string `json:"email"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
)
