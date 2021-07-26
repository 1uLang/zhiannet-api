package commands

type ListReq struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`

	System_user string `json:"system_user,omitempty"`
	Date_from   string `json:"date_from,omitempty"`
	Date_to     string `json:"date_to,omitempty"`
	Risk_level  string `json:"risk_level,omitempty"`
	Asset       string `json:"asset,omitempty"`
	Search      string `json:"search,omitempty"`

	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}
