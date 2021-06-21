package monitor_list

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	//表结构
	MonitorList struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`
		Addr        string `gorm:"column:addr" json:"addr" form:"addr"`
		Port        int    `gorm:"column:port" json:"port" form:"port"`
		Code        int    `gorm:"column:code" json:"code" form:"code"`
		Status      int    `gorm:"column:status" json:"status" form:"status"`
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`
		IsDelete    int    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`
		MonitorType int    `gorm:"column:monitor_type" json:"monitor_type" form:"monitor_type"`
	}

	//列表请求参数
	ListReq struct {
		Status      string `json:"status"`
		UserId      uint64 `json:"user_id"`
		MonitorType int    `json:"monitor_type"`
		PageNum     int    `json:"page_num"`
		PageSize    int    `json:"page_size"`

		MinId uint64 `json:"min_id"` //连续翻页时，上页最大的ID
	}

	//修改端口or code状态 请求参数
	SaveReq struct {
		Id     uint64 `json:"id"`
		Status int    `json:"status"`
	}
)

//获取节点
func GetList(req *ListReq) (list []*MonitorList, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&MonitorList{}).Where("is_delete=?", 0)
	if req != nil {
		if req.Status != "" {
			model = model.Where("state=?", req.Status)
		}
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.MonitorType > 0 {
			model = model.Where("monitor_type=?", req.MonitorType)
		}
		if req.MinId > 0 {
			model = model.Where("id>?", req.MinId)
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

//添加
func Add(req *MonitorList) (insertId uint64, err error) {
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

//修改状态
func Save(req *SaveReq) (rows int64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := model.MysqlConn.Model(&MonitorList{}).Where("id = ?", req.Id).Update("status", req.Status)
	if res.Error != nil {
		return 0, res.Error
	}
	rows = res.RowsAffected
	return
}

//删除
func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&MonitorList{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}
