package global_status

import (
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"testing"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
}

//全局状态
func Test_Global(t *testing.T) {
	InitDB()
	GetGlobalStatus(&request.ApiKey{
		Username: "",
		Password: "",
	})
	//fmt.Println(res)
	//fmt.Println(err)
}

//NAT
func Test_NAT(t *testing.T) {
	InitDB()
	GetNATList(&request.ApiKey{
		Username: "",
		Password: "",
	})
	//fmt.Println(res)
	//fmt.Println(err)
}
