package agents

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/wazuh/model"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"strings"
	"time"
)

const (
	agent_api_url                = "/agents"
	agent_summary_status_api_url = "/agents/summary/status?pretty=true"
	agent_scan_url               = "/syscheck"
	agent_sca_url                = "/sca/"
	agent_sca_details_url        = "/sca/%s/checks/%s"
	agent_syscheck_url           = "/syscheck/"
	agent_ciscat_url             = "/ciscat/%s/results"
	agent_vulnerability_url      = "/vulnerability/%s"
)

var paramString = `{
    "params":{
        "index":"wazuh-alerts-*",
        "body":{
            "version":true,
            "_source":{"excludes":["@timestamp"]},
            "aggs":{
                "2":{
                    "date_histogram":{
                        "field":"timestamp",
                        "fixed_interval":"1s",
                        "time_zone":"Asia/Shanghai",
                        "min_doc_count":1
                    }
                }
            },
            "size":500,
            "from":1,
            "sort":[
                {
                    "timestamp":{
                        "order":"desc",
                        "unmapped_type":"boolean"
                    }
                }
            ],
            "stored_fields":[
                "*"
            ],
            "script_fields":{

            },
            "docvalue_fields":[
                {
                    "field":"@timestamp",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.createdAt",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.end",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.resource.instanceDetails.launchTime",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.service.eventFirstSeen",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.service.eventLastSeen",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.start",
                    "format":"date_time"
                },
                {
                    "field":"data.aws.updatedAt",
                    "format":"date_time"
                },
                {
                    "field":"data.timestamp",
                    "format":"date_time"
                },
                {
                    "field":"data.vulnerability.published",
                    "format":"date_time"
                },
                {
                    "field":"data.vulnerability.updated",
                    "format":"date_time"
                },
                {
                    "field":"syscheck.mtime_after",
                    "format":"date_time"
                },
                {
                    "field":"syscheck.mtime_before",
                    "format":"date_time"
                },
                {
                    "field":"timestamp",
                    "format":"date_time"
                }
            ],
            "query":{
                "bool":{
                    "must":[

                    ],
                    "filter":[
                        {
                            "match_all":{

                            }
                        },
                        {
                            "match_phrase":{
                                "manager.name":{
                                    "query":"wauzh.novalocal"
                                }
                            }
                        },
                        {
                            "match_phrase":{
                                "rule.groups":{
                                    "query":"vulnerability-detector"
                                }
                            }
                        },
                        {
                            "match_phrase":{
                                "agent.id":{
                                    "query":"001"
                                }
                            }
                        },
                        {
                            "match_phrase":{
                                "data.vulnerability.severity":"Critical"
                            }
                        },
                        {
                            "range":{
                                "timestamp":{
                                    "gte":"2021-09-07T10:37:15.711Z",
                                    "lte":"2021-09-08T10:37:15.711Z",
                                    "format":"strict_date_optional_time"
                                }
                            }
                        },
						{
							"bool":{
								"should":[
									{
										"match_phrase":{
											"rule.groups":"active_response"
										}
									}
								],
								"minimum_should_match":1
							}
						}
                    ],
                    "should":[

                    ],
                    "must_not":[

                    ]
                }
            },
            "highlight":{
                "pre_tags":[
                    "@kibana-highlighted-field@"
                ],
                "post_tags":[
                    "@/kibana-highlighted-field@"
                ],
                "fields":{
                    "*":{

                    }
                },
                "fragment_size":2147483647
            }
        },
        "preference":1631083814638
    }
}`

