package reports

import (
	"github.com/1uLang/zhiannet-api/awvs/model/reports"
)

//scans 扫描接口层

//List 扫描列表
func List(req *reports.ListReq) (info map[string]interface{}, err error) {
	return reports.List(req)
}

//Create 新建扫描
func Create(req *reports.CreateResp) (info map[string]interface{}, err error) {
	return reports.Create(req)
}

//Delete 删除扫描
func Delete(scan_id string) (err error) {
	return reports.Delete(scan_id)
}

//下载报表 去掉awvs字样
func Download(url string, pdf bool) ([]byte, string, error) {
	return reports.Download(url, pdf)
}
