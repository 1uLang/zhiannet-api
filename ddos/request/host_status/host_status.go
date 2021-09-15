package host_status

import (
	"encoding/xml"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"net/http"
	"strings"
	"sync"
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

type (
	HostReq struct { //获取主机信息请求参数
		Addr []string
	}
	//主机状态 返回数据
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

	//主机列表 返回数据
	StatusHost struct {
		HostAddress            string `xml:"host_address"`
		SortType               string `xml:"sort_type"`
		View                   string `xml:"view"`
		Netaddr                string `xml:"netaddr"`
		FlowThreshold          string `xml:"flow_threshold"`
		SwitchAutoProtection   string `xml:"switch_auto_protection"`
		SwitchManualProtection string `xml:"switch_manual_protection"`
		SwitchExtraProtection  string `xml:"switch_extra_protection"`
		Host                   []struct {
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
	//新增修改 返回的操作
	Success struct {
		Delay  string `xml:"delay"`
		Info   string `xml:"info"`
		URL    string `xml:"url"`
		Params string `xml:"params"`
	}

	//屏蔽列表请求参数
	ShieldListReq struct {
		Page int
		Addr string
	}
	//屏蔽列表响应参数
	StatusFblink struct {
		Filter   string `xml:"filter"`
		Page     string `xml:"page"`
		Total    string `xml:"total"`
		SortType string `xml:"sort_type"`
		Fblink   []struct {
			LocalAddress  string `xml:"local_address"`
			RemoteAddress string `xml:"remote_address"`
			ReleaseTime   string `xml:"release_time"`
			ForbidReason  string `xml:"forbid_reason"`
		} `xml:"fblink" json:"fblink,omitempty"`
	}

	//释放屏蔽列表请求参数
	ReleaseShieldReq struct {
		Addr []string
	}

	//链接列表请求参数
	LinkListReq struct {
		Page int
		Addr string
	}

	//链接列表响应参数
	StatusLink struct {
		Filter      string `xml:"filter"`
		Page        string `xml:"page"`
		Total       string `xml:"total"`
		SortType    string `xml:"sort_type"`
		LinkConnIn  string `xml:"link_conn_in"`
		LinkConnOut string `xml:"link_conn_out"`
		Link        []struct {
			LocalAddress  string `xml:"local_address"`
			RemoteAddress string `xml:"remote_address"`
			PortLinks     string `xml:"port_links"`
			TotalLinks    string `xml:"total_links"`
		} `xml:"link" json:"link,omitempty"`
	}

	//主机详情响应参数
	StatusHostResp struct {
		SettingAddress  string `xml:"setting_address"`
		Address         string `xml:"address"`
		Prefix          string `xml:"prefix"`
		Exist           string `xml:"exist"`
		RecordFlow      string `xml:"record_flow"`
		GatewayIp       string `xml:"gateway_ip"`
		GatewayMac      string `xml:"gateway_mac"`
		FlowIncomingBps string `xml:"flow_incoming_bps"`
		FlowOutgoingBps string `xml:"flow_outgoing_bps"`
		FlowIncomingPps string `xml:"flow_incoming_pps"`
		FlowOutgoingPps string `xml:"flow_outgoing_pps"`
		Ignore          string `xml:"ignore"`
		Forbid          string `xml:"forbid"`
		ForbidOverflow  string `xml:"forbid_overflow"`
		RejectForeign   string `xml:"reject_foreign"`
		ParamSet        string `xml:"param_set"`
		FilterSet       string `xml:"filter_set"`
		PortproSetTcp   string `xml:"portpro_set_tcp"`
		PortproSetUdp   string `xml:"portpro_set_udp"`
		TcpPlugin       []struct {
			Protocol    string `xml:"protocol"`
			ID          string `xml:"id"`
			Name        string `xml:"name"`
			Enabled     string `xml:"enabled"`
			SslCertList []struct {
				SslCertName string `xml:"ssl_cert_name"`
			} `xml:"ssl_cert_list" json:"ssl_cert_list,omitempty"`
		} `xml:"tcp_plugin" json:"tcp_plugin,omitempty"`
		UdpPlugin []struct {
			Protocol string `xml:"protocol"`
			ID       string `xml:"id"`
			Name     string `xml:"name"`
			Enabled  string `xml:"enabled"`
		} `xml:"udp_plugin" json:"udp_plugin,omitempty"`
		CapturedFiles string `xml:"captured_files"`
	}

	//主机设置参数设置
	HostSetReq struct {
		Addr       string
		Ignore     bool //忽略所有流量
		ProtectSet int  //防护参数集
		FilterSet  int  //过滤参数集
		SetTcp     int  //tcp端口集
		SetUdp     int  //udp端口集
	}
)

//
////主机状态API
//func HostStatus(req *HostReq, retry bool) (res *StatusFBlink, err error) {
//	// Create a Resty Client
//	resp, err := client.R().
//		SetCookie(&http.Cookie{
//			Name:  "sid",
//			Value: audit_db.Cookie,
//		}).
//		Get(_const.DDOS_HOST + _const.DDOS_HOST_STATUS_URL)
//	fmt.Println(string(resp.Body()), err)
//
//	res = &StatusFBlink{}
//	err = xml.Unmarshal(resp.Body(), res)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	if len(res.FBlink) == 0 { //可能登陆信息过期
//		failure := &audit_db.Failure{}
//		xml.Unmarshal(resp.Body(), &failure)
//		if retry && failure.Info == _const.FAILURE_INFO {
//			return HostStatus(req, false)
//		}
//	}
//
//	fmt.Println(res)
//	return res, err
//}

//主机列表
func HostList(req *HostReq, loginReq *request.LoginReq, retry bool) (res []*StatusHost, err error) {
	// Create a Resty Client
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_HOST_STATUS_URL
	wg := &sync.WaitGroup{}
	for _, v := range req.Addr {
		wg.Add(1)

		resp, err := client.R().
			SetCookie(&http.Cookie{
				Name:  "sid",
				Value: request.GetCookie(loginReq),
			}).SetFormData(map[string]string{
			"param_submit_type": "query-host", //查询
			"param_netaddr":     v,            //单个IP查询
		}).Post(url)
		//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_HOST_STATUS_URL)
		//fmt.Println("get cookie", audit_db.GetCookie(loginReq))
		//fmt.Println(string(resp.Body()), err)

		apiRes := &StatusHost{}
		err = xml.Unmarshal(resp.Body(), apiRes)
		if err != nil {
			fmt.Println(err)
			wg.Done()
			break
		}
		if apiRes == nil {
			apiRes = &StatusHost{}
		}
		apiRes.Netaddr = v
		res = append(res, apiRes)
		//if len(apiRes.Host) == 0 { //可能登陆信息过期
		//	failure := &audit_db.Failure{}
		//	xml.Unmarshal(resp.Body(), &failure)
		//	if retry && failure.Info == _const.FAILURE_INFO {
		//		return HostList(req, loginReq, false)
		//	}
		//}
		wg.Done()
	}
	wg.Wait()
	return res, err
}

//添加高防ip
//参数1 ip
//参数2 节点登陆信息
//参数3 是否重试
func AddAddr(ip string, loginReq *request.LoginReq, retry bool) (err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_HOST_STATUS_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "add-host", //添加
		"param_netaddr":     ip,         //单个IP查询
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_HOST_STATUS_URL)
	//fmt.Println("add addr res = ===",string(resp.Body()), err)

	if err != nil {
		//fmt.Println(err)
		return err
	}
	apiRes := &Success{}
	xml.Unmarshal(resp.Body(), apiRes)
	if apiRes.Info == "" {
		return fmt.Errorf("添加失败")
	}
	return err
}

//屏蔽列表
//参数1 ip
//参数2 节点登陆信息
//参数3 是否重试
func HostShieldList(req *ShieldListReq, loginReq *request.LoginReq, retry bool) (list *StatusFblink, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_FBLINK_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(
		map[string]string{
			"param_submit_type": "select",                    //添加
			"param_filter":      req.Addr,                    //单个IP查询
			"param_page":        fmt.Sprintf("%v", req.Page), //页数
		},
	).
		Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_FBLINK_URL)
	//fmt.Println("add addr res = ===", string(resp.Body()), err)

	if err != nil {
		//fmt.Println(err)
		return list, err
	}
	list = &StatusFblink{}
	err = xml.Unmarshal(resp.Body(), list)
	if err != nil {
		return list, err
	}
	return list, err
}

//重置屏蔽列表
//参数1 ip
//参数2 节点登陆信息
//参数3 是否重试
func ReleaseShield(req *ReleaseShieldReq, loginReq *request.LoginReq, retry bool) (info *Success, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_FBLINK_URL
	filter := strings.Join(req.Addr, ",")
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "reset", //重置
		"param_filter":      filter,  //ip
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_FBLINK_URL)
	//fmt.Println("add addr res = ===", string(resp.Body()), err)

	if err != nil {
		//fmt.Println(err)
		return info, err
	}
	info = &Success{}
	err = xml.Unmarshal(resp.Body(), info)
	if err != nil {
		return info, err
	}
	return info, err
}

