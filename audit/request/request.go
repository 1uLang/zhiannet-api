package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/utils"
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func Request(loginReq *LoginReq, retry bool) (res []byte, err error) {
	if loginReq.ReqType == "post" {
		return Post(loginReq, true)
	}
	return Get(loginReq, true)
}

//请求
func Get(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)

	resp, err := client.R().SetAuthToken(loginReq.Token).
		//SetQueryParams(loginReq.QueryParams).
		SetBody(loginReq.QueryParams).
		Get(url)

	fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}

//Post
func Post(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)

	resp, err := client.R().SetAuthToken(loginReq.Token).
		//SetQueryParams(loginReq.QueryParams).
		SetBody(loginReq.QueryParams).
		Post(url)

	fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}
