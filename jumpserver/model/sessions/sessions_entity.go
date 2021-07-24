package sessions

type ListReq struct {
	Limit       int    `json:"limit,omitempty"`
	Offset      int    `json:"offset,omitempty"`
	Is_finished string `json:"is_finished"` //在线 - 历史 1 - 0
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}
