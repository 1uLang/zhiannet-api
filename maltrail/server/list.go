package server

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/maltrail/const"
	"github.com/1uLang/zhiannet-api/maltrail/request"
	"regexp"
	"strings"
	"time"
)

type (
	ListReq struct {
		Date string `json:"date"`
	}

	ListResp struct {
		Time      string `json:"time"`
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

func GetList(req *ListReq) (list []*ListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo()
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
	return MustList(resp)

}

func MustList(str []byte) (resp []*ListResp, err error) {
	strs := string(str)
	strSli := strings.Split(strs, "\n")
	if len(strSli) > 1 {
		//re,_ := regexp.Compile(`\"\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2}\.\d{2,}\" \S{1,} \S{1,} \d{1,} \S{1,} \d{1,} \S{1,} \S{1,} \S{1,} \".*?\" \".*?\"`)
		re, _ := regexp.Compile(`\".*?\"`)
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
				ips.Info = ll[1]
				ips.Reference = s[len(s)-1]
				if len(ll) >= 3 {
					//ips.Trail = ll[1]
					ips.Reference = ll[2]
				}
				if len(ll) >= 4 {
					ips.Trail = ll[1]
					ips.Info = ll[2]
					ips.Reference = ll[3]
				}
			}
			if ips.Trail != "" {
				resp = append(resp, ips)
			}
		}
	}
	return
}
