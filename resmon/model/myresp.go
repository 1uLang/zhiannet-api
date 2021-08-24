package model

// AgentInfo 代理主机信息
type AgentInfo struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Host   string `json:"host"`
	OS     string `json:"os"`
	Cpu    string `json:"cpu"`
	Mem    string `json:"mem"`
	Disk   string `json:"disk"`
	Status bool   `json:"status"`
	On     bool   `json:"on"`
	Key    string `json:"key"`
	OsType uint8  `json:"os_type"`
}

type AgentList struct {
	Total int         `json:"total"`
	List  []AgentInfo `json:"list"`
}

type DownInfo struct {
	Host    string `json:"host"`
	DownUrl string `json:"down_url"`
}
