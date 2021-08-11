package conversation

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
)

type (
	ConReq struct {
		Current      string //页数
		RowCount     string //每页条数
		SearchPhrase string //关键词
	}
	ListResp struct {
		Current  int `json:"current"`
		RowCount int `json:"rowCount"`
		Rows     []struct {
			Proto   string      `json:"proto"`
			Dir     string      `json:"dir"`
			SrcAddr string      `json:"src_addr"`
			SrcPort string      `json:"src_port"`
			DstAddr string      `json:"dst_addr"`
			DstPort string      `json:"dst_port"`
			GwAddr  interface{} `json:"gw_addr"`
			GwPort  interface{} `json:"gw_port"`
			State   string      `json:"state"`
			Age     int         `json:"age"`
			Expire  int         `json:"expire"`
			Pkts    int         `json:"pkts"`
			Bytes   int         `json:"bytes"`
			Avg     int         `json:"avg"`
			Rule    string      `json:"rule"`
			Label   string      `json:"label"`
			Descr   string      `json:"descr"`
		} `json:"rows"`
		Total int `json:"total"`
	}
)

//var client = resty.New().SetTimeout(time.Second * 60).SetDebug(false)

//获取日志
func GetList(req *ConReq, apiKey *request.ApiKey) (list *ListResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_DIAGNOSTICS_LIST_URL)
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
