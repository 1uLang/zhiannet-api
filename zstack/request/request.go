package request

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/utils"
)

//var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

func Request(loginReq *LoginReq, retry bool) (res []byte, err error) {
	switch loginReq.ReqType {
	case "post":
		return Post(loginReq, true)
	case "put":
		return Put(loginReq, true)
	case "delete":
		return Delete(loginReq, true)
	case "get":
		return Get(loginReq, true)
	default:
		return Get(loginReq, true)
	}

}

//请求
func Get(loginReq *LoginReq, retry bool) (res []byte, err error) {
	fmt.Println("Authorization", "OAuth "+loginReq.UUID)
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)
	resp, err := client.SetDebug(false).R(). //SetAuthToken(loginReq.UUID).
							SetHeader("Authorization", "OAuth "+loginReq.UUID).
							SetBody(loginReq.QueryParams).
							Get(url)

	//fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}

//Post
func Post(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)

	resp, err := client.R().SetHeader("Authorization", "OAuth "+loginReq.UUID).
		//SetQueryParams(loginReq.QueryParams).
		SetBody(loginReq.QueryParams).
		Post(url)

	//fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}

func Put(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)

	resp, err := client.R().SetHeader("Authorization", "OAuth "+loginReq.UUID).
		//SetQueryParams(loginReq.QueryParams).
		SetBody(loginReq.QueryParams).
		Put(url)

	//fmt.Println("body===", string(resp.Body()), err)
	res = resp.Body()
	return res, err
}

func Delete(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)

	resp, err := client.R().SetHeader("Authorization", "OAuth "+loginReq.UUID).
		//SetQueryParams(loginReq.QueryParams).
		SetBody(loginReq.QueryParams).
		Delete(url)

	//fmt.Println(resp, err)
	//fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}
