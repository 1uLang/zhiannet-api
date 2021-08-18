package reports

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	WebscanReport struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		ReportId    string `gorm:"column:report_id" json:"report_id" form:"report_id"`             //报告ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8    `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除 0未删除
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	ReportListReq struct {
		UserId      uint64 `json:"user_id" gorm:"column:user_id"`                      // 用户ID
		AdminUserId uint64 `json:"admin_user_id" gorm:"column:admin_user_id"`          // admin用户ID
		ReportId    string `gorm:"column:report_id" json:"report_id" form:"report_id"` //报告ID
		PageNum     int    `json:"page_num" `                                          //页数
		PageSize    int    `json:"page_size" `                                         //每页条数
	}
)

//初始化建表
func InitTable() {
	err := model.MysqlConn.AutoMigrate(&WebscanReport{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}
//获取列表信息
func GetList(req *ReportListReq) (list []*WebscanReport, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&WebscanReport{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.ReportId != "" {
			model = model.Where("report_id=?", req.ReportId)
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
	err = model.Debug().Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//添加数据
func AddAddr(req *WebscanReport) (insertId uint64, err error) {
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

//用reportID 删除
func DeleteByTargetIds(ids []string) (err error) {
	res := model.MysqlConn.Model(&WebscanReport{}).Where("report_id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&WebscanReport{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}
