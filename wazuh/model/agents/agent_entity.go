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

type VulnerabilityListResp struct {
	AffectedItems []struct {
		Cve          string `json:"cve"`
		Architecture string `json:"architecture"`
		Version      string `json:"version"`
		Name         string `json:"name"`
	} `json:"affected_items"`
	TotalAffectedItems int           `json:"total_affected_items"`
	TotalFailedItems   int           `json:"total_failed_items"`
	FailedItems        []interface{} `json:"failed_items"`
}
type CiscatListResp struct {
	AffectedItems []struct {
		Benchmark  string `json:"benchmark"`
		Error      int64  `json:"error"`
		Fail       int64  `json:"fail"`
		Notchecked int64  `json:"notchecked"`
		Pass       int64  `json:"pass"`
		Profile    string `json:"profile"`
		Scan       struct {
			ID   int64  `json:"id"`
			Time string `json:"time"`
		} `json:"scan"`
		Score   int64 `json:"score"`
		Unknown int64 `json:"unknown"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

type SysCheckListResp struct {
	AffectedItems []struct {
		Changes int64  `json:"changes"`
		Date    string `json:"date"`
		File    string `json:"file"`
		Gid     string `json:"gid"`
		Gname   string `json:"gname"`
		Inode   int64  `json:"inode"`
		Md5     string `json:"md5"`
		Mtime   string `json:"mtime"`
		Perm    string `json:"perm"`
		Sha1    string `json:"sha1"`
		Sha256  string `json:"sha256"`
		Size    int64  `json:"size"`
		Type    string `json:"type"`
		UID     string `json:"uid"`
		Uname   string `json:"uname"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

type SCADetailsListResp struct {
	AffectedItems []struct {
		Compliance []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"compliance"`
		Condition   string `json:"condition"`
		Description string `json:"description"`
		File        string `json:"file"`
		ID          int64  `json:"id"`
		PolicyID    string `json:"policy_id"`
		Rationale   string `json:"rationale"`
		References  string `json:"references"`
		Remediation string `json:"remediation"`
		Result      string `json:"result"`
		Title       string `json:"title"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
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
