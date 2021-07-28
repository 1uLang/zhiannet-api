package users

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	Edgeusers struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		Username    string `gorm:"column:username" json:"username" form:"username"`                //username
	}
)

//获取用户id
func GetUserIdByUsername(req *Edgeusers) (id uint64, err error) {
	if req == nil ||req.Username == ""{
		err = fmt.Errorf("参数错误")
		return
	}
	var entity Edgeusers
	err = db_model.MysqlConn.Model(&Edgeusers{}).Where("username=?", req.Username).Find(&entity).Error
	id = entity.Id
	return
}