type matchPhrase struct {
	ManagerName *struct {
		Query string `json:"query"`
	} `json:"manager.name,omitempty"`
	RuleGroups *struct {
		Query string `json:"query"`
	} `json:"rule.groups,omitempty"`
	AgentId *struct {
		Query *string `json:"query"`
	} `json:"agent.id,omitempty"`
	SyscheckPath              string  `json:"syscheck.path,omitempty"`
	DataVulnerabilitySeverity *string `json:"data.vulnerability.severity,omitempty"`
}
type multiMatch struct {
	Lenient bool   `json:"lenient"`
	Query   string `json:"query"`
	Type    string `json:"type"`
}
type filter struct {
	MatchAll *struct {
	} `json:"match_all,omitempty"`
	Exists *struct {
		Field string `json:"field,omitempty"`
	} `json:"exists,omitempty"`
	MultiMatch  *multiMatch  `json:"multi_match"`
	MatchPhrase *matchPhrase `json:"match_phrase,omitempty"`
	Bool        *struct {
		Should []struct {
			MatchPhrase *struct {
				RuleGroups string `json:"rule.groups,omitempty"`
			} `json:"match_phrase,omitempty"`
		} `json:"should,omitempty"`
		MinimumShouldMatch int `json:"minimum_should_match,omitempty"`
	} `json:"bool,omitempty"`
	Range *struct {
		Timestamp *struct {
			Gte    time.Time `json:"gte"`
			Lte    time.Time `json:"lte"`
			Format *string   `json:"format"`
		} `json:"timestamp"`
	} `json:"range,omitempty"`
}
type esParams struct {
	Params struct {
		Index string `json:"index"`
		Body  struct {
			Version bool `json:"version"`
			Source  struct {
				Excludes []string `json:"excludes"`
			} `json:"_source"`
			Aggs struct {
				Field1 struct {
					DateHistogram struct {
						Field         string `json:"field"`
						FixedInterval string `json:"fixed_interval"`
						TimeZone      string `json:"time_zone"`
						MinDocCount   int    `json:"min_doc_count"`
					} `json:"date_histogram"`
				} `json:"2"`
			} `json:"aggs"`
			Size int `json:"size"`
			From int `json:"from,omitempty"`
			Sort []struct {
				Timestamp struct {
					Order        string `json:"order"`
					UnmappedType string `json:"unmapped_type"`
				} `json:"timestamp"`
			} `json:"sort"`
			StoredFields []string `json:"stored_fields"`
			ScriptFields struct {
			} `json:"script_fields"`
			DocvalueFields []struct {
				Field  string `json:"field"`
				Format string `json:"format"`
			} `json:"docvalue_fields"`
			Query struct {
				Bool struct {
					Must    []interface{} `json:"must"`
					Filter  []filter      `json:"filter"`
					Should  []interface{} `json:"should"`
					MustNot []interface{} `json:"must_not"`
				} `json:"bool"`
			} `json:"query"`
			Highlight struct {
				PreTags  []string `json:"pre_tags"`
				PostTags []string `json:"post_tags"`
				Fields   struct {
					Field1 struct {
					} `json:"*"`
				} `json:"fields"`
				FragmentSize int `json:"fragment_size"`
			} `json:"highlight"`
		} `json:"body"`
		Preference int64 `json:"preference"`
	} `json:"params"`
}

