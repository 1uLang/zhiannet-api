package users

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	Edgeusers struct {
		Id       uint64 `gorm:"column:id" json:"id" form:"id"`                   //id
		Username string `gorm:"column:username" json:"username" form:"username"` //username
		Name     string `gorm:"column:fullname" json:"fullname" form:"fullname"`             //name
		Email    string `gorm:"column:email" json:"email" form:"email"`          //email
	}
)

//获取用户id
func GetUserInfoByUsername(req *Edgeusers) (ent Edgeusers, err error) {
	if req == nil || req.Username == "" {
		err = fmt.Errorf("参数错误")
		return
	}
	err = db_model.MysqlConn.Model(&Edgeusers{}).Where("username=?", req.Username).Find(&ent).Error

	return
}
