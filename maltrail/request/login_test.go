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
	token, err := GetLoginInfo(1)
	fmt.Println(token, err)
}

func Test_change(t *testing.T) {
	Change()
}
