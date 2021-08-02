package user

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/audit/request"
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
	list, err := GetRole(&request.UserReq{UserId: 2})
	fmt.Println(list)
	fmt.Println(err)
}

////添加
func Test_add(t *testing.T) {
	list, err := AddUser(&AddUserReq{
		User:        &request.UserReq{AdminUserId: 1},
		Email:       "4048216340@qq.com",
		IsAdmin:     1,
		NickName:    "test4",
		Opt:         1,
		Password:    "18113470664",
		Phonenumber: "18113470664",
		RoleIds:     []uint64{},
		RoleName:    "平台管理员",
		Sex:         1,
		Status:      1,
		UserName:    "test4",
	})
	fmt.Println(list)
	fmt.Println(err)
}
