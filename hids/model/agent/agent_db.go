package agent

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type (
	hidsAgents struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		IP          string `gorm:"column:ip" json:"ip" form:"ip"`                                  //agent ip
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int64  `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	ListReq struct {
		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
)

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&hidsAgents{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

func checkAgentIP(req *hidsAgents) (bool, error) {

	model := db_model.MysqlConn.Model(&hidsAgents{}).Where("is_delete=?", 0).Where("ip=?", req.IP)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
	}

	var total int64
	err := model.Count(&total).Error
	return total > 0, err
}

func addAgent(req *hidsAgents) (err error) {

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	ok, err := checkAgentIP(req)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("该IP已添加")
	}
	return db_model.MysqlConn.Create(&req).Error
}
func updateAgent(req *hidsAgents) (err error) {
	if req == nil || req.Id == 0 {
		return fmt.Errorf("参数错误")
	}
	ent := hidsAgents{}
	md := db_model.MysqlConn.Model(req).Where("id=?", req.Id)
	err = md.Find(&ent).Error
	if err != nil {
		return err
	}

	ent.IP = req.IP

	return db_model.MysqlConn.Model(ent).Where("id=?", req.Id).Save(&ent).Error
}
func deleteAgent(id uint64) error {
	res := db_model.MysqlConn.Model(&hidsAgents{}).Where("id = ?", id).Update("is_delete", 1)
	return res.Error
}
func GetList(req *ListReq) (list []*hidsAgents, total int64, err error) {
	//从数据库获取
	model := db_model.MysqlConn.Model(&hidsAgents{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
	}
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	err = model.Debug().Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}
func GetUserListByAgentIP(ip string) (list []*hidsAgents, err error) {
	model := db_model.MysqlConn.Model(&hidsAgents{}).Where("is_delete=?", 0).Where("ip=?",ip)

	err = model.Debug().Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}
