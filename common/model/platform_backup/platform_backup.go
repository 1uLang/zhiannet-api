package platform_backup

import (
	"fmt"
	model "github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	PlatformBackup struct {
		Id         int64  `gorm:"column:id" json:"id" form:"id"`                            //id
		Name       string `gorm:"column:name" json:"name" form:"name"`                      //文件名
		Size       string `gorm:"column:size" json:"size" form:"size"`                      //文件大小
		CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"` //创建时间
	}

	ListReq struct {
		PageNum  int       `json:page_num`
		PageSize int       `json:page_size`
		Name     string    `json:"name"`
		LastTime time.Time `json:"time"`
		Ids      []int64   `json:"ids"`
	}
)

//初始化建表
func InitTable() {
	err := model.MysqlConn.AutoMigrate(&PlatformBackup{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

//获取列表
func GetList(req *ListReq) (list []*PlatformBackup, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&PlatformBackup{})
	if req.Name != "" {
		model = model.Where("name like ?", "%"+req.Name+"%")
	}
	if !req.LastTime.IsZero() {
		model = model.Where("create_time <?", req.LastTime.Format("2006-01-02 15:04:05"))
	}
	if len(req.Ids) > 0 {
		model = model.Where("id in(?)", req.Ids)
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
	if err != nil {
		return
	}
	return
}

//新增更新
func Save(req *PlatformBackup) (rows int64, err error) {
	var entity PlatformBackup

	err = model.MysqlConn.Where("name=?", req.Name).Find(&entity).Error
	if err != nil {
		return
	}
	entity.Name = req.Name
	entity.Size = req.Size
	entity.CreateTime = time.Now().Unix()
	res := model.MysqlConn.Model(&PlatformBackup{}).Where("id=?", entity.Id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected

	return
}

func Delete(id []int64) error {
	res := model.MysqlConn.Where("id in (?)", id).Delete(&PlatformBackup{})
	return res.Error
}

func GetInfo(id int64) (info *PlatformBackup, err error) {
	err = model.MysqlConn.Where("id =?", id).First(&info).Error
	return
}
