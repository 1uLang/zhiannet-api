package agent

type SearchReq struct {
	UserName      string `json:"userName,omitempty"`
	PageNo        int    `json:"pageNo"`
	PageSize      int    `json:"pageSize"`
	ServerIp      string `json:"serverIp,omitempty"`
	ServerLocalIp string `json:"ServerLocalIp,omitempty"`

	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}
type SearchResp struct {
	PageNo    int                      `json:"pageNo"`
	PageSize  int                      `json:"pageSize"`
	TotalData int                      `json:"totalData"`
	TotalPage int                      `json:"totalPage"`
	List      []map[string]interface{} `json:"remoteServerAgentStateInfoList"`
	//agentState 	1:启用中，2：已启用，3：停用中，4：已停用，5：卸载中，6：已卸载
}
type CreateReq struct {
	AgentIp string

	UserId      uint64
	AdminUserId uint64
}

type UpdateReq struct {
	Id      uint64
	AgentIp string

	UserId      uint64
	AdminUserId uint64
}
type DeleteReq struct {
	Id uint64
}
