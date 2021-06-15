package global_status

import (
	"crypto/tls"
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/go-resty/resty/v2"
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
					Load          []string `json:"load"`
					Maxfreq       string   `json:"max.freq"`
					Model         string   `json:"model"`
					Nice          string   `json:"nice"`
					Sys           string   `json:"sys"`
					Used          string   `json:"used"`
					User          string   `json:"user"`
				} `json:"cpu"`
				Config struct {
					LastChange     string `json:"last_change"`
					LastChangeFrmt string `json:"last_change_frmt"`
				} `json:"config"`
				DateFrmt string `json:"date_frmt"`
				DateTime string `json:"date_time"`
				Disk     struct {
					Devices []struct {
						Available  string `json:"available"`
						Capacity   string `json:"capacity"`
						Device     string `json:"device"`
						Mountpoint string `json:"mountpoint"`
						Size       string `json:"size"`
						Type       string `json:"type"`
						Used       string `json:"used"`
					} `json:"devices"`
					Swap []struct {
						Device string `json:"device"`
						Total  string `json:"total"`
						Used   string `json:"used"`
					} `json:"swap"`
				} `json:"disk"`
				Kernel struct {
					Mbuf struct {
						Max   string `json:"max"`
						Total string `json:"total"`
					} `json:"mbuf"`
					Memory struct {
						Total string `json:"total"`
						Used  string `json:"used"`
					} `json:"memory"`
					Pf struct {
						Maxstates string `json:"maxstates"`
						States    string `json:"states"`
					} `json:"pf"`
				} `json:"kernel"`
				Uptime   string   `json:"uptime"`
				Versions []string `json:"versions"`
			} `json:"system"`
		} `json:"data"`
		Plugins []string `json:"plugins"`
		System  string   `json:"system"`
	}
)

var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

//获取全局统计
func GetStatusGlobal(apiKey *request.ApiKey, retry bool) (err error) {
	resp, err := client.R().
		SetBasicAuth("XUWAuB5AYrfT3BmO6KYWYU7nP6fdMdA4ljLEi5deK9+Rfxp4oJI6ZBfoOTR553HGnUl3Pq45iA4Usv3b", "b+Q/b1HRDASLYdM6DDVmmvdTMm1MIUgLMPhkRINxkSNSWUanBejXRqhE71aQrHYCBzeOoUN0RNbYzhlE").
		SetQueryParams(map[string]string{
			//"param_submit_type": "add-host", //
		}).
		Get("https://182.150.0.109:5443/widgets/api/get.php?load=system%2Cinterfaces")
	fmt.Println(string(resp.Body()), err)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return err
}

//nat 规则列表
func GetNATList(apiKey *request.ApiKey, retry bool) (err error) {
	resp, err := client.R().
		SetBasicAuth("XUWAuB5AYrfT3BmO6KYWYU7nP6fdMdA4ljLEi5deK9+Rfxp4oJI6ZBfoOTR553HGnUl3Pq45iA4Usv3b", "b+Q/b1HRDASLYdM6DDVmmvdTMm1MIUgLMPhkRINxkSNSWUanBejXRqhE71aQrHYCBzeOoUN0RNbYzhlE").
		SetQueryParams(map[string]string{
			//"param_submit_type": "add-host", //
		}).
		Get("https://182.150.0.109:5443/api/firewall/filter/searchRule?current=1&rowCount=7&searchPhrase=")
	fmt.Println(string(resp.Body()), err)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return err
}
