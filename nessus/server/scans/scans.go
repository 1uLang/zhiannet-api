package scans

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/nessus/model/scans"
)

// 创建扫描
func Create(req *scans.AddReq) error {

	id, err := scans.Create(req)
	if err != nil {
		return err
	}
	fmt.Println("create nessus scans success : ", id)
	return nil
}

// 扫描列表
func List(req *scans.ListReq) ([]interface{}, error) {
	return scans.List(req)
}

// 扫描
func Scans(req *scans.ScanReq) error {
	return scans.Scans(req)
}

// 暂停
func Pause(req *scans.PauseReq) error {
	return scans.Pause(req)
}

// 重新扫描
func Resume(req *scans.ResumeReq) error {
	return scans.Resume(req)
}

// 导出
func Export(req *scans.ExportReq) (*scans.ExportResp, error) {
	return scans.Export(req)
}

// 导出
func ExportFile(req *scans.ExportFileReq) ([]byte, string, error) {
	return scans.ExportFile(req)
}

// 漏洞列表
func Vulnerabilities(req *scans.VulnerabilitiesReq) ([]interface{}, error) {
	return scans.Vulnerabilities(req)
}

// 漏洞详情
func Plugins(req *scans.PluginsReq) (map[string]interface{}, error) {
	return scans.Plugins(req)
}

// 删除扫描
func Delete(req *scans.DeleteReq) error {
	return scans.Delete(req)
}

// 重置扫描
func Reset(req *scans.ResetReq) error {
	return scans.Reset(req)
}

// 扫描记录
func History(req *scans.HistoryReq) ([]interface{}, error) {
	return scans.History(req)
}

// 扫描记录
func DelHistory(req *scans.DelHistoryReq) error {
	return scans.DelHistory(req)
}

// 扫描记录
func CreateReport(req *scans.CreateReportReq) error {
	return scans.CreateReport(req)
}

// 扫描记录
func ListReport(req *scans.ScansListReq) ([]interface{}, error) {
	return scans.ListReport(req)
}

// 扫描记录
func DeleteReport(req []string) error {
	return scans.DeleteReport(req)
}
