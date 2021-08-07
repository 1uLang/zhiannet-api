package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"time"
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

//登陆获取cookie
func Login(req *ApiKey) (CookieMap map[string]string, err error) {
	url := req.Addr
	Client = GetHttpClient(req)
	CookieMap = make(map[string]string)
	// https://182.150.0.109:5443/
	//访问登陆页 获取登陆需要的唯一凭证 key-value
	index, err := Client.R().Get(url)
	if err != nil {
		return CookieMap, err
	}
	//fmt.Println(index.StatusCode())
	//fmt.Println(string(index.Body()))

	//解析html
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(index.Body()))
	if err != nil {
		return CookieMap, err
	}
	key, value := "", ""
	//doc.Find("form input[type='hidden']").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the band and title
	//	name, _ = s.Attr("name")
	//	value, _ = s.Attr("value")
	//	fmt.Printf("Review %d: %s - %s\n", i, name, value)
	//})
	//通过标签匹配获取key - value
	input := doc.Find("form input[type='hidden']").First()
	key, _ = input.Attr("name")
	value, _ = input.Attr("value")
	//登陆 返回cookies
	resp, err := Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"login":       "1",
			"usernamefld": req.Username,
			"passwordfld": req.Password,
			key:           value,
		}).
		Post(url)
	//fmt.Println("code========", resp.StatusCode())
	//fmt.Println(string(resp.Body()))
	//fmt.Println("key=",name,"value=",value)
	if err != nil {
		return CookieMap, err
	}
	if resp.StatusCode() == 200 {
		//获取cookie
		Cookies := resp.Cookies()
		if len(Cookies) > 0 {
			CookieMap["cookie"] = Cookies[0].Value
			CookieMap["x-csrftoken"] = value //接口调用凭证
		}

		fmt.Println("cookies", Cookies)
		//fmt.Println("login in Cookie=", Cookie)
	}
	return CookieMap, err
}

//获取cookie和接口凭证 x-csrftoken
func GetCookie(req *ApiKey) (cookie, x_csrftoken string, err error) {

	key := fmt.Sprintf("opnsense-cookie-%v:%v", req.Addr, req.Port)
	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return Login(req)
	}, 600, true)
	if err != nil {
		return cookie, x_csrftoken, err
	}
	var resByte []byte
	resByte, err = json.Marshal(resp)
	cookie = gjson.ParseBytes(resByte).Get("cookie").String()
	x_csrftoken = gjson.ParseBytes(resByte).Get("x-csrftoken").String()
	return cookie, x_csrftoken, err

}

//设置cookie
func SetCookie(req *ApiKey) (err error) {
	req.Cookie, req.XCsrfToken, err = GetCookie(req)
	return err
}

//获取请求客户端
func GetHttpClient(req *ApiKey) *resty.Client {
	if req.IsSsl {
		Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return Client
}