func Info(req *request.Request, id string) (*ResInfo, error) {
	req.Method = "get"
	req.Path = agent_api_url
	req.Params = map[string]interface{}{
		"agents_list": id,
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	info := &ResInfo{}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &info)
	return info, err
}
func Statistics(req *request.Request) (*StatisticsResp, error) {

	req.Method = "get"
	req.Path = agent_summary_status_api_url
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	list := &StatisticsResp{}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}
func Delete(req *request.Request, ids []string) error {
	req.Method = "delete"
	req.Path = agent_api_url
	req.Params = map[string]interface{}{
		"agents_list": strings.Join(ids, ","),
		"status":      "all",
		"older_than":  "0s",
		"pretty":      true,
		"purge":       true, //从密钥库中删除agent
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	list := &StatisticsResp{}
	if resp.Error != 0 {
		return fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	if err == nil {
		_ = deleteAgent(ids)
	}
	return err
}
func List(req *request.Request, args *ListReq) (*ListResp, error) {

	if args.AdminUserId != 0 {
		args.Group = fmt.Sprintf("admin_%v", args.AdminUserId)
	} else if args.UserId != 0 {
		args.Group = fmt.Sprintf("user_%v", args.UserId)
	} else {
		args.Group = ""
	}

	req.Method = "get"
	req.Path = agent_api_url
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &ListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	if err != nil {
		return nil, err
	}
	for idx, item := range list.AffectedItems {
		list.AffectedItems[idx].Remake, _ = getInfo(item.ID)
	}
	return list, err
}
func Scan(req *request.Request, agent []string) error {
	req.Method = "put"
	req.Path = agent_scan_url
	req.Params = map[string]interface{}{
		"agents_list": strings.Join(agent, ","),
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Error != 0 {
		return fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	return err
}

func Check(req *request.Request, agent string) error {

	req.Method = "put"
	req.Path = agent_api_url + "/" + agent + "/restart"
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	list := &StatisticsResp{}
	if resp.Error != 0 {
		return fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return err

}
func Update(args UpdateReq) error {

	return updateAgent(&HIDSAgent{AgentId: args.ID, Remake: args.Remake})
}

//SCADetailsList 合规基线详情
func SCADetailsList(req *request.Request, args SCADetailsListReq) (*SCADetailsListResp, error) {
	req.Method = "get"
	req.Path = fmt.Sprintf(agent_sca_details_url, args.Agent, args.Policy)
	if args.Limit == 0 {
		args.Limit = 20
	}
	req.Params = model.ToMap(map[string]interface{}{
		"limit":  args.Limit,
		"offset": args.Offset,
		"result": args.Result,
	})
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SCADetailsListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

//SCAList Security configuration assessment 合规基线列表
func SCAList(req *request.Request, args SCAListReq) (*SCAListResp, error) {

	req.Method = "get"
	req.Path = agent_sca_url + args.Agent
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SCAListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func SysCheckList(req *request.Request, args SysCheckListReq) (*SysCheckListResp, error) {

	req.Method = "get"
	req.Path = agent_syscheck_url + args.Agent
	req.Params = map[string]interface{}{
		"limit":  args.Limit,
		"offset": args.Offset,
		"type":   "file",
	}
	if args.File != "" {
		req.Params["file"] = args.File
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &SysCheckListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func CiscatList(req *request.Request, agent string) (*CiscatListResp, error) {

	req.Method = "get"
	req.Path = fmt.Sprintf(agent_ciscat_url, agent)
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &CiscatListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

func VulnerabilityList(req *request.Request, agent string) (*VulnerabilityListResp, error) {

	req.Method = "get"
	req.Path = fmt.Sprintf(agent_vulnerability_url, agent)
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf("主机防护服务异常：%s", resp.Message)
	}
	list := &VulnerabilityListResp{}
	bytes, _ := json.Marshal(resp.Data)
	err = json.Unmarshal(bytes, &list)
	return list, err
}

//VulnerabilityESList 模拟登录 es 漏洞风险
func VulnerabilityESList(req *request.Request, args ESListReq) (*VulnerabilityHitsResp, error) {

	req.Method = "post"
	req.Path = "/internal/search/es"

	var paramStruct esParams
	_ = json.Unmarshal([]byte(paramString), &paramStruct)

	newFilter := paramStruct.Params.Body.Query.Bool.Filter[:3]

	newFilter[2].MatchPhrase.RuleGroups.Query = "vulnerability-detector"

	timeFilter := paramStruct.Params.Body.Query.Bool.Filter[5]
	if args.Agent != "" { //指定agent
		paramStruct.Params.Body.Query.Bool.Filter[3].MatchPhrase.AgentId.Query = &args.Agent
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[3])
	}
	if args.Severity != "" { //指定等级
		paramStruct.Params.Body.Query.Bool.Filter[4].MatchPhrase.DataVulnerabilitySeverity = &args.Severity
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[4])
	}
	if args.Start != 0 && args.End != 0 && args.Start < args.End {
		timeFilter.Range.Timestamp.Gte = time.Unix(args.Start, 0)
		timeFilter.Range.Timestamp.Lte = time.Unix(args.End, 0)
		newFilter = append(newFilter, timeFilter)
	}
	if args.Limit == 0 {
		args.Limit = 20
	}

	paramStruct.Params.Body.Size = 1
	paramStruct.Params.Body.From = 0
	paramStruct.Params.Body.Query.Bool.Filter = newFilter

	resp, err := req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	vuls := &vulnerabilityESList{}

	err = json.Unmarshal(resp, &vuls)
	if err != nil {
		return nil, err
	}

	if vuls.StatusCode == 401 {
		return nil, fmt.Errorf(vuls.Message)
	}

	if vuls.RawResponse.Hits.Total > 0 {

		timestamp := vuls.RawResponse.Hits.Hits[0].Source.Timestamp
		timestamp = timestamp[:13]
		start, _ := time.Parse("2006-01-02T15", timestamp)

		//设置时区 6小时
		timeFilter.Range.Timestamp.Gte = start.Add(-8 * time.Hour)
		timeFilter.Range.Timestamp.Lte = start.Add(-7 * time.Hour)
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, timeFilter)
	} else {

		return &vuls.RawResponse.Hits, nil
	}

	paramStruct.Params.Body.Size = args.Limit
	paramStruct.Params.Body.From = args.Offset
	resp, err = req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	vuls = &vulnerabilityESList{}
	err = json.Unmarshal(resp, &vuls)
	if err != nil {
		return nil, err
	}

	if vuls.StatusCode == 401 {
		return nil, fmt.Errorf(vuls.Message)
	}
	return &vuls.RawResponse.Hits, nil
}

//VirusESList 模拟登录 es 病毒管理
func VirusESList(req *request.Request, args ESListReq) (*VirusESHitsListResp, error) {

	req.Method = "post"
	req.Path = "/internal/search/es"

	var paramStruct esParams
	_ = json.Unmarshal([]byte(paramString), &paramStruct)

	//过滤
	paramStruct.Params.Body.Query.Bool.MustNot = []interface{}{
		map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"data.virustotal.positives": "0",
			},
		},
	}
	newFilter := paramStruct.Params.Body.Query.Bool.Filter[:3]
	newFilter[2].MatchPhrase.RuleGroups.Query = "virustotal"

	timeFilter := paramStruct.Params.Body.Query.Bool.Filter[5]
	if args.Agent != "" { //指定agent
		paramStruct.Params.Body.Query.Bool.Filter[3].MatchPhrase.AgentId.Query = &args.Agent
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[3])
	}
	if args.Start != 0 && args.End != 0 && args.Start < args.End {
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Gte = time.Unix(args.Start, 0)
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Lte = time.Unix(args.End, 0)
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[5])
	}
	if args.Limit == 0 {
		args.Limit = 20
	}
	paramStruct.Params.Body.Size = 1
	paramStruct.Params.Body.From = 0
	paramStruct.Params.Body.Query.Bool.Filter = newFilter

	resp, err := req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	virus := &virusESList{}

	err = json.Unmarshal(resp, &virus)
	if err != nil {
		return nil, err
	}

	if virus.StatusCode == 401 {
		return nil, fmt.Errorf(virus.Message)
	}

	if virus.RawResponse.Hits.Total > 0 {

		timestamp := virus.RawResponse.Hits.Hits[0].Source.Timestamp
		timestamp = timestamp[:13]
		start, _ := time.Parse("2006-01-02T15", timestamp)

		//设置时区 6小时
		timeFilter.Range.Timestamp.Gte = start.Add(-8 * time.Hour)
		timeFilter.Range.Timestamp.Lte = start.Add(-7 * time.Hour)
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, timeFilter)
	} else {

		return &virus.RawResponse.Hits, nil
	}

	paramStruct.Params.Body.Size = args.Limit
	paramStruct.Params.Body.From = args.Offset
	resp, err = req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	virus = &virusESList{}
	err = json.Unmarshal(resp, &virus)
	if err != nil {
		return nil, err
	}

	if virus.StatusCode == 401 {
		return nil, fmt.Errorf(virus.Message)
	}
	return &virus.RawResponse.Hits, nil
}

//SysCheckESList 文件监控列表
func SysCheckESList(req *request.Request, args ESListReq) (*SysCheckESHitsListResp, error) {

	req.Method = "post"
	req.Path = "/internal/search/es"

	var paramStruct esParams
	_ = json.Unmarshal([]byte(paramString), &paramStruct)
	//过滤
	paramStruct.Params.Body.Query.Bool.MustNot = []interface{}{
		map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"agent.id": "000",
			},
		},
	}
	newFilter := paramStruct.Params.Body.Query.Bool.Filter[:3]
	newFilter[2].MatchPhrase.RuleGroups.Query = "syscheck"
	timeFilter := paramStruct.Params.Body.Query.Bool.Filter[5]
	if args.Agent != "" { //指定agent
		paramStruct.Params.Body.Query.Bool.Filter[3].MatchPhrase.AgentId.Query = &args.Agent
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[3])
	}
	if args.Path != "" { //指定文件路径
		newFilter = append(newFilter, filter{MatchPhrase: &matchPhrase{SyscheckPath: args.Path}})
	}
	if args.Start != 0 && args.End != 0 && args.Start < args.End {
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Gte = time.Unix(args.Start, 0)
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Lte = time.Unix(args.End, 0)
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[5])
	}
	if args.Limit == 0 {
		args.Limit = 20
	}
	paramStruct.Params.Body.Size = 1
	paramStruct.Params.Body.From = 0
	paramStruct.Params.Body.Query.Bool.Filter = newFilter

	resp, err := req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	syscheck := &sysCheckESList{}

	err = json.Unmarshal(resp, &syscheck)
	if err != nil {
		return nil, err
	}

	if syscheck.StatusCode == 401 {
		return nil, fmt.Errorf(syscheck.Message)
	}

	if syscheck.RawResponse.Hits.Total > 0 {

		timestamp := syscheck.RawResponse.Hits.Hits[0].Source.Timestamp
		timestamp = timestamp[:13]
		start, _ := time.Parse("2006-01-02T15", timestamp)

		//设置时区 6小时
		timeFilter.Range.Timestamp.Gte = start.Add(-8 * time.Hour)
		timeFilter.Range.Timestamp.Lte = start.Add(-7 * time.Hour)
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, timeFilter)
	} else {

		return &syscheck.RawResponse.Hits, nil
	}

	paramStruct.Params.Body.Size = args.Limit
	paramStruct.Params.Body.From = args.Offset
	resp, err = req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	syscheck = &sysCheckESList{}
	err = json.Unmarshal(resp, &syscheck)
	if err != nil {
		return nil, err
	}

	if syscheck.StatusCode == 401 {
		return nil, fmt.Errorf(syscheck.Message)
	}
	return &syscheck.RawResponse.Hits, nil
}

