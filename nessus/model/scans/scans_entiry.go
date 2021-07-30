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

type ListReq struct {
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
	Targets     bool   `json:"-"`
	Scan        bool   `json:"-"`
	Report      bool   `json:"-"`
}
type ScanReq struct {
	ID string
}
type PauseReq struct {
	ID string
}

type ResumeReq struct {
	ID string
}

type ExportReq struct {
	ID     string
	Format string
}
type ExportResp struct {
	Token string
	File  float64
}
type VulnerabilitiesReq struct {
	ID string
}

type PluginsReq struct {
	ScanId string
	VulId  string
}

type DeleteReq struct {
	ID string
}
type ResetReq struct {
	ID string
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}
