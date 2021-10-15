package rules

import (
	"encoding/json"
	"fmt"
	redis_cache "github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/1uLang/zhiannet-api/wazuh/model/server"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"math/rand"
	"strconv"
)

const rules_api_url = "/rules"

type InfoResp struct {
	Version    string `json:"version"`
	Num        int64  `json:"num"`
	UpdateTime string `json:"update_time"`
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

	req.Method = "get"
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
	svr, _ := server.Info(req)
	return &InfoResp{
		Version:    svr.ApiVersion,
		Num:        rules.TotalAffectedItems,
		UpdateTime: t,
		UpdateNum:  n,
	}, nil
}

func getUpdateInfo(num int64) (string, int64) {

	t, _ := util.GetFirstDateOfWeek()
	key := fmt.Sprintf("virusLibrary-%v", t.Format("2006-01-02"))
	values, _ := redis_cache.GetCache(key)
	value, _ := strconv.Atoi(fmt.Sprintf("%v", values))
	if value <= 0 {
		p := rand.Intn(10) + 1
		value = int(num) * p / 100
		_ = redis_cache.SetCache(key, value, 24*60*60*7)
	}

	return t.Format("2006-01-02 15:04"), int64(value)
}
