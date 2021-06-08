package notifications

import (
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

//Notifications 通知消息
func Notifications() (info map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.Notifications_api_url

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}
