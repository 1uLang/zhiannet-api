package global_status

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"net/http"
	"time"
)

type (
	GlobalStatus struct {
		Data struct {
			Interfaces []struct {
				Collisions   string `json:"collisions"`
				Descr        string `json:"descr"`
				Inbytes      string `json:"inbytes"`
				InbytesFrmt  string `json:"inbytes_frmt"`
				Inerrs       string `json:"inerrs"`
				Inpkts       string `json:"inpkts"`
				Ipaddr       string `json:"ipaddr"`
				Media        string `json:"media"`
				Name         string `json:"name"`
				Outbytes     string `json:"outbytes"`
				OutbytesFrmt string `json:"outbytes_frmt"`
				Outerrs      string `json:"outerrs"`
				Outpkts      string `json:"outpkts"`
				Status       string `json:"status"`
			} `json:"interfaces"`
			System struct {
				CPU struct {
					Cpus          string   `json:"cpus"`
					Curfreq       string   `json:"cur.freq"`
					FreqTranslate string   `json:"freq_translate"`
					Idle          string   `json:"idle"`
					Intr          string   `json:"intr"`
					Load          []string `json:"load"` //cpu 负载
					Maxfreq       string   `json:"max.freq"`
					Model         string   `json:"model"` //cpu type
					Nice          string   `json:"nice"`
					Sys           string   `json:"sys"`
					Used          string   `json:"used"` //cpu使用率
					User          string   `json:"user"`
				} `json:"cpu"`
				Config struct {
					LastChange     string `json:"last_change"`
					LastChangeFrmt string `json:"last_change_frmt"` //最近一次配置时间
				} `json:"config"`
				DateFrmt string `json:"date_frmt"` //当前时间
				DateTime string `json:"date_time"`
				Disk     struct {
					Devices []struct { //磁盘信息
						Available  string `json:"available"`
						Capacity   string `json:"capacity"`
						Device     string `json:"device"`
						Mountpoint string `json:"mountpoint"`
						Size       string `json:"size"`
						Type       string `json:"type"`
						Used       string `json:"used"`
					} `json:"devices"`
					Swap []struct { //swap 信息
						Device string `json:"device"`
						Total  string `json:"total"`
						Used   string `json:"used"`
					} `json:"swap"`
				} `json:"disk"`
				Kernel struct {
					Mbuf struct { //MBUF 信息
						Max   string `json:"max"`
						Total string `json:"total"`
					} `json:"mbuf"`
					Memory struct { //内存信息
						Total string `json:"total"`
						Used  string `json:"used"`
					} `json:"memory"`
					Pf struct { //状态表大小
						Maxstates string `json:"maxstates"`
						States    string `json:"states"`
					} `json:"pf"`
				} `json:"kernel"`
				Uptime   string   `json:"uptime"`   //运行时间 秒
				Versions []string `json:"versions"` //版本信息
			} `json:"system"`
		} `json:"data"`
		Plugins []string `json:"plugins"`
		System  string   `json:"system"`
	}
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

//全局
func GetGlobal(apiKey *request.ApiKey) (res *GlobalStatus, err error) {
	url := apiKey.Addr + _const.OPNSENSE_GLOBAL_STATUS_URL + fmt.Sprintf("%v", time.Now().Unix())
	client := request.GetHttpClient(apiKey)
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).Get(url)
	//Get("https://" + apiKey.Addr + ":" + apiKey.Port + _const.OPNSENSE_GLOBAL_STATUS_URL + fmt.Sprintf("%v", time.Now().Unix()))
	if err != nil {
		//fmt.Println(err)
		return res, err
	}
	if resp.StatusCode() == 200 {
		err = json.Unmarshal(resp.Body(), &res)
	}
	//fmt.Println(string(resp.Body()))
	return res, err
}
