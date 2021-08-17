package audit_assets_relation_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/audit_assets_relation"
	"testing"
)

func init() {
	model.InitMysqlLink()
}

func Test_list(t *testing.T) {
	list, total, e := GetList(&audit_assets_relation.ListReq{
		UserId: 11,
	})
	fmt.Println(list[0], total, e)
}

func Test_add(t *testing.T) {
	e := Reset(&audit_assets_relation.AddReq{
		UserId:      1,
		AdminUserId: 0,
		AuditId:     []string{"112312312312"},
		AssetsType:  0,
	})
	fmt.Println(e)
}
