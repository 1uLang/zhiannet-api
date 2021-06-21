package monitor_notice

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	MonitorNotice struct {
		Id            uint64 `gorm:"column:id" json:"id" form:"id"`
		MonitorListId uint64 `gorm:"column:monitor_list_id" json:"monitor_list_id" form:"monitor_list_id"`
		Title         string `gorm:"column:title" json:"title" form:"title"`
		Message       string `gorm:"column:message" json:"message" form:"message"`
		UserId        uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`
		IsDelete      int    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
		CreateTime    int    `gorm:"column:create_time" json:"create_time" form:"create_time"`
	}
	ListReq struct {
		MonitorListId uint64 `json:"monitor_list_id"`
		Message       string `json:"message"`
		PageNum       int    `json:"page_num"`
		PageSize      int    `json:"page_size"`
	}
)

func GetList(req *ListReq) (list []*MonitorNotice, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&MonitorNotice{}).Where("is_delete=?", 0)
	if req != nil {
		if req.Message != "" {
			model = model.Where("message like ?", "%"+req.Message+"%")
		}
		if req.MonitorListId > 0 {
			model = model.Where("monitor_list_id = ?", req.MonitorListId)
		}

	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	return
}

//删除
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&MonitorNotice{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

//添加
func Add(req *MonitorNotice) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := model.MysqlConn.Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}
