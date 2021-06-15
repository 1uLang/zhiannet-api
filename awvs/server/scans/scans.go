package scans

import "github.com/1uLang/zhiannet-api/awvs/model/scans"

//scans 扫描接口层

//List 扫描列表
func List(req *scans.ListReq) (info map[string]interface{}, err error) {
	return scans.List(req)
}

//Add 新建扫描
func Add(req *scans.AddReq) (err error) {
	return scans.Add(req)
}

//ScanningProfiles 扫描配置文件列表
func ScanningProfiles() (list map[string]interface{}, err error) {
	return scans.ScanningProfiles()
}

//ReportTemplates 报表列表
func ReportTemplates() (list map[string]interface{}, err error) {
	return scans.ReportTemplates()
}

//Abort 停止扫描
func Abort(scan_id string) (err error) {
	return scans.Abort(scan_id)
}

//Delete 删除扫描
func Delete(scan_id string) (err error) {
	return scans.Delete(scan_id)
}
