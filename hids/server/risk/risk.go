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
func SystemDistributed(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.SystemDistributed(req)
}

func WeakList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.WeakList(req)
}
func WeakDetail(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.WeakDetail(req)
}
func ProcessWeak(req *risk.ProcessResp) error {
	return risk.ProcessWeak(req)
}

func DangerAccountList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.DangerAccountList(req)
}
func DangerAccountDetail(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.DangerAccountDetail(req)
}
func ProcessDangerAccount(req *risk.ProcessResp) error {
	return risk.ProcessDangerAccount(req)
}

func ConfigDefectList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.ConfigDefectList(req)
}
func ConfigDefectDetail(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.ConfigDefectDetail(req)
}
func ProcessConfigDefect(req *risk.ProcessResp) error {
	return risk.ProcessConfigDefect(req)
}