//ATTCKESList 安全事件监控列表
func ATTCKESList(req *request.Request, args ESListReq) (*ATTCKESHitsListResp, error) {

	req.Method = "post"
	req.Path = "/internal/search/es"

	var paramStruct esParams
	_ = json.Unmarshal([]byte(paramString), &paramStruct)
	//过滤
	paramStruct.Params.Body.Query.Bool.MustNot = []interface{}{
		map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"agent.id": "000",
			},
		},
	}
	newFilter := paramStruct.Params.Body.Query.Bool.Filter[:3]
	timeFilter := paramStruct.Params.Body.Query.Bool.Filter[5]
	newFilter[2].MatchPhrase = nil
	newFilter[2].Exists = &struct {
		Field string `json:"field,omitempty"`
	}{Field: "rule.mitre.id"}

	if args.Agent != "" { //指定agent
		paramStruct.Params.Body.Query.Bool.Filter[3].MatchPhrase.AgentId.Query = &args.Agent
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[3])
	}
	if args.Start != 0 && args.End != 0 && args.Start < args.End {
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Gte = time.Unix(args.Start, 0)
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Lte = time.Unix(args.End, 0)
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[5])
	}
	if args.Limit == 0 {
		args.Limit = 20
	}
	paramStruct.Params.Body.Size = 1
	paramStruct.Params.Body.From = 0
	paramStruct.Params.Body.Query.Bool.Filter = newFilter
	if args.Warning {
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, filter{
			MultiMatch: &multiMatch{
				Type:    "best_fields",
				Query:   "Brute Force",
				Lenient: true,
			},
		})
	}

	resp, err := req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	attck := &aTTCKESList{}

	err = json.Unmarshal(resp, &attck)
	if err != nil {
		return nil, err
	}

	if attck.StatusCode == 401 {
		return nil, fmt.Errorf(attck.Message)
	}

	if attck.RawResponse.Hits.Total > 0 {

		timestamp := attck.RawResponse.Hits.Hits[0].Source.Timestamp
		timestamp = timestamp[:13]
		start, _ := time.Parse("2006-01-02T15", timestamp)

		//设置时区 6小时
		timeFilter.Range.Timestamp.Gte = start.Add(-8 * time.Hour)
		timeFilter.Range.Timestamp.Lte = start.Add(-7 * time.Hour)
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, timeFilter)
	} else {

		return &attck.RawResponse.Hits, nil
	}

	paramStruct.Params.Body.Size = args.Limit
	paramStruct.Params.Body.From = args.Offset
	resp, err = req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	attck = &aTTCKESList{}
	err = json.Unmarshal(resp, &attck)
	if err != nil {
		return nil, err
	}

	if attck.StatusCode == 401 {
		return nil, fmt.Errorf(attck.Message)
	}
	return &attck.RawResponse.Hits, nil
}

