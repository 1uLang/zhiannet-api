package server

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/rules"
	"github.com/1uLang/zhiannet-api/wazuh/request"
)

//RulesInfo 病毒库版本信息
func RulesInfo() (*rules.InfoResp, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return rules.Info(req)
}
