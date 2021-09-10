package server

import (
	"github.com/1uLang/zhiannet-api/wazuh/model/agents"
	"github.com/1uLang/zhiannet-api/wazuh/request"
)

//AgentList 资产列表
func AgentList(args *agents.ListReq) (*agents.ListResp, error) {

	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.List(req, args)
}

//AgentDelete 资产删除
func AgentDelete(ids []string) error {

	req, err := request.NewRequest()
	if err != nil {
		return err
	}
	return agents.Delete(req, ids)
}

//VulnerabilityList 漏洞风险列表
func VulnerabilityList(args agents.ESListReq) (*agents.VulnerabilityHitsResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.VulnerabilityESList(req, args)
}

//VirusList 病毒管理列表
func VirusList(args agents.ESListReq) (*agents.VirusESHitsListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.VirusESList(req, args)
}

//BaselineList 合规基线列表
func BaselineList(args agents.SCAListReq) (*agents.SCAListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.SCAList(req, args)
}

//BaselineDetailsList 合规基线详情列表
func BaselineDetailsList(args agents.SCADetailsListReq) (*agents.SCADetailsListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.SCADetailsList(req, args)
}
