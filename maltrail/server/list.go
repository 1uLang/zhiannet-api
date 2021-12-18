package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	_const "github.com/1uLang/zhiannet-api/maltrail/const"
	"github.com/1uLang/zhiannet-api/maltrail/request"
	"regexp"
	"strings"
	"time"
)

type (
	ListReq struct {
		Date   string `json:"date"`
		NodeId uint64 `json:"node_id"`
		Type   string `json:"type"` //类型 url=网络防病毒
	}

	ListResp struct {
		Time      string `json:"time"`
		Events    int    `json:"events"`
		Sensor    string `json:"sensor"`
		SrcIp     string `json:"src_ip"`
		SrcPort   string `json:"src_port"`
		DstIp     string `json:"dst_ip"`
		DstPort   string `json:"dst_port"`
		Proto     string `json:"proto"`
		Type      string `json:"type"`
		Trail     string `json:"trail"`
		Info      string `json:"info"`
		Reference string `json:"reference"`
	}
)

//apt 节点列表
func GetMaltrailNodeList() (list []*subassemblynode.Subassemblynode, total int64, err error) {
	list, total, err = subassemblynode.GetList(&subassemblynode.NodeReq{Type: 11, State: "1"})
	return
}
func GetList(req *ListReq) (list []*ListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(req.NodeId)
	if err != nil {
		return
	}
	if req.Date == "" {
		Now := time.Now()
		req.Date = time.Date(Now.Year(), Now.Month(), Now.Day(), 0, 0, 0, 0, time.Local).Format("2006-01-02")
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.MALTRAIL_LIST)
	logReq.QueryParams = map[string]string{
		"date": req.Date,
	}
	//获取用户关联的审计ID
	resp, err := request.Get(
		logReq, true,
	)
	if err != nil {
		return
	}
	return MustList(resp, req.Type)

}

func MustList(str []byte, t string) (newResp []*ListResp, err error) {
	strs := string(str)
	resp := make([]*ListResp, 0)
	newResp = make([]*ListResp, 0)

	strSli := strings.Split(strs, "\n")
	if len(strSli) > 1 {
		//re,_ := regexp.Compile(`\"\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2}\.\d{2,}\" \S{1,} \S{1,} \d{1,} \S{1,} \d{1,} \S{1,} \S{1,} \S{1,} \".*?\" \".*?\"`)
		re, _ := regexp.Compile(`\".*?\"`)
		strMap := map[string]int{}
		for _, v := range strSli {
			ips := &ListResp{}
			s := strings.Split(v, " ")
			if len(s) > 10 {
				ips = &ListResp{
					Sensor:  s[2],
					SrcIp:   s[3],
					SrcPort: s[4],
					DstIp:   s[5],
					DstPort: s[6],
					Proto:   s[7],
					Type:    s[8],
					Trail:   s[9],
				}
			}
			ll := re.FindAllString(v, -1)
			//fmt.Println(len(ll),ll)
			if len(ll) >= 2 {
				firshTime, _ := time.ParseInLocation("2006-01-02 15:04:05.000000", strings.Trim(ll[0], "\""), time.Local)
				ips.Time = firshTime.Format("2006-01-02 15:04:05")
				ips.Info = strings.TrimRight(strings.TrimLeft(ll[1], "\""), "\"")
				ips.Reference = s[len(s)-1]
				if len(ll) >= 3 {
					//ips.Trail = ll[1]
					ips.Reference = ll[2]
				}
				if len(ll) >= 4 {
					ips.Trail = ll[1]
					ips.Info = strings.TrimRight(strings.TrimLeft(ll[2], "\""), "\"")
					ips.Reference = ll[3]
				}
			}
			if ips.Trail != "" {
				//event 计数
				if num, ok := strMap[ips.SrcIp]; ok {
					strMap[ips.SrcIp] = num + 1
				} else {
					strMap[ips.SrcIp] = 1
				}
				resp = append(resp, ips)
			}
		}
		reg := regexp.MustCompile(`(\(.*?\))$`)
		for _, v := range resp {
			//fmt.Println("--------",t,strings.TrimSpace(v.Type))
			if t == "url" { //网络防病毒
				if strings.TrimSpace(v.Type) == "URL" {
					if num, ok := strMap[v.SrcIp]; ok {
						value := &ListResp{}
						*value = *v
						value.Events = num
						value.Trail = reg.ReplaceAllString(value.Trail, ``) //182.150.0.117(/board.cgi) => 182.150.0.117 正则替换，去掉后面括号和括号内的内容
						newResp = append(newResp, value)
						delete(strMap, v.SrcIp)
					}
				}

			} else {
				if num, ok := strMap[v.SrcIp]; ok {
					value := &ListResp{}
					*value = *v
					value.Events = num
					newResp = append(newResp, value)
					delete(strMap, v.SrcIp)
				}
			}

		}
	}
	return newResp, err
}
