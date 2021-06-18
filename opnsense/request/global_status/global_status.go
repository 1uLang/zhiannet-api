package global_status

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/opnsense/const"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/go-resty/resty/v2"
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

//全局
func GetGlobal(apiKey *request.ApiKey) (res *GlobalStatus, err error) {
	//https://182.150.0.109:5443/widgets/api/get.php?load=system%2Cinterfaces&_=1623836610177
	//PHPSESSID=f8eea58ee17da3ce16ba39bf17312346
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("x-csrftoken", apiKey.XCsrfToken).
		SetCookie(&http.Cookie{
			Name:  "PHPSESSID",
			Value: apiKey.Cookie,
		}).
		Get("https://" + apiKey.Addr + ":" + apiKey.Port + _const.OPNSENSE_GLOBAL_STATUS_URL + fmt.Sprintf("%v", time.Now().Unix()))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode() == 200 {
		err = json.Unmarshal(resp.Body(), &res)
	}
	fmt.Println(string(resp.Body()))
	return res, err
}
