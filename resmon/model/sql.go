package model

// ResMonModel 资源监控表
type ResMonModel struct {
	ID     string `gorm:"column:id"`
	OSType uint8  `gorm:"column:os_type"`
}
