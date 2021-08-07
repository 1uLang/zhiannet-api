package server

import (
	"fmt"
	dbmodel "github.com/1uLang/zhiannet-api/common/model"
	model "github.com/1uLang/zhiannet-api/edgeUsers/model"
	"testing"
)

func init() {
	dbmodel.InitMysqlLink()
}

func Test_list(t *testing.T) {
	list, err := ListEnabledUsers(
		&model.ListReq{
			UserId: 1,
		})
	fmt.Println(list[0].OtpParams, err)
}
