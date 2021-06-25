package reports

type CreateResp struct {
	Source struct {
		IDS  []string `json:"id_list"`
		Type string   `json:"list_type"`
	} `json:"source"`
	TemplateId string `json:"template_id"`
	//TargetIds   []string `json:"target_ids"`
	AdminUserId uint64 `json:"admin_user_id"`
	UserId      uint64 `json:"user_id"`
}

type ListReq struct {
	Limit       int    `json:"l,omitempty"` //限制条数
	C           int    `json:"c,omitempty"` //偏移量
	UserId      uint64 `json:"user_id"`
	AdminUserId uint64 `json:"admin_user_id"`
}
