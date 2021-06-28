package logs

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"net/http"
	"time"
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

type (
	AttackLogReq struct { //黑白名单请求参数
		Addr       string    `json:"addr"`
		StartTime  time.Time `json:"start_time"`
		EndTime    time.Time `json:"end_time"`
		AttackType string    `json:"attack_type"`
		Status     int       `json:"status"`
	}

	LogsReportAttack struct { //攻击日志响应参数
		Address   string `xml:"address"`
		StartDate string `xml:"start_date"`
		EndDate   string `xml:"end_date"`
		CurFlags  string `xml:"cur_flags"`
		CurStatus string `xml:"cur_status"`
		Page      string `xml:"page"`
		Total     string `xml:"total"`
		Report    []struct {
			Sid         string `xml:"sid"`
			DstAddress  string `xml:"dst_address"`
			DstPort     string `xml:"dst_port"`
			Begin       string `xml:"begin"`
			End         string `xml:"end"`
			Last        string `xml:"last"`
			Flags       string `xml:"flags"`
			BpsIn       string `xml:"bps_in"`
			FromAddress string `xml:"from_address"`
			Status      string `xml:"status"`
			Archive     string `xml:"archive"`
			Highproto   string `xml:"highproto"`
		} `xml:"report" json:"report,omitempty"`
	}

	//新增修改 返回的操作
	Success struct {
		Delay  string `xml:"delay"`
		Info   string `xml:"info"`
		URL    string `xml:"url"`
		Params string `xml:"params"`
	}
)

//攻击日志列表
func AttackLogList(req *AttackLogReq, loginReq *request.LoginReq, retry bool) (res *LogsReportAttack, err error) {
	// Create a Resty Client
	client := request.GetHttpClient(loginReq)
	url := request.CheckHttpUrl("http://"+loginReq.Addr+_const.DDOS_LOGS_REPORT_ATTACK_URL, loginReq)
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "submit",                                    //查询
		"param_start_date":  req.StartTime.Format("2006-01-02 15:04:05"), //开始时间
		"param_end_date":    req.EndTime.Format("2006-01-02 15:04:05"),   //结束时间
		"param_address":     req.Addr,                                    //单个IP查询
		"param_flags":       req.AttackType,                              //类型
		"param_status":      fmt.Sprintf("%v", req.Status),               //状态
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_LOGS_REPORT_ATTACK_URL)

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
