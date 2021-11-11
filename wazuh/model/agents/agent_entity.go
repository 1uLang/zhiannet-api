package agents

type ResInfo struct {
	Affected_items []struct {
		Group []string `json:"group"`
		IP    string   `json:"ip"`
		Name  string   `json:"name"`
	} `json:"affected_items"`
}
type StatisticsResp struct {
	Active         int `json:"active"`
	Disconnected   int `json:"disconnected"`
	NeverConnected int `json:"never_connected"`
	Pending        int `json:"pending"`
	Total          int `json:"total"`
}
type ListReq struct {
	Group  string `json:"group,omitempty"`
	IP     string `json:"ip,omitempty"`
	Limit  int64  `json:"limit,omitempty"`
	Offset int64  `json:"offset,omitempty"`

	UserId      int64 `json:"-"`
	AdminUserId int64 `json:"-"`
}
type UpdateReq struct {
	ID     string `json:"group"`
	Remake string `json:"ip"`

	UserId      uint64 `json:"-"`
	AdminUserId uint64 `json:"-"`
}
type SysCheckListReq struct {
	Agent  string `json:"-"`
	File   string `json:"file,omitempty"`
	Limit  int64  `json:"limit,omitempty"`
	Offset int64  `json:"offset,omitempty"`
}
type Affected struct {
}
type ListResp struct {
	AffectedItems []struct {
		ConfigSum     string   `json:"configSum"`
		DateAdd       string   `json:"dateAdd"` //创建时间
		Group         []string `json:"group"`
		ID            string   `json:"id"`
		IP            string   `json:"ip"`
		LastKeepAlive string   `json:"lastKeepAlive"` //最后在线时间
		Manager       string   `json:"manager"`
		MergedSum     string   `json:"mergedSum"`
		Name          string   `json:"name"`
		NodeName      string   `json:"node_name"`
		Remake        string   `json:"remake"`
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
		Command     string `json:"command"`
		Registry    string `json:"registry"`
		Bopen       bool   `json:"bOpen"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

type SCAListReq struct {
	Agent  string `json:"-"`
	Limit  int64  `json:"limit,omitempty"`
	Offset int64  `json:"offset,omitempty"`
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
		AgentID     string `json:"agent_id"`
		AgentIP     string `json:"agent_ip"`
		AgentName   string `json:"agent_name"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

type SCADetailsListReq struct {
	Agent  string
	Policy string
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Result string `json:"result"`
}
type ESListReq struct {
	Agent    string
	Severity string //Critical   High  Medium Low
	Path     string
	Start    int64
	End      int64
	Limit    int
	Offset   int
	Warning  bool `json:"-"`
}
type VulnerabilityHitsResp struct {
	Total    int         `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []struct {
		Index   string      `json:"_index"`
		Type    string      `json:"_type"`
		Id      string      `json:"_id"`
		Version int         `json:"_version"`
		Score   interface{} `json:"_score"`
		Source  struct {
			Predecoder struct {
			} `json:"predecoder"`
			Cluster struct {
				Name string `json:"name"`
			} `json:"cluster"`
			Agent struct {
				Ip   string `json:"ip"`
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"agent"`
			Manager struct {
				Name string `json:"name"`
			} `json:"manager"`
			Data struct {
				Vulnerability struct {
					Severity string `json:"severity"`
					Package  struct {
						Condition    string `json:"condition"`
						Name         string `json:"name"`
						Source       string `json:"source,omitempty"`
						Version      string `json:"version"`
						Architecture string `json:"architecture"`
					} `json:"package"`
					References   []string `json:"references"`
					CveVersion   string   `json:"cve_version"`
					Assigner     string   `json:"assigner"`
					Published    string   `json:"published"`
					CweReference string   `json:"cwe_reference"`
					Title        string   `json:"title"`
					Rationale    string   `json:"rationale,omitempty"`
					Cve          string   `json:"cve"`
					State        string   `json:"state"`
					Cvss         struct {
						Cvss2 struct {
							BaseScore string `json:"base_score"`
							Vector    struct {
								IntegrityImpact       string `json:"integrity_impact"`
								ConfidentialityImpact string `json:"confidentiality_impact"`
								Availability          string `json:"availability"`
								AttackVector          string `json:"attack_vector"`
								AccessComplexity      string `json:"access_complexity"`
								Authentication        string `json:"authentication"`
							} `json:"vector"`
						} `json:"cvss2"`
						Cvss3 struct {
							BaseScore string `json:"base_score"`
							Vector    struct {
								UserInteraction       string `json:"user_interaction"`
								IntegrityImpact       string `json:"integrity_impact"`
								Scope                 string `json:"scope"`
								ConfidentialityImpact string `json:"confidentiality_impact"`
								Availability          string `json:"availability"`
								AttackVector          string `json:"attack_vector"`
								AccessComplexity      string `json:"access_complexity"`
								PrivilegesRequired    string `json:"privileges_required"`
							} `json:"vector"`
						} `json:"cvss3,omitempty"`
					} `json:"cvss"`
					Updated            string   `json:"updated"`
					BugzillaReferences []string `json:"bugzilla_references,omitempty"`
				} `json:"vulnerability"`
			} `json:"data"`
			Sampledata bool `json:"@sampledata"`
			Rule       struct {
				Firedtimes  int      `json:"firedtimes"`
				Mail        bool     `json:"mail"`
				Level       int      `json:"level"`
				PciDss      []string `json:"pci_dss"`
				Tsc         []string `json:"tsc"`
				Description string   `json:"description"`
				Groups      []string `json:"groups"`
				Id          string   `json:"id"`
				Gdpr        []string `json:"gdpr"`
			} `json:"rule"`
			Location string `json:"location"`
			Id       string `json:"id"`
			Decoder  struct {
				Name string `json:"name"`
			} `json:"decoder"`
			Timestamp string `json:"timestamp"`
		} `json:"_source"`
		Fields struct {
			DataVulnerabilityPublished []string `json:"data.vulnerability.published"`
			DataVulnerabilityUpdated   []string `json:"data.vulnerability.updated"`
			Timestamp                  []string `json:"timestamp"`
		} `json:"fields"`
		Highlight struct {
			AgentId     []string `json:"agent.id"`
			ManagerName []string `json:"manager.name"`
			RuleGroups  []string `json:"rule.groups"`
		} `json:"highlight"`
		Sort []int64 `json:"sort"`
	} `json:"hits"`
}
type vulnerabilityESList struct {
	IsPartial   bool `json:"isPartial"`
	IsRunning   bool `json:"isRunning"`
	RawResponse struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits         VulnerabilityHitsResp `json:"hits"`
		Aggregations struct {
			Field1 struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         int64  `json:"key"`
					DocCount    int    `json:"doc_count"`
				} `json:"buckets"`
			} `json:"2"`
		} `json:"aggregations"`
	} `json:"rawResponse"`
	Total      int    `json:"total"`
	Loaded     int    `json:"loaded"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type VirusESHitsListResp struct {
	Total    int         `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []struct {
		Index   string      `json:"_index"`
		Type    string      `json:"_type"`
		Id      string      `json:"_id"`
		Version int         `json:"_version"`
		Score   interface{} `json:"_score"`
		Source  struct {
			Predecoder struct {
			} `json:"predecoder"`
			Cluster struct {
				Name string `json:"name"`
			} `json:"cluster"`
			Agent struct {
				Ip   string `json:"ip"`
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"agent"`
			Manager struct {
				Name string `json:"name"`
			} `json:"manager"`
			Data struct {
				Virustotal struct {
					Malicious interface{} `json:"malicious"`
					Found     string      `json:"found"`
					Source    struct {
						Sha1    string `json:"sha1"`
						File    string `json:"file"`
						AlertId string `json:"alert_id"`
						Md5     string `json:"md5"`
					} `json:"source"`
					Total     string `json:"total,omitempty"`
					Positives string `json:"positives,omitempty"`
					Permalink string `json:"permalink,omitempty"`
					ScanDate  string `json:"scan_date,omitempty"`
				} `json:"virustotal"`
			} `json:"data"`
			Sampledata bool `json:"@sampledata"`
			Rule       struct {
				Mail        bool     `json:"mail"`
				Level       int      `json:"level"`
				Description string   `json:"description"`
				Groups      []string `json:"groups"`
				Id          string   `json:"id"`
			} `json:"rule"`
			Location string `json:"location"`
			Id       string `json:"id"`
			Decoder  struct {
			} `json:"decoder"`
			Timestamp string `json:"timestamp"`
		} `json:"_source"`
		Fields struct {
			Timestamp []string `json:"timestamp"`
		} `json:"fields"`
		Highlight struct {
			AgentId     []string `json:"agent.id"`
			ManagerName []string `json:"manager.name"`
			RuleGroups  []string `json:"rule.groups"`
		} `json:"highlight"`
		Sort []int64 `json:"sort"`
	} `json:"hits"`
}
type virusESList struct {
	IsPartial   bool `json:"isPartial"`
	IsRunning   bool `json:"isRunning"`
	RawResponse struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits         VirusESHitsListResp `json:"hits"`
		Aggregations struct {
			Field1 struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         int64  `json:"key"`
					DocCount    int    `json:"doc_count"`
				} `json:"buckets"`
			} `json:"2"`
		} `json:"aggregations"`
	} `json:"rawResponse"`
	Total      int    `json:"total"`
	Loaded     int    `json:"loaded"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type SysCheckESHitsListResp struct {
	Total    int         `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []struct {
		Index   string      `json:"_index"`
		Type    string      `json:"_type"`
		Id      string      `json:"_id"`
		Version int         `json:"_version"`
		Score   interface{} `json:"_score"`
		Source  struct {
			Syscheck struct {
				SizeAfter         string   `json:"size_after,omitempty"`
				Md5Before         string   `json:"md5_before,omitempty"`
				Sha256Before      string   `json:"sha256_before,omitempty"`
				Mode              string   `json:"mode"`
				Path              string   `json:"path"`
				Sha1After         string   `json:"sha1_after,omitempty"`
				ChangedAttributes []string `json:"changed_attributes"`
				Arch              string   `json:"arch"`
				Event             string   `json:"event"`
				Md5After          string   `json:"md5_after,omitempty"`
				Sha1Before        string   `json:"sha1_before,omitempty"`
				Sha256After       string   `json:"sha256_after,omitempty"`
				ValueName         string   `json:"value_name,omitempty"`
				UnameAfter        string   `json:"uname_after,omitempty"`
				GnameAfter        string   `json:"gname_after,omitempty"`
				MtimeAfter        string   `json:"mtime_after,omitempty"`
				UidAfter          string   `json:"uid_after,omitempty"`
				GidAfter          string   `json:"gid_after,omitempty"`
				WinPermAfter      []struct {
					Allowed []string `json:"allowed"`
					Name    string   `json:"name"`
				} `json:"win_perm_after,omitempty"`
				MtimeBefore string `json:"mtime_before,omitempty"`
				SizeBefore  string `json:"size_before,omitempty"`
			} `json:"syscheck"`
			Input struct {
				Type string `json:"type"`
			} `json:"input"`
			Agent struct {
				Ip   string `json:"ip"`
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"agent"`
			Manager struct {
				Name string `json:"name"`
			} `json:"manager"`
			Rule struct {
				Mail        bool     `json:"mail"`
				Level       int      `json:"level"`
				PciDss      []string `json:"pci_dss"`
				Hipaa       []string `json:"hipaa"`
				Tsc         []string `json:"tsc"`
				Description string   `json:"description"`
				Groups      []string `json:"groups"`
				Nist80053   []string `json:"nist_800_53"`
				Gdpr        []string `json:"gdpr"`
				Firedtimes  int      `json:"firedtimes"`
				Mitre       struct {
					Technique []string `json:"technique"`
					Id        []string `json:"id"`
					Tactic    []string `json:"tactic"`
				} `json:"mitre"`
				Id    string   `json:"id"`
				Gpg13 []string `json:"gpg13"`
			} `json:"rule"`
			Location string `json:"location"`
			Decoder  struct {
				Name string `json:"name"`
			} `json:"decoder"`
			Id        string `json:"id"`
			FullLog   string `json:"full_log"`
			Timestamp string `json:"timestamp"`
		} `json:"_source"`
		Fields struct {
			Timestamp           []string `json:"timestamp"`
			SyscheckMtimeAfter  []string `json:"syscheck.mtime_after,omitempty"`
			SyscheckMtimeBefore []string `json:"syscheck.mtime_before,omitempty"`
		} `json:"fields"`
		Highlight struct {
			AgentId     []string `json:"agent.id"`
			ManagerName []string `json:"manager.name"`
			RuleGroups  []string `json:"rule.groups"`
		} `json:"highlight"`
		Sort []int64 `json:"sort"`
	} `json:"hits"`
}
type sysCheckESList struct {
	IsPartial   bool `json:"isPartial"`
	IsRunning   bool `json:"isRunning"`
	RawResponse struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits         SysCheckESHitsListResp `json:"hits"`
		Aggregations struct {
			Field1 struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         int64  `json:"key"`
					DocCount    int    `json:"doc_count"`
				} `json:"buckets"`
			} `json:"2"`
		} `json:"aggregations"`
	} `json:"rawResponse"`
	Total      int    `json:"total"`
	Loaded     int    `json:"loaded"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type ATTCKESHitsListResp struct {
	Total    int         `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []struct {
		Index   string      `json:"_index"`
		Type    string      `json:"_type"`
		Id      string      `json:"_id"`
		Version int         `json:"_version"`
		Score   interface{} `json:"_score"`
		Source  struct {
			Input struct {
				Type string `json:"type"`
			} `json:"input"`
			Agent struct {
				Ip   string `json:"ip"`
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"agent"`
			Manager struct {
				Name string `json:"name"`
			} `json:"manager"`
			Data struct {
				Win struct {
					Eventdata struct {
						SubjectLogonId            string `json:"subjectLogonId,omitempty"`
						SubjectDomainName         string `json:"subjectDomainName,omitempty"`
						TargetLinkedLogonId       string `json:"targetLinkedLogonId,omitempty"`
						ImpersonationLevel        string `json:"impersonationLevel,omitempty"`
						AuthenticationPackageName string `json:"authenticationPackageName,omitempty"`
						TargetLogonId             string `json:"targetLogonId,omitempty"`
						LogonProcessName          string `json:"logonProcessName,omitempty"`
						LogonGuid                 string `json:"logonGuid,omitempty"`
						TargetUserName            string `json:"targetUserName,omitempty"`
						KeyLength                 string `json:"keyLength,omitempty"`
						ElevatedToken             string `json:"elevatedToken,omitempty"`
						SubjectUserSid            string `json:"subjectUserSid,omitempty"`
						ProcessId                 string `json:"processId,omitempty"`
						ProcessName               string `json:"processName,omitempty"`
						TargetDomainName          string `json:"targetDomainName,omitempty"`
						TargetUserSid             string `json:"targetUserSid,omitempty"`
						VirtualAccount            string `json:"virtualAccount,omitempty"`
						LogonType                 string `json:"logonType,omitempty"`
						SubjectUserName           string `json:"subjectUserName,omitempty"`
						ServiceType               string `json:"serviceType,omitempty"`
						ImagePath                 string `json:"imagePath,omitempty"`
						StartType                 string `json:"startType,omitempty"`
						ServiceName               string `json:"serviceName,omitempty"`
						IpAddress                 string `json:"ipAddress,omitempty"`
						WorkstationName           string `json:"workstationName,omitempty"`
						IpPort                    string `json:"ipPort,omitempty"`
					} `json:"eventdata"`
					System struct {
						EventID         string `json:"eventID"`
						Keywords        string `json:"keywords"`
						ProviderGuid    string `json:"providerGuid"`
						Level           string `json:"level"`
						Channel         string `json:"channel"`
						Opcode          string `json:"opcode"`
						Message         string `json:"message"`
						Version         string `json:"version"`
						SystemTime      string `json:"systemTime"`
						EventRecordID   string `json:"eventRecordID"`
						ThreadID        string `json:"threadID"`
						Computer        string `json:"computer"`
						Task            string `json:"task"`
						ProcessID       string `json:"processID"`
						SeverityValue   string `json:"severityValue"`
						ProviderName    string `json:"providerName"`
						EventSourceName string `json:"eventSourceName,omitempty"`
					} `json:"system"`
				} `json:"win,omitempty"`
				ExtraData string `json:"extra_data,omitempty"`
			} `json:"data,omitempty"`
			Rule struct {
				Mail        bool     `json:"mail"`
				Level       int      `json:"level"`
				PciDss      []string `json:"pci_dss,omitempty"`
				Hipaa       []string `json:"hipaa,omitempty"`
				Tsc         []string `json:"tsc,omitempty"`
				Description string   `json:"description"`
				Groups      []string `json:"groups"`
				Nist80053   []string `json:"nist_800_53,omitempty"`
				Gdpr        []string `json:"gdpr,omitempty"`
				Firedtimes  int      `json:"firedtimes"`
				Mitre       struct {
					Technique []string `json:"technique"`
					Id        []string `json:"id"`
					Tactic    []string `json:"tactic"`
				} `json:"mitre"`
				Id    string   `json:"id"`
				Gpg13 []string `json:"gpg13,omitempty"`
			} `json:"rule"`
			Location string `json:"location"`
			Decoder  struct {
				Name   string `json:"name"`
				Parent string `json:"parent,omitempty"`
			} `json:"decoder"`
			Id        string `json:"id"`
			Timestamp string `json:"timestamp"`
			FullLog   string `json:"full_log,omitempty"`
			Syscheck  struct {
				Mode              string   `json:"mode"`
				Path              string   `json:"path"`
				ChangedAttributes []string `json:"changed_attributes"`
				UnameAfter        string   `json:"uname_after,omitempty"`
				GnameAfter        string   `json:"gname_after,omitempty"`
				MtimeAfter        string   `json:"mtime_after,omitempty"`
				UidAfter          string   `json:"uid_after,omitempty"`
				GidAfter          string   `json:"gid_after,omitempty"`
				Arch              string   `json:"arch"`
				WinPermAfter      []struct {
					Allowed []string `json:"allowed"`
					Name    string   `json:"name"`
				} `json:"win_perm_after,omitempty"`
				MtimeBefore  string `json:"mtime_before,omitempty"`
				Event        string `json:"event"`
				SizeAfter    string `json:"size_after,omitempty"`
				Md5Before    string `json:"md5_before,omitempty"`
				Sha256Before string `json:"sha256_before,omitempty"`
				Sha1After    string `json:"sha1_after,omitempty"`
				Md5After     string `json:"md5_after,omitempty"`
				Sha1Before   string `json:"sha1_before,omitempty"`
				Sha256After  string `json:"sha256_after,omitempty"`
				ValueName    string `json:"value_name,omitempty"`
				SizeBefore   string `json:"size_before,omitempty"`
			} `json:"syscheck,omitempty"`
		} `json:"_source"`
		Fields struct {
			Timestamp           []string `json:"timestamp"`
			SyscheckMtimeAfter  []string `json:"syscheck.mtime_after,omitempty"`
			SyscheckMtimeBefore []string `json:"syscheck.mtime_before,omitempty"`
		} `json:"fields"`
		Highlight struct {
			AgentId     []string `json:"agent.id"`
			ManagerName []string `json:"manager.name"`
		} `json:"highlight"`
		Sort []int64 `json:"sort"`
	} `json:"hits"`
}
type aTTCKESList struct {
	IsPartial   bool `json:"isPartial"`
	IsRunning   bool `json:"isRunning"`
	RawResponse struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits         ATTCKESHitsListResp `json:"hits"`
		Aggregations struct {
			Field1 struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         int64  `json:"key"`
					DocCount    int    `json:"doc_count"`
				} `json:"buckets"`
			} `json:"2"`
		} `json:"aggregations"`
	} `json:"rawResponse"`
	Total      int    `json:"total"`
	Loaded     int    `json:"loaded"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

type InvadeThreatESHitsListResp struct {
	Total    int         `json:"total"`
	MaxScore interface{} `json:"max_score"`
	Hits     []struct {
		Index   string      `json:"_index"`
		Type    string      `json:"_type"`
		Id      string      `json:"_id"`
		Version int         `json:"_version"`
		Score   interface{} `json:"_score"`
		Source  struct {
			Input struct {
				Type string `json:"type"`
			} `json:"input"`
			Agent struct {
				Ip   string `json:"ip"`
				Name string `json:"name"`
				Id   string `json:"id"`
			} `json:"agent"`
			Manager struct {
				Name string `json:"name"`
			} `json:"manager"`
			Data struct {
				Origin struct {
					Module string `json:"module"`
				} `json:"origin"`
				Version    string `json:"version"`
				Parameters struct {
					ExtraArgs []interface{} `json:"extra_args"`
					Program   string        `json:"program"`
				} `json:"parameters"`
				Command string `json:"command"`
			} `json:"data"`
			Rule struct {
				Firedtimes  int      `json:"firedtimes"`
				Mail        bool     `json:"mail"`
				Level       int      `json:"level"`
				PciDss      []string `json:"pci_dss"`
				Tsc         []string `json:"tsc"`
				Description string   `json:"description"`
				Groups      []string `json:"groups"`
				Id          string   `json:"id"`
				Nist80053   []string `json:"nist_800_53"`
				Gdpr        []string `json:"gdpr"`
			} `json:"rule"`
			Location string `json:"location"`
			Decoder  struct {
				Name string `json:"name"`
			} `json:"decoder"`
			Id        string `json:"id"`
			FullLog   string `json:"full_log"`
			Timestamp string `json:"timestamp"`
		} `json:"_source"`
		Fields struct {
			Timestamp []string `json:"timestamp"`
		} `json:"fields"`
		Highlight struct {
			AgentId     []string `json:"agent.id"`
			ManagerName []string `json:"manager.name"`
			RuleGroups  []string `json:"rule.groups"`
		} `json:"highlight"`
		Sort []int64 `json:"sort"`
	} `json:"hits"`
}
type invadeThreatESList struct {
	IsPartial   bool `json:"isPartial"`
	IsRunning   bool `json:"isRunning"`
	RawResponse struct {
		Took     int  `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards   struct {
			Total      int `json:"total"`
			Successful int `json:"successful"`
			Skipped    int `json:"skipped"`
			Failed     int `json:"failed"`
		} `json:"_shards"`
		Hits         InvadeThreatESHitsListResp `json:"hits"`
		Aggregations struct {
			Field1 struct {
				Buckets []struct {
					KeyAsString string `json:"key_as_string"`
					Key         int64  `json:"key"`
					DocCount    int    `json:"doc_count"`
				} `json:"buckets"`
			} `json:"2"`
		} `json:"aggregations"`
	} `json:"rawResponse"`
	Total      int    `json:"total"`
	Loaded     int    `json:"loaded"`
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}
