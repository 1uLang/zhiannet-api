package global_status

import (
	"crypto/tls"
	"fmt"
	"github.com/1uLang/zhiannet-api/opnsense/request"
	"github.com/go-resty/resty/v2"
)

type ()

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
