package model

// CPUUsage cpu使用率
type CPUUsage struct {
	Usage struct {
		All []float64 `json:"all"`
		Avg float64   `json:"avg"`
	} `json:"usage"`
}

// MemUsage 内存使用量
type MemUsage struct {
	Usage struct {
		// SwapFree int `json:"swapFree"`
		// SwapPercent int `json:"swapPercent"`
		// SwapTotal int `json:"swapTotal"`
		// SwapUsed int `json:"swapUsed"`
		// VirtualBuffers int `json:"virtualBuffers"`
		// VirtualCached int `json:"virtualCached"`
		// VirtualFree float64 `json:"virtualFree"`
		// VirtualPercent float64 `json:"virtualPercent"`
		VirtualTotal float64 `json:"virtualTotal"`
		VirtualUsed  float64 `json:"virtualUsed"`
		// VirtualWired float64 `json:"virtualWired"`
	} `json:"usage"`
}

// DiskUsage 磁盘使用量
type DiskUsage struct {
	Partitions []struct {
		// Free          int64   `json:"free"`
		// Fstype        string  `json:"fstype"`
		// InodesFree    int64   `json:"inodesFree"`
		// InodesPercent float64 `json:"inodesPercent"`
		// InodesTotal   int64   `json:"inodesTotal"`
		// InodesUsed    int     `json:"inodesUsed"`
		Name string `json:"name"`
		// Percent       float64 `json:"percent"`
		Total int64 `json:"total"`
		Used  int64 `json:"used"`
	} `json:"partitions"`
}

// Agents agent 列表
type Agents []struct {
	Config struct {
		ID   string `json:"id"`
		On   bool   `json:"on"`
		Name string `json:"name"`
		Host string `json:"host"`
		Key  string `json:"key"`
	} `json:"config"`
}

// AgentState 代理主机状态
type AgentState struct {
	Version  string  `json:"version"`   // 版本号
	OSName   string  `json:"os_name"`   // 操作系统
	Speed    float64 `json:"speed"`     // 连接速度，ms
	IP       string  `json:"ip"`        // IP地址
	IsActive bool    `json:"is_active"` // 是否在线
}

type BaseResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AddAgentResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AgentID string `json:"agentId"`
	} `json:"data"`
}
