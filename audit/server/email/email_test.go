package email

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

//数据库
func Test_email_info(t *testing.T) {
	list, err := GetEmail(&EmailInfoReq{
		User: &request.UserReq{
			//AdminUserId: 1,
			//UserId: 1,
		},
	})
	fmt.Println(list)
	fmt.Println(err)
}

func Test_email_set(t *testing.T) {
	list, err := SetEmail(&SetEmailReq{
		User:     &request.UserReq{},
		Host:     "smtp.exmail.qq.com",
		Password: "6ratKdr9f9ggHFf4",
		Port:     465,
		Username: "dengb@zhiannet.com",
	})
	fmt.Println(list)
	fmt.Println(err)
}

func Test_email_check(t *testing.T) {
	list, err := CheckEmail(&SetEmailReq{
		User:     &request.UserReq{},
		Host:     "smtp.exmail.qq.com",
		Password: "6ratKdr9f9ggHFf4",
		Port:     465,
		Username: "dengb@zhiannet.com",
		To:       "404821634@qq.com",
	})
	fmt.Println(list)
	fmt.Println(err)
}

// smtp.exmail.qq.com 465 dengb@zhiannet.com 6ratKdr9f9ggHFf4
