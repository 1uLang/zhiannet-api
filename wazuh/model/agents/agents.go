package agents

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/model"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"strings"
)

const (
	agent_api_url                = "/agents"
	agent_summary_status_api_url = "/agents/summary/status?pretty=true"
	agent_scan_url               = "/syscheck"
	agent_sca_url                = "/sca/"
	agent_sca_details_url        = "/sca/%s/checks/%s"
	agent_syscheck_url           = "/syscheck/"
	agent_ciscat_url             = "/ciscat/%s/results"
	agent_vulnerability_url      = "/vulnerability/%s"
)

func Statistics(req *request.Request) (*StatisticsResp, error) {

	req.Method = "get"
	req.Path = agent_summary_status_api_url
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	list := &StatisticsResp{}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}
func Delete(req *request.Request, ids []string) error {
	req.Method = "delete"
	req.Path = agent_api_url
	req.Params = map[string]interface{}{
		"agents_list": strings.Join(ids, ","),
		"status":      "all",
		"older_than":  "0s",
		"pretty":      true,
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	list := &StatisticsResp{}
	fmt.Println(resp)
	if resp.Error != 0 {
		return fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return err
}
func List(req *request.Request, args *ListReq) (*ListResp, error) {
	req.Method = "get"
	req.Path = agent_api_url
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &ListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}
func Scan(req *request.Request, agent []string) error {
	req.Method = "put"
	req.Path = agent_scan_url
	req.Params = map[string]interface{}{
		"agents_list": strings.Join(agent, ","),
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Error != 0 {
		return fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	fmt.Println(resp.Data)
	return err
}

//SCADetailsList 合规基线详情
func SCADetailsList(req *request.Request, agent, policy string) (*SCADetailsListResp, error) {
	req.Method = "get"
	req.Path = fmt.Sprintf(agent_sca_details_url, agent, policy)
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SCADetailsListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

//SCAList Security configuration assessment 合规基线
func SCAList(req *request.Request, agent string) (*SCAListResp, error) {

	req.Method = "get"
	req.Path = agent_sca_url + agent
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SCAListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func SysCheckList(req *request.Request, agent string) (*SysCheckListResp, error) {

	req.Method = "get"
	req.Path = agent_syscheck_url + agent
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SysCheckListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func CiscatList(req *request.Request, agent string) (*CiscatListResp, error) {

	req.Method = "get"
	req.Path = fmt.Sprintf(agent_ciscat_url, agent)
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &CiscatListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func VulnerabilityList(req *request.Request, agent string) (*VulnerabilityListResp, error) {
	req.Method = "get"
	req.Path = fmt.Sprintf(agent_vulnerability_url, agent)
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &VulnerabilityListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}
