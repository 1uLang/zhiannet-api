package scans

type AddReq struct {
	UUID     string `json:"uuid"`
	Settings struct {
		Name         string `json:"name"`
		Text_targets string `json:"text_targets"`
		Description  string `json:"description"`
	} `json:"settings"`

	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}
