package agents

type StatisticsResp struct {
	Active         int `json:"active"`
	Disconnected   int `json:"disconnected"`
	NeverConnected int `json:"never_connected"`
	Pending        int `json:"pending"`
	Total          int `json:"total"`
}
type ListReq struct {
	IP string `json:"ip,omitempty"`
}
type ListResp struct {
	AffectedItems []struct {
		ConfigSum     string   `json:"configSum"`
		DateAdd       string   `json:"dateAdd"`
		Group         []string `json:"group"`
		ID            string   `json:"id"`
		IP            string   `json:"ip"`
		LastKeepAlive string   `json:"lastKeepAlive"`
		Manager       string   `json:"manager"`
		MergedSum     string   `json:"mergedSum"`
		Name          string   `json:"name"`
		NodeName      string   `json:"node_name"`
		Os            struct {
			Arch     string `json:"arch"`
			Codename string `json:"codename"`
			Major    string `json:"major"`
			Minor    string `json:"minor"`
			Name     string `json:"name"`
			Platform string `json:"platform"`
			Uname    string `json:"uname"`
			Version  string `json:"version"`
		} `json:"os"`
		RegisterIP string `json:"registerIP"`
		Status     string `json:"status"`
		Version    string `json:"version"`
	} `json:"affected_items"`
	TotalAffectedItems int64 `json:"total_affected_items"`
}

type SCAListResp struct {
	AffectedItems []struct {
		Description string `json:"description"`
		EndScan     string `json:"end_scan"`
		Fail        int64  `json:"fail"`
		HashFile    string `json:"hash_file"`
		Invalid     int64  `json:"invalid"`
		Name        string `json:"name"`
		Pass        int64  `json:"pass"`
		PolicyID    string `json:"policy_id"`
		References  string `json:"references"`
		Score       int64  `json:"score"`
		StartScan   string `json:"start_scan"`
		TotalChecks int64  `json:"total_checks"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}
