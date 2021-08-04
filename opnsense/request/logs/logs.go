package logs

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
)

type (
	LogReq struct {
		Current      string //页数
		RowCount     string //每页条数
		SearchPhrase string //关键词
	}
	LogListResp struct {
		Current  int    `json:"current"`
		Filters  string `json:"filters"`
		Origin   string `json:"origin"`
		RowCount int    `json:"rowCount"`
		Rows     []struct {
			Line        string `json:"line"`
			Parser      string `json:"parser"`
			ProcessName string `json:"process_name"`
			Rnum        int    `json:"rnum"`
			Timestamp   string `json:"timestamp"`
		} `json:"rows"`
		Total     int `json:"total"`
		TotalRows int `json:"total_rows"`
	}
	ClearLogResp struct {
		Status string `json:"status"`
	}
)

//var client = resty.New().SetTimeout(time.Second * 60).SetDebug(false)

//获取日志
func GetLogsList(req *LogReq, apiKey *request.ApiKey) (list *LogListResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_LOGS_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetFormData(map[string]string{
			"current":      req.Current, //
			"rowCount":     req.RowCount,
			"searchPhrase": req.SearchPhrase,
		}).SetCookie(&http.Cookie{
		Name:  "PHPSESSID",
		Value: apiKey.Cookie,
	}).
		//Get("https://182.150.0.109:5443/firewall_nat_edit.php")
		//Get("https://182.150.0.109:5443/api/diagnostics/log/core/suricata")
		Post(url)
	//fmt.Println("edge_logs_server query list == ",string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//清除所有日志
func ClearLog(apiKey *request.ApiKey) (res bool, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_CLEAR_LOGS_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		//SetBasicAuth(apiKey.Username, apiKey.Password).
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Post(url)
	//fmt.Println(string(resp.Body()), err)
	clearRes := ClearLogResp{}
	err = json.Unmarshal(resp.Body(), &clearRes)
	if err != nil {
		return res, err
	}
	if clearRes.Status == "ok" {
		res = true
	}
	return res, err
}
