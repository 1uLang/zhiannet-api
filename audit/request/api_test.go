package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/model/audit_user_relation"
	"testing"
	"time"
)

func Test_ApiLogin(t *testing.T) {

	cook, err := GetToken(&audit_user_relation.AuditReq{
		AdminUserId: 1,
	})
	fmt.Println("res=", cook, err)
	time.Sleep(time.Second * 5)
}

//
////全局统计  GetStatusGlobal
//func Test_status_global(t *testing.T) {
//	global_status.GetStatusGlobal()
//}
//
////负载信息
//func Test_load(t *testing.T) {
//	global_status.GetLoad()
//}

////主机列表
//func Test_host(t *testing.T){
//	HostStatusApi()
//}
