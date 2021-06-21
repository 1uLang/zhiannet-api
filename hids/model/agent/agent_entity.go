package agent

type SearchReq struct {
	UserName      string `json:"userName,omitempty"`
	PageNo        int    `json:"pageNo"`
	PageSize      int    `json:"pageSize"`
	ServerIp      string `json:"serverIp"`
	ServerLocalIp string `json:"ServerLocalIp"`
}
type SearchResp struct {
	PageNo    int                      `json:"pageNo"`
	PageSize  int                      `json:"pageSize"`
	TotalData int                      `json:"totalData"`
	TotalPage int                      `json:"totalPage"`
	List      []map[string]interface{} `json:"remoteServerAgentStateInfoList"`
}
