package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}
func Test_login(t *testing.T) {
	res, err := GetLoginInfo(&UserReq{})
	fmt.Println(res)
	fmt.Println(err)
}
