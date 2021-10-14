package rules

import (
	"encoding/json"
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"math/rand"
	"time"
)

const rules_api_url = "/rules"

type InfoResp struct {
	Version    string `json:"version"`
	Num        int64  `json:"num"`
	UpdateTime int64  `json:"update_time"`
	UpdateNum  int64  `json:"update_num"`
}

type rulesResp struct {
	AffectedItems []struct {
		ConfigSum string `json:"configSum"`
		Count     int64  `json:"count"`
		MergedSum string `json:"mergedSum"`
		Name      string `json:"name"`
	} `json:"affected_items"`
	FailedItems        []interface{} `json:"failed_items"`
	TotalAffectedItems int64         `json:"total_affected_items"`
	TotalFailedItems   int64         `json:"total_failed_items"`
}

func Info(req *request.Request) (*InfoResp, error) {

	req.Method = "post"
	req.Path = rules_api_url
	req.Params = map[string]interface{}{
		"offset": 0,
		"limit":  1,
	}
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Error != 0 {
		return nil, fmt.Errorf(resp.Message)
	}
	rules := rulesResp{}
	bytes, _ := json.Marshal(resp.Data)
	_ = json.Unmarshal(bytes, &rules)
	t, n := getUpdateInfo(rules.TotalAffectedItems)
	return &InfoResp{
		Version:    "virusLibrary_" + time.Now().Format("20060102"),
		Num:        rules.TotalAffectedItems,
		UpdateTime: t,
		UpdateNum:  n,
	}, nil
}

func getUpdateInfo(num int64) (int64, int64) {

	key := "virusLibrary"

	//当周一 0点
	tm := time.Now()
	week := tm.Weekday()
	if week != time.Monday {
		if week == time.Sunday {
			tm.AddDate(0, 0, -6)
		} else {
			tm.AddDate(0, 0, 1-int(week))
		}
	}
	t, _ := time.Parse(tm.Format("20060102"), "20060102")

	value, _ := redis_cache.GetInt(key)
	if value <= 0 {
		p := rand.Intn(10) + 1
		value = int(num) * p / 100
		_ = redis_cache.SetCache(key, value, 0)
	}

	return t.Unix(), int64(value)
}
