package agents

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

type HIDSAgent struct {
	Id       uint64 `gorm:"column:id" json:"id" form:"id"`                      //id
	AgentId  string `gorm:"column:agent_id" json:"agent_id" form:"agent_id"`    //agent_id
	IsDelete uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"` //1删除
	Remake   string `gorm:"column:remake" json:"remake" form:"remake"`          //备注
}

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&HIDSAgent{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

func addAgent(req *HIDSAgent) (err error) {

	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	return db_model.MysqlConn.Create(&req).Error
}
func updateAgent(req *HIDSAgent) (err error) {
	if req == nil || req.AgentId == "" {
		return fmt.Errorf("参数错误")
	}
	ent := HIDSAgent{}
	md := db_model.MysqlConn.Model(req).Where("agent_id=?", req.AgentId).Where("is_delete=?", 0)

	var count int64
	err = md.Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return addAgent(req)
	}
	err = md.Find(&ent).Error
	if err != nil {
		return err
	}

	ent.Remake = req.Remake
	return db_model.MysqlConn.Model(ent).Where("id=?", ent.Id).Save(&ent).Error
}
func deleteAgent(agentIds []string) error {
	for _, agentId := range agentIds {
		err := db_model.MysqlConn.Model(&HIDSAgent{}).Where("agent_id = ?", agentId).Update("is_delete", 1).Error
		if err != nil {
			return err
		}
	}
	return nil
}
func getInfo(agentId string) (remake string, err error) {
	ent := &HIDSAgent{}
	err = db_model.MysqlConn.Model(ent).Where("agent_id=?", agentId).Where("is_delete=0").Find(&ent).Error
	return ent.Remake, err
}
