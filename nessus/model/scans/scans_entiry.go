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
type HistoryReq struct {
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}
type DelHistoryReq struct {
	ID        string
	HistoryId string
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
	HistoryId string
	Format string
}
type ExportResp struct {
	Token string
	File  float64
}
type VulnerabilitiesReq struct {
	ID        string
	HistoryId string
}

type PluginsReq struct {
	ID        string
	HistoryId string
	VulId     string
}

type DeleteReq struct {
	ID string
}
type ResetReq struct {
	ID          string
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}

type CreateReportReq struct {
	ID        string `json:"id"`
	HistoryId string `json:"history_id"`
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}

type ListReportReq struct {
	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}