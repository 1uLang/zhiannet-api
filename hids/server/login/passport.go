package request

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"time"
)

var Client = resty.New().SetDebug(false).SetTimeout(time.Second * 60)

var publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCYJzVL82l9rPGTvJKtaSM2p27v
ujEsJhlq8QgUHB9958ZVg1i0t5wPhbJsK0ASlRLPa7jIV2rNxKSoqZR8Jkhj9Xm8
ipUX0+qlf5r6z9vHSa29UaWSBH4QznSxkKB0jhdISnwcVVlBSxuwOj0uVgqjIkK4
6E4fNu3yGx1FfL9rXQIDAQAB
-----END PUBLIC KEY-----
`

func encrypt(pswd string) string {
	//获取public key
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic("failed to parse DER encoded public key: " + err.Error())
	}

	//类型断言
	publicKey := pub.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(fmt.Sprintf("%x",md5.Sum([]byte(pswd+"safedog")))))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cipherText))
	return string(cipherText)
	//返回密文
	//return "QoEKoQs9ichARCIlruqwinurNVr1kE14PzxBD1igRUbzIUIJYcy5Kb50p9F+DF5dGEG03rDueAE6WDcTJAje0Z1LTkCB74jeOHzOaFju4eXBhYP5gGfg2/54fs3m2gddOrYEg1E+fYoU5qSSYDIZW9gRhXDOsshjJrC9zplFCLQ="
}

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
	name, value := "", ""
	//fmt.Println(doc.Html())
	doc.Find("div class=\"yunlei_login_nr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		name, _ = s.Attr("userName")
		value, _ = s.Attr("password")
		fmt.Printf("Review %d: %s - %s\n", i, name, value)
	})
	//通过标签匹配获取key - value
	//input := doc.Find("form input[type='hidden']").First()
	//key, _ = input.Attr("name")
	//value, _ = input.Attr("value")
	////登陆 返回cookies
	resp, err := Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("Host", "hids.zhiannet.com").
		SetHeader("Origin", "https://hids.zhiannet.com").
		SetHeader("Referer", "https://hids.zhiannet.com/manager/login").
		SetHeader("Referer", "https://hids.zhiannet.com/manager/login").
		SetFormData(map[string]string{
			"userName":     req.Username,
			"password":     encrypt(req.Password),
			"validateCode": "6xmb",
		}).Post("https://hids.zhiannet.com/manager/loginSubmit" + fmt.Sprintf("?__t=0.5070945322296552"))

	if err != nil {
		return CookieMap, err
	}
	fmt.Println("code========", resp.StatusCode())
	fmt.Println(string(resp.Body()))
	fmt.Println("key=", name, "value=", value)

	if resp.StatusCode() == 200 {
		//获取cookie
		Cookies := resp.Cookies()
		if len(Cookies) > 0 {
			CookieMap["cookie"] = Cookies[0].Value
			CookieMap["x-csrftoken"] = value //接口调用凭证
		}

		fmt.Println("cookies", Cookies)
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
