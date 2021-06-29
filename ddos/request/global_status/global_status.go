package global_status

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"io/ioutil"
	"net/http"
)

type (
	StatusGlobal struct { //全局统计响应参数
		WanInputBps                string `xml:"wan_input_bps"`
		WanInputPps                string `xml:"wan_input_pps"`
		WanSubmitBps               string `xml:"wan_submit_bps"`
		WanSubmitPps               string `xml:"wan_submit_pps"`
		LanInputBps                string `xml:"lan_input_bps"`
		LanInputPps                string `xml:"lan_input_pps"`
		LanSubmitBps               string `xml:"lan_submit_bps"`
		LanSubmitPps               string `xml:"lan_submit_pps"`
		AnonymousIncomingBps       string `xml:"anonymous_incoming_bps"`
		AnonymousIncomingSubmitBps string `xml:"anonymous_incoming_submit_bps"`
		AnonymousOutgoingBps       string `xml:"anonymous_outgoing_bps"`
		AnonymousOutgoingSubmitBps string `xml:"anonymous_outgoing_submit_bps"`
		TcpConnIn                  string `xml:"tcp_conn_in"`
		TcpConnOut                 string `xml:"tcp_conn_out"`
		UdpConn                    string `xml:"udp_conn"`
		HostCreated                string `xml:"host_created"`
		HostAutoProtected          string `xml:"host_auto_protected"`
		HostManualProtected        string `xml:"host_manual_protected"`
		AttackStatus               []struct {
			Flags  string `xml:"flags"`
			Status string `xml:"status"`
			Counts string `xml:"counts"`
		} `xml:"attack_status" json:"attack_status,omitempty"`
	}
	StatusHealth struct {
		ThisRackUnit string `xml:"this_rack_unit"`
		RackUnitList []struct {
			RackUnit string `xml:"rack_unit"`
		} `xml:"rack_unit_list" json:"rack_unit_list,omitempty"`
		CpuUsage    string `xml:"cpu_usage"`
		MemoryUsage string `xml:"memory_usage"`
		DeviceStat  []struct {
			Name     string `xml:"name"`
			RecvFlow string `xml:"recv_flow"`
			RecvPps  string `xml:"recv_pps"`
			RecvErr  string `xml:"recv_err"`
			SendFlow string `xml:"send_flow"`
			SendPps  string `xml:"send_pps"`
			SendErr  string `xml:"send_err"`
		} `xml:"device_stat" json:"device_stat,omitempty"`
	}
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

//获取全局统计
func GetStatusGlobal(loginReq *request.LoginReq, retry bool) (res *StatusGlobal, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_GLOBAL_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetQueryParams(map[string]string{
		//"param_submit_type": "add-host", //
	}).Get(url)
	//Get("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_GLOBAL_URL)
	fmt.Println(string(resp.Body()), err)
	if err != nil {
		//fmt.Println(err)
		return res, err
	}
	err = xml.Unmarshal(resp.Body(), &res)
	if err != nil {
		return res, err
	}
	return res, err
}

//获取负载信息-(小时|天|月)
func GetLoad(loginReq *request.LoginReq, retry bool) (res *StatusHealth, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_HEALTH_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetQueryParams(map[string]string{
		//"param_submit_type": "add-host", //
	}).Get(url)
	//Get("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_HEALTH_URL)
	fmt.Println(string(resp.Body()), err)
	if err != nil {
		//fmt.Println(err)
		return res, err
	}
	err = xml.Unmarshal(resp.Body(), &res)
	return res, err
}

func GlobalImg(loginReq *request.LoginReq, retry bool) (res []byte, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + "/cgi-bin/rateview.cgi?width=958&height=120&level=2&scale=0.25&rand=0.6897267381111953"
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetQueryParams(map[string]string{
		//"param_submit_type": "add-host", //
	}).Get(url)
	//Get("https://182.131.30.171:28443/cgi-bin/rateview.cgi?width=958&height=120&level=2&scale=0.25&rand=0.6897267381111953")
	//fmt.Println("err ===", err)
	//fmt.Println(string(resp.Body()), err)
	res, err = ioutil.ReadAll(resp.RawBody())
	if err != nil {
		fmt.Println("读取图片失败", err)
		return
	}
	return

}
