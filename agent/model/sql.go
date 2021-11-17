package model

import (
	"fmt"
	db_model "github.com/1uLang/zhiannet-api/common/model"
)

// AgentFile Agent文件详情表
type AgentFile struct {
	ID        int64  `gorm:"column:id;primaryKey" json:"id"`      // 主键ID
	Name      string `gorm:"column:name" json:"name"`             // 文件名
	Describe  string `gorm:"column:describe" json:"describe"`     // 文件描述信息
	Size      int    `gorm:"column:size" json:"size"`             // 文件大小
	Format    string `gorm:"column:format" json:"format"`         // 文件格式
	State     uint8  `gorm:"column:state"`                        // 文件状态 1正常 0删除
	Path      string `gorm:"column:path"`                         // 文件存储路径
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 文件上传时间
	UpdatedAt int64  `gorm:"column:updated_at"`                   // 文件更新时间
}

// TableName 表明映射
func (AgentFile) TableName() string {
	return "agent_file"
}

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&AgentFile{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}
