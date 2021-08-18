package audit_assets_relation_server

import "github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"

func GetList(req *audit_assets_relation.ListReq) (list []*audit_assets_relation.AuditAssetsRelation, total int64, err error) {
	return audit_assets_relation.GetList(req)
}

func Reset(req *audit_assets_relation.AddReq) (err error) {
	return audit_assets_relation.Reset(req)
}

func Add(req *audit_assets_relation.AuditAssetsRelation) (id uint64, err error) {
	return audit_assets_relation.Add(req)
}
