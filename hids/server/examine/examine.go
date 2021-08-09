package examine

import "github.com/1uLang/zhiannet-api/hids/model/examine"

func List(req *examine.SearchReq) (examine.SearchResp, error) {
	return examine.List(req)
}
func Details(args *examine.DetailsReq) (info examine.DetailsResp, err error) {
	return examine.Details(args)
}
func ScanServerNow(req *examine.ScanReq) error {
	return examine.ScanServerNow(req)
}
func ScanServerCancel(macCodes []string) error {
	return examine.ScanServerCancel(macCodes)
}
