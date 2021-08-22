package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"testing"
	"time"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//测试获取cook
func Test_get_cook(t *testing.T) {
	cook, _, err := GetCookie(&ApiKey{
		Username: "root",
		Password: "21ops.com",
		Port:     "5443",
		Addr:     "182.150.0.109",
	})
	fmt.Println(cook)
	fmt.Println(err)
}

//测试获取一个接口
func Test_global(t *testing.T) {
	GetGlobal(&ApiKey{
		Username: "root",
		Password: "21ops.com",
		Port:     "5443",
		Addr:     "182.150.0.109",
	})
}

//测试获取nat 1:1 接口
func Test_nat_1to1(t *testing.T) {
	//Nat1to1(&ApiKey{
	//	Username: "root",
	//	Password: "21ops.com",
	//	Port:     "5443",
	//	Addr:     "182.150.0.109",
	//})
}

//测试获取日志接口
func Test_logs(t *testing.T) {
	//GetLogsList(&ApiKey{
	//	Username: "root",
	//	Password: "21ops.com",
	//	Port:     "5443",
	//	Addr:     "182.150.0.109",
	//})
}

////获取唯一key
//func Test_get_keys(t *testing.T){
//	key,value := GetUniqueKey(&ApiKey{
//		Username: "root",
//		Password: "21ops.com",
//		Port: "5443",
//		Addr: "182.150.0.109",
//	})
//	fmt.Println(key)
//	fmt.Println(value)
//}
//func Test_login(t *testing.T){
//	cook,err := CollyLongin(&ApiKey{
//		Username: "root",
//		Password: "21ops.com",
//		Port: "5443",
//		Addr: "182.150.0.109",
//	})
//	fmt.Println(cook)
//	fmt.Println(err)
//	/// [PHPSESSID=fd75860aa1e86e96d1d27220a7edbc29 cookie_test=a299bc1183cd5cef45ec415cfd569f9d]
//}

func Test_insert_mesg(t *testing.T) {
	res, err := edge_messages.Add(&edge_messages.Edgemessages{
		Level:     "error",
		Subject:   "组件状态异常",
		Body:      "云防火墙状态不可用",
		Type:      "AdminAssembly",
		Params:    "{}",
		Createdat: uint64(time.Now().Unix()),
		Day:       time.Now().Format("20060102"),
		Hash:      "",
		Role:      "admin",
	})
	fmt.Println(res, err)
}
