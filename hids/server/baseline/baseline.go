package baseline

import "github.com/1uLang/zhiannet-api/hids/model/baseline"

func List(req *baseline.SearchReq) (baseline.SearchResp, error) {
	return baseline.List(req)
}
func Check(req *baseline.CheckReq) error {
	return baseline.Check(req)
}
func TemplateList(req *baseline.TemplateSearchReq) (baseline.TemplateSearchResp, error) {
	return baseline.TemplateList(req)
}
func TemplateDetail(req *baseline.TemplateDetailReq) (baseline.TemplateSearchResp, error) {
	return baseline.TemplateDetail(req)
}
func Detail(req *baseline.DetailReq) (baseline.DetailResp, error) {
	return baseline.Detail(req)
}
