package vulnerabilities

type ListReq struct {
	Limit       int    `json:"l,omitempty"` //限制条数
	C           int    `json:"c,omitempty"` //偏移量
	Query       string `json:"q,omitempty"` //筛选器
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}
