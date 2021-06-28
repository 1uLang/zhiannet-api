package logs

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"net/http"
)

type (
	LinkLogReq struct { //流量日志请求参数
		Addr  string `json:"addr"`
		Level int    `json:"level"`
	}

	LogsReportLink struct {
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
			TcpInMax     string `xml:"tcp_in_max"`
			TcpInAve     string `xml:"tcp_in_ave"`
			TcpOutMax    string `xml:"tcp_out_max"`
			TcpOutAve    string `xml:"tcp_out_ave"`
			UdpMax       string `xml:"udp_max"`
			UdpAve       string `xml:"udp_ave"`
			InputBpsMax  string `xml:"input_bps_max"`
			OutputBpsMax string `xml:"output_bps_max"`
		} `xml:"report" json:"report,omitempty"`
	}
)

//攻击日志列表
func LinkLogList(req *LinkLogReq, loginReq *request.LoginReq, retry bool) (res *LogsReportLink, err error) {
	// Create a Resty Client
	client := request.GetHttpClient(loginReq)
	url := request.CheckHttpUrl("http://"+loginReq.Addr+_const.DDOS_LOGS_REPORT_LINK_URL, loginReq)
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "", //查询
		"param_level":       fmt.Sprintf("%v", req.Level),
		"param_address":     req.Addr, //单个IP查询
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_LOGS_REPORT_LINK_URL)

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
