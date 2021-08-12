package clamav

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
)

type (
	ClamAVResp struct {
		Version struct {
			Daily      string `json:"daily"`
			Main       string `json:"main"`
			Bytecode   string `json:"bytecode"`
			Signatures string `json:"signatures"`
			Clamav     string `json:"clamav"`
		} `json:"version"`
	}
)

//var client = resty.New().SetTimeout(time.Second * 60).SetDebug(false)

//病毒库版本
func GetClamAV(apiKey *request.ApiKey) (list *ClamAVResp, err error) {
	url := fmt.Sprintf("%v%v", apiKey.Addr, _const.OPNSENSE_CLAMAV_INFO_URL)
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
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
