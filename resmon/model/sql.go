package model

// ResMonModel 资源监控表
type ResMonModel struct {
	ID     string `gorm:"column:id"`
	OSType uint8  `gorm:"column:os_type"`
}

func (ResMonModel)TableName() string {
	return "resmon"
}

// CPUTypeMap CPU架构映射
var CPUTypeMap = map[uint8]string{
	1: "monit-agent-linux-amd64-v0.1.9.zip",    // Linux 64位，通用cpu架构
	2: "monit-agent-windows-amd64-v0.1.9.zip",  // Windows 64位
	3: "monit-agent-darwin-amd64-v0.1.9.zip",   // MacOS 64位
	4: "monit-agent-freebsd-amd64-v0.1.9.zip",  // Freebsd 64位
	5: "monit-agent-linux-arm64-v0.1.9.zip",    // Linux ARM 64位，ARM架构CPU专用
	6: "monit-agent-linux-mips64-v0.1.9.zip",   // Linux Mips 64，MIPS架构CPU专用
	7: "monit-agent-linux-mips64le-v0.1.9.zip", // Linux Mips 64 LE，MIPS64LE架构CPU专用
}

func GetCPUType(id uint8) string {
	if value, ok := CPUTypeMap[id]; ok {
		return value
	}

	return CPUTypeMap[1]
}
