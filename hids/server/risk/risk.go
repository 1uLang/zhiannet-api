package risk

import (
	"github.com/1uLang/zhiannet-api/hids/model/risk"
)

func SystemRiskList(req *risk.SearchReq) (risk.SearchResp, error) {
	return risk.SystemRiskList(req)
}
func SystemRiskDetail(macCode, riskId string, state bool) (info map[string]interface{}, err error) {
	return risk.SystemRiskDetail(macCode, riskId, state)
}
func Dashboard(userName string) (risk.DashboardResp, error) {
	return risk.Dashboard(userName)
}

//漏洞风险

func ProcessRisk(req *risk.ProcessReq) error {
	return risk.ProcessRisk(req)
}
func SystemDistributed(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.SystemDistributed(req)
}

func WeakList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.WeakList(req)
}
func WeakDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.WeakDetailList(req)
}
func ProcessWeak(req *risk.ProcessReq) error {
	return risk.ProcessWeak(req)
}

func DangerAccountList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.DangerAccountList(req)
}
func DangerAccountDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.DangerAccountDetailList(req)
}
func ProcessDangerAccount(req *risk.ProcessReq) error {
	return risk.ProcessDangerAccount(req)
}

func ConfigDefectList(req *risk.SearchReq) (info risk.SystemDistributedResp, err error) {
	return risk.ConfigDefectList(req)
}
func ConfigDefectDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.ConfigDefectDetailList(req)
}
func ProcessConfigDefect(req *risk.ProcessReq) error {
	return risk.ProcessConfigDefect(req)
}

//入侵威胁

func VirusList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.VirusList(req)
}
func VirusDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.VirusDetail(macCode, id)
}
func VirusDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.VirusDetailList(req)
}
func ProcessVirus(req *risk.ProcessReq) error {
	return risk.VirusProcess(req)
}

func WebShellList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.WebShellList(req)
}
func WebShellDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.WebShellDetail(macCode, id)
}
func WebShellDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.WebShellDetailList(req)
}
func ProcessWebShell(req *risk.ProcessReq) error {
	return risk.WebShellProcess(req)
}

func ReboundShellList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.ReboundList(req)
}
func ReboundShellDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.ReboundShellDetail(macCode, id)
}
func ReboundShellDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.ReboundDetailList(req)
}
func ProcessReboundShell(req *risk.ProcessReq) error {
	return risk.ReboundShellProcess(req)
}
func AbnormalAccountList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.AbnormalAccountList(req)
}
func AbnormalAccountDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.AbnormalAccountDetail(macCode, id)
}
func AbnormalAccountDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.AbnormalAccountDetailList(req)
}
func ProcessAbnormalAccount(req *risk.ProcessReq) error {
	return risk.AbnormalAccountProcess(req)
}

func LogDeleteList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.LogDeleteList(req)
}
func LogDeleteDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.LogDeleteDetail(macCode, id)
}
func LogDeleteDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.LogDeleteDetailList(req)
}
func ProcessLogDelete(req *risk.ProcessReq) error {
	return risk.LogDeleteProcess(req)
}

func AbnormalLoginList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.AbnormalLoginList(req)
}
func AbnormalLoginDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.AbnormalLoginDetail(macCode, id)
}
func AbnormalLoginDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.AbnormalLoginDetailList(req)
}
func ProcessAbnormalLogin(req *risk.ProcessReq) error {
	return risk.AbnormalLoginProcess(req)
}

func AbnormalProcessList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.AbnormalProcessList(req)
}
func AbnormalProcessDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.AbnormalProcessDetail(macCode, id)
}
func AbnormalProcessDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.AbnormalProcessDetailList(req)
}
func ProcessAbnormalProcess(req *risk.ProcessReq) error {
	return risk.AbnormalProcessProcess(req)
}

func SystemCmdList(req *risk.RiskSearchReq) (risk.RiskSearchResp, error) {
	return risk.SystemCmdList(req)
}
func SystemCmdDetail(macCode, id string) (map[string]interface{}, error) {
	return risk.SystemCmdDetail(macCode, id)
}
func SystemCmdDetailList(req *risk.DetailReq) (info risk.DetailResp, err error) {
	return risk.SystemCmdDetailList(req)
}
func ProcessSystemCmd(req *risk.ProcessReq) error {
	return risk.SystemCmdProcess(req)
}
