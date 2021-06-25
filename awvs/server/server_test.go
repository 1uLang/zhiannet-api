package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"testing"
)

func init() {
	model.InitMysqlLink()
}
func Test_server(t *testing.T) {
	res, err := GetWebScan()
	fmt.Println(res)
	fmt.Println(err)
}
