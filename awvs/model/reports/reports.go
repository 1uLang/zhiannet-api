package reports

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	"github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
)

//List 报表列表
func List(limit int) (list map[string]interface{}, err error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}

	req.Method = "get"
	req.Url = _const.Awvs_server + _const.Reports_api_url
	req.Params = map[string]interface{}{
		"l": limit,
	}

	resp, err := req.Do()
	if err != nil {
		return nil, err
	}
	return model.ParseResp(resp)
}

//Delete 删除报表
func Delete(report_id string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}

	req.Method = "DELETE"
	req.Url = _const.Awvs_server + _const.Reports_api_url + "/" + report_id

	resp, err := req.Do()
	if err != nil {
		return err
	}
	ret, err := model.ParseResp(resp)
	fmt.Println(ret)
	return err
}
