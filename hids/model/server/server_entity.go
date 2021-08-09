package server

type SearchReq struct {
	UserName     string `json:"userName,omitempty"`
	PageNo       int    `json:"pageNo"`
	PageSize     int    `json:"pageSize"`
	ServerIp     string `json:"serverIp,omitempty"`
	HealthLevel  string `json:"healthLevel,omitempty"`   //体检等级 0 （0-59） 1 （60-89） 2 （90-100）
	ServerStatus string `json:"server_status,omitempty"` //主机状态 0 离线 1 在线

	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}

type SearchResp struct {
	PageNo         int                      `json:"pageNo"`
	PageSize       int                      `json:"pageSize"`
	TotalData      int                      `json:"totalData"`
	TotalPage      int                      `json:"totalPage"`
	ServerInfoList []map[string]interface{} `json:"serverInfoList"`
}
type InfoResp struct {
	LocalIp  string
	HostName string
	System   string
	OsType   string
}
