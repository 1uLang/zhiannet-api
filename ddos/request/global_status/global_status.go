package global_status

import (
	"crypto/tls"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/ddos/const"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/go-resty/resty/v2"
	"net/http"
)

//获取负载信息-(小时|天|月)
func GetStatusGlobal() {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetHeader("Content-Type", "text/xml").
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.Cookie,
		}).
		Get(_const.DDOS_HOST + _const.DDOS_STATUS_GLOBAL_URL)
	fmt.Println(string(resp.Body()), err)
	return
}

//获取负载信息-(小时|天|月)
func GetLoad() {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().
		SetHeader("Content-Type", "text/xml").
		SetCookie(&http.Cookie{
			Name:  "sid",
			Value: request.Cookie,
			//Path:     "/",
			//Domain:   "http://182.131.30.171",
			//MaxAge:   36000,
			//HttpOnly: true,
			//Secure:   false,
		}).
		Get(_const.DDOS_HOST + _const.DDOS_STATUS_HEALTH_URL)
	fmt.Println(string(resp.Body()), err)
	return
}
