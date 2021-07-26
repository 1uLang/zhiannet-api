package audit_host

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

//
func Test_list(t *testing.T) {
	list, err := GetAuditHostList(&ReqSearch{
		User:     &request.UserReq{AdminUserId: 1},
		PageSize: 1,
		//Name: "罗兵",
		Status: "1",
	})
	fmt.Println(list)
	fmt.Println(err)
}

////添加
func Test_add_host(t *testing.T) {
	list, err := AddHost(&HostReq{
		User:     &request.UserReq{AdminUserId: 1},
		Uid:      1,
		Name:     "test",
		IP:       "1.2.2.2",
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
	list, err := EditHost(&HostEditReq{
		User: &request.UserReq{AdminUserId: 1},
		Name: "aaaa",
		Id:   29,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//删除 DelDb
func Test_del(t *testing.T) {
	list, err := DelHost(&DelHostReq{
		User: &request.UserReq{AdminUserId: 1},
		Id:   29,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//授权
func Test_audit_host(t *testing.T) {
	list, err := AuthHost(&server.AuthReq{
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
		User: &request.UserReq{AdminUserId: 1},
		Id:   26,
	})
	fmt.Println(list)
	fmt.Println(err)
}

//日志列表
func Test_log_list(t *testing.T) {
	list, err := GetHostLog(&HostLogReq{
		UserId:   &request.UserReq{AdminUserId: 1},
		AuditId:  []string{"host-dengpang"},
		TimeType: "30day",
		Page:     1, Size: 10,
		//Message: "1626523801266",
		//Export: true,
	})
	fmt.Println(list)
	fmt.Println(err)
}
