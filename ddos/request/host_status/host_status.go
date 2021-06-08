package host_status

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type (
	HostReq struct {
		Addr []string
	}
	//主机状态
	StatusFBlink struct {
		Filter    string   `xml:"filter"`
		Page      string   `xml:"page"`
		Total     string   `xml:"total"`
		PageTitle string   `xml:"page_title"`
		sortType  string   `xml:"sort_type"`
		FBlink    []Fblink `xml:"fblink"`
	}
	Fblink struct {
		LocalAddress  string `xml:"local_address"`
		RemoteAddress string `xml:"remote_address"`
		ReleaseTime   string `xml:"release_time"`
		ForbidReason  string `xml:"forbid_reason"`
	}

	//主机列表
	StatusHost struct {
		//XMLName                xml.Name `xml:"status_host" json:"status_host,omitempty"`
		//Text                   string   `xml:",chardata" json:"text,omitempty"`
		HostAddress            string `xml:"host_address"`
		SortType               string `xml:"sort_type"`
		View                   string `xml:"view"`
		Netaddr                string `xml:"netaddr"`
		FlowThreshold          string `xml:"flow_threshold"`
		SwitchAutoProtection   string `xml:"switch_auto_protection"`
		SwitchManualProtection string `xml:"switch_manual_protection"`
		SwitchExtraProtection  string `xml:"switch_extra_protection"`
		Host                   []struct {
			//Text              string `xml:",chardata" json:"text,omitempty"`
			Type              string `xml:"type"`
			Address           string `xml:"address"`
			InputBps          string `xml:"input_bps"`
			InputPps          string `xml:"input_pps"`
			InputSubmitBps    string `xml:"input_submit_bps"`
			InputSubmitPps    string `xml:"input_submit_pps"`
			OutputBps         string `xml:"output_bps"`
			OutputPps         string `xml:"output_pps"`
			OutputSubmitBps   string `xml:"output_submit_bps"`
			OutputSubmitPps   string `xml:"output_submit_pps"`
			BaselineReference string `xml:"baseline_reference"`
			SynRate           string `xml:"syn_rate"`
			AckRate           string `xml:"ack_rate"`
			UdpRate           string `xml:"udp_rate"`
			IcmpRate          string `xml:"icmp_rate"`
			FragRate          string `xml:"frag_rate"`
			NonipRate         string `xml:"nonip_rate"`
			NewTcpRate        string `xml:"new_tcp_rate"`
			NewUdpRate        string `xml:"new_udp_rate"`
			TcpConnIn         string `xml:"tcp_conn_in"`
			TcpConnOut        string `xml:"tcp_conn_out"`
			UdpConn           string `xml:"udp_conn"`
			IcmpConn          string `xml:"icmp_conn"`
			Status            string `xml:"status"`
			ProtectCredits    string `xml:"protect_credits"`
		} `xml:"host" json:"host,omitempty"`
	}
)

//主机状态API
func HostStatus(req *HostReq, retry bool) (res *StatusFBlink, err error) {
	// Create a Resty Client
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.Cookie,
		}).
		Get(_const.DDOS_HOST + _const.DDOS_HOST_STATUS_URL)
	fmt.Println(string(resp.Body()), err)

	res = &StatusFBlink{}
	err = xml.Unmarshal(resp.Body(), res)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(res.FBlink) == 0 { //可能登陆信息过期
		failure := &request.Failure{}
		xml.Unmarshal(resp.Body(), &failure)
		if retry && failure.Info == _const.FAILURE_INFO {
			return HostStatus(req, false)
		}
	}

	fmt.Println(res)
	return res, err
}

//主机列表
func HostList(req *HostReq, loginReq *request.LoginReq, retry bool) (res []*StatusHost, err error) {
	// Create a Resty Client
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	for _, v := range req.Addr {
		resp, err := client.R().
			SetCookie(&http.Cookie{
				Name:  "sid",
				Value: request.GetCookie(loginReq),
			}).SetFormData(map[string]string{
			"param_submit_type": "query-host", //查询
			"param_netaddr":     v,            //单个IP查询
		}).
			Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_HOST_STATUS_URL)

		fmt.Println(string(resp.Body()), err)

		apiRes := &StatusHost{}
		err = xml.Unmarshal(resp.Body(), apiRes)
		if err != nil {
			fmt.Println(err)
			break
		}
		res = append(res, apiRes)
		//if len(apiRes.Host) == 0 { //可能登陆信息过期
		//	failure := &request.Failure{}
		//	xml.Unmarshal(resp.Body(), &failure)
		//	if retry && failure.Info == _const.FAILURE_INFO {
		//		return HostList(req, loginReq, false)
		//	}
		//}
	}
	return res, err
}
