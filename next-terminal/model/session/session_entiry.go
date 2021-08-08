package session

type (
	ListReq struct {
		PageIndex   int    `json:"pageIndex"`
		PageSize    int    `json:"pageSize"`
		Status 		string `json:"status"`
		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
	DeleteReq struct {
		Id 	string
	}
	ReplayReq struct {
		Id 	string
	}
	ReplayResp struct {

	}
	MonitorReq struct {

	}
	MonitorResp struct {

	}
	DisConnectReq struct {
		Id string
	}
)