package agent

import "github.com/1uLang/zhiannet-api/hids/model/agent"

func Create(args *agent.CreateReq) (err error) {
	return agent.Create(args)
}

func Download(osType string) (string, error) {
	return agent.Download(osType)
}
func Install(osType string) (string, error) {
	return agent.Install(osType)
}
func List(req *agent.SearchReq) (agent.SearchResp, error) {
	return agent.List(req)
}
func Disport(macCode, opt string) error {
	return agent.Disport(macCode, opt)
}
