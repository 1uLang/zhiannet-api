package agent

import "github.com/1uLang/zhiannet-api/hids/model/agent"

func Download(username, osType string) (string, error) {
	return agent.Download(username, osType)
}
func Install(username, osType string) (string, error) {
	return agent.Install(username, osType)
}
func List(req *agent.SearchReq) (agent.SearchResp, error) {
	return agent.List(req)
}
func Disport(macCode, opt string) error {
	return agent.Disport(macCode, opt)
}
