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

//AgentUpdate 资产更新
func AgentUpdate(args agents.UpdateReq) error {

	return agents.Update(args)
}

//SysCheckList 文件列表
func SysCheckList(args agents.SysCheckListReq) (*agents.SysCheckListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.SysCheckList(req, args)
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

//ATTCKESList 安全事件监控列表
func ATTCKESList(args agents.ESListReq) (*agents.ATTCKESHitsListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.ATTCKESList(req, args)
}

//SysCheckESList 文件监控列表
func SysCheckESList(args agents.ESListReq) (*agents.SysCheckESHitsListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.SysCheckESList(req, args)
}

//InvadeThreatESList 入侵威胁列表
func InvadeThreatESList(args agents.ESListReq) (*agents.InvadeThreatESHitsListResp, error) {
	req, err := request.NewRequest()
	if err != nil {
		return nil, err
	}
	return agents.InvadeThreatESList(req, args)
}
