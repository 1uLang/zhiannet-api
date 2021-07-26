package audit_from

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//数据库
func Test_fromlist(t *testing.T) {
	list, err := GetAuditFromList(&ReqSearch{
		User:     &request.UserReq{AdminUserId: 1},
		PageSize: 1,
		//Name: "罗兵",
	})
	fmt.Println(list)
	fmt.Println(err)
}

////添加
func Test_add_db(t *testing.T) {
	list, err := AddFrom(&FromReq{
		User:       &request.UserReq{AdminUserId: 1},
		UserId:     1,
		Name:       "test",
		Cycle:      1,
		CycleDay:   1,
		SendTime:   "12:01",
		Format:     1,
		AssetsType: 1,
		AssetsId:   1,
		Email:      "12333@qq.com",
	})
	fmt.Println(list)
	fmt.Println(err)
}

//
////修改
func Test_edit(t *testing.T) {
	list, err := EditFrom(&FromReq{
		User:       &request.UserReq{AdminUserId: 1},
		UserId:     1,
		Name:       "test",
		Cycle:      1,
		CycleDay:   1,
		SendTime:   "12:01",
		Format:     1,
		AssetsType: 1,
		AssetsId:   1,
		Email:      "11111@qq.com",
		Id:         14,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//删除 DelDb
func Test_del(t *testing.T) {
	list, err := DelFrom(&DelFromReq{
		User: &request.UserReq{AdminUserId: 1},
		Id:   14,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//获取授权列表
func Test_info(t *testing.T) {
	list, err := GetFrom(&GetFromReq{
		User: &request.UserReq{AdminUserId: 1},
		Id:   14,
	})
	fmt.Println(list)
	fmt.Println(err)
}
