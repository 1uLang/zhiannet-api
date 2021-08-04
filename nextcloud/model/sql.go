package model

// NextCloudToken nextcloud token表
type NextCloudToken struct {
	ID    int64  `gorm:"column:id"`
	User  string `gorm:"column:user;unqiue"` // 主站的用户名
	UID   int64  `gorm:"column:uid"`         // 主站用户id
	Kind  uint8  `gorm:"column:kind"`        // 0用户端 1管理端
	Token string `gorm:"column:token"`       // nextcloud token
}

// Subassemblynode 节点配置表
type Subassemblynode struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Addr     string `gorm:"column:addr"`
	State    uint8  `gorm:"column:state"`     // 1启用、0禁用
	IsDelete uint8  `gorm:"column:is_delete"` // 1删除
	IsSSL    uint8  `gorm:"is_ssl"`           // 1是 0不是
	Key      string `gorm:"column:key"`
	Secret   string `gorm:"column:secret"`
}

// TableName 表名映射
func (NextCloudToken) TableName() string {
	return "nextcloud_token"
}

// TableName 表名映射
func (Subassemblynode) TableName() string {
	return "subassemblynode"
}
