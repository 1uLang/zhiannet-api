package logs

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"net/http"
)

type (
	TrafficLogReq struct { //流量日志请求参数
		Addr  string `json:"addr"`
		Level int    `json:"level"`
	}

	LogsReportFlow struct {
		Address   string `xml:"address"`
		Level     string `xml:"level"`
		SortType  string `xml:"sort_type"`
		StartDate string `xml:"start_date"`
		EndDate   string `xml:"end_date"`
		Page      string `xml:"page"`
		Total     string `xml:"total"`
		Report    []struct {
			DateTime     string `xml:"date_time"`
			DateTimeMax  string `xml:"date_time_max"`
			InputBpsMax  string `xml:"input_bps_max"`
			InputPpsMax  string `xml:"input_pps_max"`
			InputBpsAve  string `xml:"input_bps_ave"`
			InputPpsAve  string `xml:"input_pps_ave"`
			OutputBpsMax string `xml:"output_bps_max"`
			OutputPpsMax string `xml:"output_pps_max"`
			OutputBpsAve string `xml:"output_bps_ave"`
			OutputPpsAve string `xml:"output_pps_ave"`
			SpsMax       string `xml:"sps_max"`
			SpsAve       string `xml:"sps_ave"`
			InputSum     string `xml:"input_sum"`
			OutputSum    string `xml:"output_sum"`
		} `xml:"report" json:"report,omitempty"`
	}
)

//攻击日志列表
func TrafficLogList(req *TrafficLogReq, loginReq *request.LoginReq, retry bool) (res *LogsReportFlow, err error) {
	// Create a Resty Client
	client := request.GetHttpClient(loginReq)
	url := request.CheckHttpUrl("http://"+loginReq.Addr+_const.DDOS_LOGS_REPORT_FLOW_URL, loginReq)
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "", //查询
		"param_level":       fmt.Sprintf("%v", req.Level),
		"param_address":     req.Addr, //单个IP查询
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_LOGS_REPORT_FLOW_URL)

	fmt.Println(string(resp.Body()), err)

	err = xml.Unmarshal(resp.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	//if len(apiRes.Host) == 0 { //可能登陆信息过期
	//	failure := &request.Failure{}
	//	xml.Unmarshal(resp.Body(), &failure)
	//	if retry && failure.Info == _const.FAILURE_INFO {
	//		return HostList(req, loginReq, false)
	//	}
	//}
	return res, err
}