//InvadeThreatESList 入侵威胁列表
func InvadeThreatESList(req *request.Request, args ESListReq) (*InvadeThreatESHitsListResp, error) {

	req.Method = "post"
	req.Path = "/internal/search/es"

	var paramStruct esParams
	_ = json.Unmarshal([]byte(paramString), &paramStruct)

	newFilter := paramStruct.Params.Body.Query.Bool.Filter[:2]
	newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[6])
	timeFilter := paramStruct.Params.Body.Query.Bool.Filter[5]
	if args.Agent != "" { //指定agent
		paramStruct.Params.Body.Query.Bool.Filter[3].MatchPhrase.AgentId.Query = &args.Agent
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[3])
	}
	if args.Start != 0 && args.End != 0 && args.Start < args.End {
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Gte = time.Unix(args.Start, 0)
		paramStruct.Params.Body.Query.Bool.Filter[5].Range.Timestamp.Lte = time.Unix(args.End, 0)
		newFilter = append(newFilter, paramStruct.Params.Body.Query.Bool.Filter[5])
	}
	if args.Limit == 0 {
		args.Limit = 20
	}
	paramStruct.Params.Body.Size = 1
	paramStruct.Params.Body.From = 0
	paramStruct.Params.Body.Query.Bool.Filter = newFilter

	resp, err := req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	invade := &invadeThreatESList{}

	err = json.Unmarshal(resp, &invade)
	if err != nil {
		return nil, err
	}

	if invade.StatusCode == 401 {
		return nil, fmt.Errorf(invade.Message)
	}

	if invade.RawResponse.Hits.Total > 0 {

		timestamp := invade.RawResponse.Hits.Hits[0].Source.Timestamp
		timestamp = timestamp[:13]
		start, _ := time.Parse("2006-01-02T15", timestamp)

		//设置时区 6小时
		timeFilter.Range.Timestamp.Gte = start.Add(-8 * time.Hour)
		timeFilter.Range.Timestamp.Lte = start.Add(-7 * time.Hour)
		paramStruct.Params.Body.Query.Bool.Filter = append(paramStruct.Params.Body.Query.Bool.Filter, timeFilter)
	} else {

		return &invade.RawResponse.Hits, nil
	}

	paramStruct.Params.Body.Size = args.Limit
	paramStruct.Params.Body.From = args.Offset
	resp, err = req.Do2(paramStruct)
	if err != nil {
		return nil, err
	}
	invade = &invadeThreatESList{}
	err = json.Unmarshal(resp, &invade)
	if err != nil {
		return nil, err
	}

	if invade.StatusCode == 401 {
		return nil, fmt.Errorf(invade.Message)
	}
	return &invade.RawResponse.Hits, nil
}
