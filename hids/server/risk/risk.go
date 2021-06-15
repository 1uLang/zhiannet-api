package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
)

func RiskList(req *risk.SearchReq) (risk.SearchResp, error) {
	return risk.RiskList(req)
}

func VirusList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.VirusList(req)
}
func Dashboard(userName string) (risk.DashboardResp, error) {
	return risk.Dashboard(userName)
}
func ProcessRisk(req *risk.ProcessResp) error {
	return risk.ProcessRisk(req)
}
