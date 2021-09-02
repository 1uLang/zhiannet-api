package request

import (
	"github.com/1uLang/zhiannet-api/utils"
)

type ()

func Get(loginReq *LoginReq, retry bool) (res []byte, err error) {
	client := GetHttpClient(loginReq)
	url := utils.CheckHttpUrl(loginReq.Addr, loginReq.IsSsl)
	resp, err := client.R().SetHeader("Cookie", loginReq.Cookie).
		SetQueryParams(loginReq.QueryParams).
		Get(url)

	//fmt.Println(string(resp.Body()), err)
	res = resp.Body()
	return res, err
}
