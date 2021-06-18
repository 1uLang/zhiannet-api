package filter

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"

	"github.com/go-resty/resty/v2"
)

type (
	FilterReq struct { //搜索请求参数
		Current      string //页数
		RowCount     string //每页条数
		SearchPhrase string //关键词
	}
	UuidReq struct { //获取详情,启动停止，删除 请求参数
		Uuid string `json:"uuid"`
	}
	FilterListResp struct { //filter 列表响应参数
		Current  int `json:"current"`
		RowCount int `json:"rowCount"`
		Rows     []struct {
			Description string `json:"description"`
			Enabled     string `json:"enabled"`
			Sequence    string `json:"sequence"`
			UUID        string `json:"uuid"`
		} `json:"rows"`
		Total int `json:"total"`
	}
	EditResp struct { //启停｜删除｜修改 响应参数
		Result  string `json:"result"`
		Changed bool   `json:"changed"`
	}
	FilterInfoResp struct { //规则详情响应参数
		Rule struct {
			Action struct {
				Block struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"block"`
				Pass struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"pass"`
				Reject struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"reject"`
			} `json:"action"`
			Description     string `json:"description"`
			DestinationNet  string `json:"destination_net"`
			DestinationNot  string `json:"destination_not"`
			DestinationPort string `json:"destination_port"`
			Direction       struct {
				In struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"in"`
				Out struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"out"`
			} `json:"direction"`
			Enabled string `json:"enabled"`
			Gateway struct {
				Null struct {
					Selected bool   `json:"selected"`
					Value    string `json:"value"`
				} `json:""`
				Null4 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"Null4"`
				Null6 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"Null6"`
				WANGW struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"WAN_GW"`
			} `json:"gateway"`
			Interface struct {
				Lan struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"lan"`
				Lo0 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"lo0"`
				Wan struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"wan"`
			} `json:"interface"`
			Ipprotocol struct {
				Inet struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"inet"`
				Inet6 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"inet6"`
			} `json:"ipprotocol"`
			Log      string `json:"log"`
			Protocol struct {
				AH struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"AH"`
				AN struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"A/N"`
				ARGUS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ARGUS"`
				ARIS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ARIS"`
				AX25 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"AX.25"`
				Any struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"any"`
				BBNRCC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"BBN-RCC"`
				BNA struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"BNA"`
				BRSATMON struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"BR-SAT-MON"`
				CARP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CARP"`
				CBT struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CBT"`
				CFTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CFTP"`
				CHAOS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CHAOS"`
				COMPAQPEER struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"COMPAQ-PEER"`
				CPHB struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CPHB"`
				CPNX struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CPNX"`
				CRTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CRTP"`
				CRUDP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"CRUDP"`
				DCCP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DCCP"`
				DCN struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DCN"`
				DDP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DDP"`
				DDX struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DDX"`
				DGP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DGP"`
				DIVERT struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DIVERT"`
				DSR struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"DSR"`
				EGP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"EGP"`
				EIGRP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"EIGRP"`
				EMCON struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"EMCON"`
				ENCAP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ENCAP"`
				ESP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ESP"`
				ETHERIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ETHERIP"`
				FC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"FC"`
				GGP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"GGP"`
				GMTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"GMTP"`
				GRE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"GRE"`
				HIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"HIP"`
				HMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"HMP"`
				IATP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IATP"`
				ICMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ICMP"`
				IDPR struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IDPR"`
				IDPRCMTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IDPR-CMTP"`
				IDRP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IDRP"`
				IFMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IFMP"`
				IGMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IGMP"`
				IGP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IGP"`
				IL struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IL"`
				INLSP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"I-NLSP"`
				IPCOMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPCOMP"`
				IPCV struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPCV"`
				IPENCAP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPENCAP"`
				IPIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPIP"`
				IPPC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPPC"`
				IPV6 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPV6"`
				IPV6ICMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPV6-ICMP"`
				IPXINIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IPX-IN-IP"`
				IRTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"IRTP"`
				ISIS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ISIS"`
				ISOIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ISO-IP"`
				ISOTP4 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ISO-TP4"`
				KRYPTOLAN struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"KRYPTOLAN"`
				L2TP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"L2TP"`
				LARP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"LARP"`
				LEAF1 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"LEAF-1"`
				LEAF2 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"LEAF-2"`
				MANET struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MANET"`
				MERITINP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MERIT-INP"`
				MFENSP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MFE-NSP"`
				MICP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MICP"`
				MOBILE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MOBILE"`
				MPLSINIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MPLS-IN-IP"`
				MTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MTP"`
				MUX struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"MUX"`
				NARP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"NARP"`
				NETBLT struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"NETBLT"`
				NSFNETIGP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"NSFNET-IGP"`
				NVP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"NVP"`
				OSPF struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"OSPF"`
				PC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"3PC"`
				PFSYNC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PFSYNC"`
				PGM struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PGM"`
				PIM struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PIM"`
				PIPE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PIPE"`
				PNNI struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PNNI"`
				PRM struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PRM"`
				PTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PTP"`
				PUP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PUP"`
				PVP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"PVP"`
				QNX struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"QNX"`
				RDP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"RDP"`
				ROHC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ROHC"`
				RSVP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"RSVP"`
				RSVPE2EIGNORE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"RSVP-E2E-IGNORE"`
				RVD struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"RVD"`
				SATEXPAK struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SAT-EXPAK"`
				SATMON struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SAT-MON"`
				SCCSP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SCC-SP"`
				SCPS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SCPS"`
				SCTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SCTP"`
				SDRP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SDRP"`
				SECUREVMTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SECURE-VMTP"`
				SHIM6 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SHIM6"`
				SKIP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SKIP"`
				SM struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SM"`
				SMP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SMP"`
				SNP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SNP"`
				SPRITERPC struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SPRITE-RPC"`
				SPS struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SPS"`
				SRP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SRP"`
				ST2 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"ST2"`
				STP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"STP"`
				SUNND struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SUN-ND"`
				SWIPE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"SWIPE"`
				TCF struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TCF"`
				TCP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TCP"`
				TLSP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TLSP"`
				TP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TP++"`
				TRUNK1 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TRUNK-1"`
				TRUNK2 struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TRUNK-2"`
				TTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"TTP"`
				UDP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"UDP"`
				UDPLITE struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"UDPLITE"`
				UTI struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"UTI"`
				VINES struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"VINES"`
				VISA struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"VISA"`
				VMTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"VMTP"`
				WBEXPAK struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"WB-EXPAK"`
				WBMON struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"WB-MON"`
				WESP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"WESP"`
				WSN struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"WSN"`
				XNET struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"XNET"`
				XNSIDP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"XNS-IDP"`
				XTP struct {
					Selected int    `json:"selected"`
					Value    string `json:"value"`
				} `json:"XTP"`
			} `json:"protocol"`
			Quick      string `json:"quick"`
			Sequence   string `json:"sequence"`
			SourceNet  string `json:"source_net"`
			SourceNot  string `json:"source_not"`
			SourcePort string `json:"source_port"`
		} `json:"rule"`
	}

	//保存规则请求参数
	SaveReq struct {
		Uuid string `json:"uuid"`
		Add  bool   `json:"add"`
		Rule Rule   `json:"rule"`
	}
	Rule map[string]string
)

var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetTimeout(time.Second * 2)

//获取Filter规则列表
func GetFilterList(req *FilterReq, apiKey *request.ApiKey) (list *FilterListResp, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		SetFormData(map[string]string{
			"current":      req.Current,      //页数
			"rowCount":     req.RowCount,     //每页条数
			"searchPhrase": req.SearchPhrase, //关键词
		}).
		Post(fmt.Sprintf("https://%v:%v%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_FILTER_SEARCH_URL))
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &list)
	if err != nil {
		return list, err
	}
	return list, err
}

//启用｜停用 规则
func EnableFilter(req *UuidReq, apiKey *request.ApiKey) (res bool, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		Post(fmt.Sprintf("https://%v:%v%v/%v", apiKey.Addr, apiKey.Port, _const.OPNSENSE_FILTER_ENABLE_URL, req.Uuid))
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	return editRes.Changed, err
}

//删除 规则
func DelFilter(req *UuidReq, apiKey *request.ApiKey) (res bool, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		Post(fmt.Sprintf("https://%v:%v%v/%v", apiKey.Addr, apiKey.Port,
			_const.OPNSENSE_FILTER_DEL_URL, req.Uuid))
	//fmt.Println(string(resp.Body()), err)
	editRes := EditResp{}
	err = json.Unmarshal(resp.Body(), &editRes)
	if err != nil {
		return res, err
	}
	return editRes.Result == "deleted", err
}

//获取规则详情数据
func GetFilterInfo(req *UuidReq, apiKey *request.ApiKey) (res *FilterInfoResp, err error) {
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		Post(fmt.Sprintf("https://%v:%v%v/%v", apiKey.Addr, apiKey.Port,
			_const.OPNSENSE_FILTER_INFO_URL, req.Uuid))
	//fmt.Println(string(resp.Body()), err)
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return res, err
	}
	return
}

//保存规则
func SaveFilter(req *SaveReq, apiKey *request.ApiKey) (res bool, err error) {
	url := _const.OPNSENSE_FILTER_SET_URL
	if req.Add { //添加
		url = _const.OPNSENSE_FILTER_ADD_URL
	}
	resp, err := client.R().
		SetBasicAuth(apiKey.Username, apiKey.Password).
		SetBody(req.Rule).
		Post(fmt.Sprintf("https://%v:%v%v/%v", apiKey.Addr, apiKey.Port,
			url, req.Uuid))
	//fmt.Println(string(resp.Body()), err)
	var saveRes EditResp
	err = json.Unmarshal(resp.Body(), &saveRes)
	if err != nil {
		return res, err
	}
	return saveRes.Result == "saved", err
}
