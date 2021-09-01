package bwlist

import "github.com/1uLang/zhiannet-api/hids/model/bwlist"

func AddBWList(req *bwlist.HIDSBWList) (err error) {
	return bwlist.AddBWList(req)
}
func GetBWList(req *bwlist.ListReq) (list []*bwlist.HIDSBWList, total int64, err error) {
	return bwlist.GetBWList(req)
}
func DeleteBWList(id uint64) error{
	return bwlist.DeleteBWList(id)
}
