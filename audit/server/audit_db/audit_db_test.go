package audit_db

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
	"github.com/1uLang/zhiannet-api/audit/server"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//数据库
func Test_dblist(t *testing.T) {
	list, err := GetAuditBdList(&ReqSearch{
		User: &request.UserReq{
			//AdminUserId: 1,
			UserId: 2,
		},
		PageSize: 1,
		Name:     "222",
		//Status: "1",
	})
	fmt.Println(list)
	fmt.Println(err)
}

////添加
func Test_add_db(t *testing.T) {
	list, err := AddDb(&DBReq{
		User: &request.UserReq{UserId: 2},
		//Uid:      1,
		Type:     1,
		Name:     "test",
		Version:  "8",
		IP:       "1.2.2.2",
		Port:     "3306",
		System:   1,
		Status:   1,
		TimeLong: 0,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//
////修改
func Test_edit(t *testing.T) {
	list, err := EditDb(&DBEditReq{
		User: &request.UserReq{AdminUserId: 1},
		Name: "aaaa",
		Id:   29,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//删除 DelDb
func Test_del(t *testing.T) {
	list, err := DelDb(&DelDbReq{
		User: &request.UserReq{AdminUserId: 1},
		Id:   29,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//授权
func Test_audit_db(t *testing.T) {
	list, err := AuthDb(&server.AuthReq{
		User:  &request.UserReq{AdminUserId: 1},
		Id:    26,
		Email: []string{"449588335@qq.com"},
	})
	fmt.Println(list)
	fmt.Println(err)
}

//获取授权列表
func Test_audit_list(t *testing.T) {
	list, err := GetAuthEmail(&server.AuthReq{
		User: &request.UserReq{UserId: 2},
		Id:   38,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//日志列表
func Test_log_list(t *testing.T) {
	list, err := GetDbLog(&DbLogReq{
		UserId:   &request.UserReq{AdminUserId: 1},
		AuditId:  []string{"db-dfc7e615-a483-4a41-8bdb-244fbb9687c3"},
		TimeType: "30day",
		Page:     1, Size: 10,
		Message: "1626523801266",
		Export:  true,
	})
	fmt.Println(list)
	fmt.Println(err)
}
