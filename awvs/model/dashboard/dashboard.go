package dashboard

import (
	"github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

//MeStats 获取仪表板信息
func MeStats() (info map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url += _const.MeStats_api_url

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)

}
