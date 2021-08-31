package login

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/hids/server/login/util"
	"github.com/tidwall/gjson"
	"net/url"
	"regexp"
	"strings"
	"time"
)

const (
	login_html_url = "/passport/login.html"
	login_api_url  = "/sso/login"
)

type Passport struct{}

//登陆获取cookie
func (this *Passport) Login(req *ApiKey) (CookieMap map[string]string, err error) {
	url := req.Addr
	Client := GetHttpClient(req)
	CookieMap = make(map[string]string)
	resp, err := Client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("Host", "hids.zhiannet.com").
		SetHeader("Origin", url).
		SetHeader("Referer", url+login_html_url).
		Get(url + login_html_url)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`var serviceSid \= \'.*?\'`)
	sid := re.FindString(string(resp.Body()))
	sid = strings.TrimPrefix(sid, "var serviceSid = '")
	sid = strings.TrimSuffix(sid, "'")

	cb := fmt.Sprintf("jQuery30009331923499973542_%v", time.Now().UnixNano()/1000000)
	//登陆 返回cookies
	resp, err = Client.SetDebug(true).R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("Host", "hids.zhiannet.com").
		SetHeader("Origin", url).
		SetHeader("Referer", url+login_html_url).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		SetQueryString(params2UrlEncode(map[string]string{
			"sid":      sid,
			"service":  "https://hids.zhiannet.com/passport/yunleiIndex",
			"callback": cb,
			"ajax":     "yes",
			"_":        fmt.Sprintf("%v", time.Now().UnixNano()/1000000),
		})).
		Get(url + login_api_url)
	if err != nil {
		return CookieMap, err
	}
	rest := map[string]interface{}{}
	if resp.StatusCode() == 200 {
		re := regexp.MustCompile(cb + `\(.*?\)`)
		jq := re.FindString(string(resp.Body()))
		jq = strings.TrimPrefix(jq, cb+"(")
		jq = jq[:len(jq)-1]
		fmt.Println(jq)
		err := json.Unmarshal([]byte(jq), &rest)
		if err != nil {
			return nil, err
		}
	}
	//登陆 返回cookies
	resp, err = Client.SetDebug(true).R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetHeader("Host", "hids.zhiannet.com").
		SetHeader("Origin", url).
		SetHeader("Referer", url+login_html_url).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		SetQueryString(params2UrlEncode(map[string]string{
			"callback":  cb,
			"userName":  req.Username,
			"password":  util.Encode([]byte(fmt.Sprintf("%x", md5.Sum([]byte(req.Password+"safedog"))))),
			"vcode":     "1234",
			"sid":       sid,
			"service":   "https://hids.zhiannet.com/passport/yunleiIndex",
			"execution": rest["execution"].(string),
			"_eventId":  rest["_eventId"].(string),
			"access":    "1",
			"ajax":      "yes",
			"lt":        rest["lt"].(string),
			"_":         fmt.Sprintf("%v", time.Now().UnixNano()/1000000),
		})).
		Get(url + login_api_url)

	if err != nil {
		return CookieMap, err
	}
	fmt.Println(string(resp.Body()))
	if resp.StatusCode() == 200 {
		//获取cookie
		Cookies := resp.Cookies()
		if len(Cookies) > 0 {
			CookieMap["cookie"] = Cookies[0].Value
		}
		fmt.Println("cookies", Cookies)
	}
	fmt.Println(CookieMap["cookie"])
	//resp, err = Client.R().
	//	SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
	//	SetHeader("Host", "hids.zhiannet.com").
	//	SetHeader("Origin", url).
	//	SetHeader("Cookie",base64.URLEncoding.EncodeToString([]byte("hy_data_2020_id=17ae6c03a058f9-0fb68e4840c342-2343360-2073600-17ae6c03a06e53;"+
	//		" hy_data_2020_js_sdk={\"distinct_id\":\"17ae6c03a058f9-0fb68e4840c342-2343360-2073600-17ae6c03a06e53\",\"site_id\":1142,"+
	//		"\"user_company\":1275,\"props\":{},\"device_id\":\"17ae6c03a058f9-0fb68e4840c342-2343360-2073600-17ae6c03a06e53\"}; "))+
	//		fmt.Sprintf("SESSION=%s; JSESSIONID=934BE728FC0A6B6CE9C39F39F62AB01B.safedogserver2", CookieMap["cookie"])).
	//	SetHeader("Referer", "https://hids.zhiannet.com/cloudeyes/riskInvade/attackAnalysis/attackAnalysisIndex?r_c_m=1").
	//	SetFormData(map[string]string{
	//		"pageSize":  "10",
	//		"pageNo":    "1",
	//		"startTime": "2021-08-19",
	//		"endTime":   "2021-08-25",
	//	}).Post("https://hids.zhiannet.com//cloudeyes/riskInvade/attackAnalysis/queryAttackAnalysisPage?_t_=0.7768512503905523")
	//
	//if err != nil {
	//	return CookieMap, err
	//}
	//fmt.Println(string(resp.Body()))
	//if resp.StatusCode() == 200 {
	//	rest := map[string]interface{}{}
	//	err = json.Unmarshal(resp.Body(), &rest)
	//	if err != nil {
	//		return nil, err
	//	}
	//}
	return CookieMap, err
}

//获取cookie和接口凭证 x-csrftoken
func (this *Passport) GetCookie(req *ApiKey) (cookie, x_csrftoken string, err error) {

	key := fmt.Sprintf("opnsense-cookie-%v:%v", req.Addr, req.Port)
	var resp interface{}
	resp, err = cache.CheckCache(key, func() (interface{}, error) {
		return this.Login(req)
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
func (this *Passport) SetCookie(req *ApiKey) (err error) {
	req.Cookie, req.XCsrfToken, err = this.GetCookie(req)
	return err
}
func params2UrlEncode(params map[string]string) string {

	q := (&url.URL{}).Query()
	for k, v := range params {
		q.Add(k, fmt.Sprintf("%v", v))
	}
	return q.Encode()
}