//链接列表
//参数1 ip
//参数2 节点登陆信息
//参数3 是否重试
func LinkList(req *LinkListReq, loginReq *request.LoginReq, retry bool) (list *StatusLink, err error) {
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_LINK_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_submit_type": "select",                    //查询
		"param_page":        fmt.Sprintf("%v", req.Page), //页数
		"param_filter":      req.Addr,                    //ip
		"link_conn_in":      "on",                        //入链接
		"link_conn_out":     "on",                        //出链接
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_LINK_URL)
	//fmt.Println("add addr res = ===", string(resp.Body()), err)

	if err != nil {
		//fmt.Println(err)
		return list, err
	}
	list = &StatusLink{}
	err = xml.Unmarshal(resp.Body(), list)
	if err != nil {
		return list, err
	}
	return list, err
}

//主机详细信息
//参数1 ip + 修改参数
//参数2 节点登陆信息
//参数3 是否重试
func GetHostInfo(req *HostReq, loginReq *request.LoginReq, retry bool) (res *StatusHostResp, err error) {
	addr := ""
	if len(req.Addr) > 0 {
		addr = req.Addr[0]
	}
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_HOSTSET_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetQueryParams(map[string]string{
		"hostaddr": addr, //ip
	}).Get(url)
	//Get("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_HOSTSET_URL)
	//fmt.Println("add addr res = ===", string(resp.Body()), err)

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

//主机设置
//参数1 ip + 修改参数
//参数2 节点登陆信息
//参数3 是否重试
func SetHost(req *HostSetReq, loginReq *request.LoginReq, retry bool) (res *Success, err error) {
	//修改之前先查询之前的数据
	info := &StatusHostResp{}
	info, err = GetHostInfo(&HostReq{
		Addr: []string{req.Addr},
	}, loginReq, retry)
	if err != nil || info == nil {
		return
	}
	//修改请求
	param_ignore := "on"
	if !req.Ignore {
		param_ignore = ""
	}
	client := request.GetHttpClient(loginReq)
	url := loginReq.Addr + _const.DDOS_STATUS_HOSTSET_URL
	resp, err := client.R().
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.GetCookie(loginReq),
		}).SetFormData(map[string]string{
		"param_setting_addr":    req.Addr,                          //ip
		"param_ignore":          param_ignore,                      //忽略所有流量
		"param_param_set":       fmt.Sprintf("%v", req.ProtectSet), //防护参数集
		"param_filter_set":      fmt.Sprintf("%v", req.FilterSet),  //过滤参数集
		"param_portpro_set_tcp": fmt.Sprintf("%v", req.SetTcp),     //tcp端口集
		"param_portpro_set_udp": fmt.Sprintf("%v", req.SetUdp),     //udp端口集
		//原来的参数保持原状
		"param_prefix":                 info.Prefix,
		"param_exist":                  check(info.Exist),
		"param_record_flow":            check(info.RecordFlow),
		"param_gateway_ip":             info.GatewayIp,
		"param_gateway_mac":            info.GatewayMac,
		"param_flowlimit_incoming_bps": info.FlowIncomingBps,
		"param_flowlimit_incoming_pps": info.FlowIncomingPps,
		"param_flowlimit_outgoing_bps": info.FlowOutgoingBps,
		"param_flowlimit_outgoing_pps": info.FlowOutgoingPps,
		"param_forbid":                 check(info.Forbid),
		"param_forbid_overflow":        check(info.ForbidOverflow),
		"param_reject_foreign_access":  check(info.RejectForeign),
		"param_plugin_tcp_0":           TcpPluginCheck("0", info),
		"param_plugin_tcp_1":           TcpPluginCheck("1", info),
		"param_plugin_tcp_2":           TcpPluginCheck("2", info),
		"param_plugin_tcp_3":           TcpPluginCheck("3", info),
		"param_plugin_tcp_5":           TcpPluginCheck("5", info),
		"param_ssl_cert_list":          GetSslCert(info),
		"param_plugin_tcp_6":           TcpPluginCheck("6", info),
		"param_plugin_udp_1":           UdpPluginCheck("1", info),
		"param_plugin_udp_4":           UdpPluginCheck("4", info),
		"param_plugin_udp_5":           UdpPluginCheck("5", info),
	}).Post(url)
	//Post("https://" + loginReq.Addr + ":" + loginReq.Port + _const.DDOS_STATUS_HOSTSET_URL)
	//fmt.Println("add addr res = ===", string(resp.Body()), err)

	if err != nil {
		//fmt.Println(err)
		return res, err
	}
	res = &Success{}
	err = xml.Unmarshal(resp.Body(), res)
	if err != nil {
		return res, err
	}
	return res, err
}

//是否选中
func check(name string) (on string) {
	if name == "checked" {
		on = "on"
	}
	return
}

//tcp 插件是否选中状态
func TcpPluginCheck(id string, resp *StatusHostResp) (on string) {
	for _, v := range resp.TcpPlugin {
		if v.ID == id {
			if v.Enabled == "checked" {
				on = "on"
			}
			break
		}
	}
	return
}

//获取ssl插件证书
func GetSslCert(resp *StatusHostResp) (name string) {
	for _, v := range resp.TcpPlugin {
		if v.ID == "5" {
			if len(v.SslCertList) > 0 {
				return v.SslCertList[0].SslCertName
			}
		}
	}
	return
}

//udp插件是否选中状态
func UdpPluginCheck(id string, resp *StatusHostResp) (on string) {
	for _, v := range resp.UdpPlugin {
		if v.ID == id {
			if v.Enabled == "checked" {
				on = "on"
			}
			break
		}
	}
	return
}
