package server

import (
	"fmt"
	"testing"
)

//func init() {
//	model.InitMysqlLink()
//}

//获取主机防护系统节点信息
func Test_get_hids(t *testing.T) {
	info, err := GetHideInfo()
	fmt.Println(info)
	fmt.Println(err)
}
