package encrypt

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func do(param map[string]interface{}) {
	url := "http://192.168.137.8:8999/assets"
	method := "POST"
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}

	buf, _ := json.Marshal(param)
	body := bytes.NewReader(buf)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	//请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	retBuf, err := ioutil.ReadAll(resp.Body)
	ret := struct {
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Message string      `json:"message"`
	}{}
	err = json.Unmarshal(retBuf, &ret)
	if err != nil {
		panic(err)
	}
	if ret.Code != 1 {
		panic(fmt.Errorf(ret.Message))
	}
	fmt.Println("add asset success... ",param["name"])
}
func ipAdd(ip string) string {

	ips := strings.Split(ip,".")
	var add  func(i []int,idx int)

	add = func(i []int,idx int) {
		if idx < 0 {
			return
		}
		if i[idx] > 255 {
			add(i,idx -1)
			return
		}else{
			i[idx] ++
		}
	}
	ints := make([]int,4)
	for k,v := range ips{
		ints[k],_ = strconv.Atoi(v)
	}
	add(ints,3)

	for k,v := range ints {
		ips[k] = strconv.Itoa(v)
	}
	return strings.Join(ips,".")
}
func TestNextTerminal(t *testing.T) {

	param := map[string]interface{}{
		"accountType": "custom",
		"ip":          "192.168.137.8",
		"name":        "opnsense",
		"protocol":    "ssh",
		"ssh-mode":    "",
		"tags":        "",
		"port":        22,
	}
	name := "next_terminal"
	for idx := 0; idx < 100; idx++ {
		param["name"] = fmt.Sprintf("%v-%v",name, strconv.Itoa(idx))
		//param["ip"] = ipAdd(param["ip"].(string))
		do(param)
	}

}
func TestMagicKeyEncode(t *testing.T) {
	var code string
	var year int
	var month int
	var day int
	var hour int
	var err error
	code = "19776ce3be75401480a49cc097a93b54"
	year = 0
	month = 1
	day = 4
	now := time.Now()
	nowt := now.Unix()
	//addt := now.AddDate(year,month,day).Add(time.Duration(hour) * time.Hour)
	addt := now.AddDate(year, month, day).Add(time.Duration(hour) * time.Hour)
	renewal := addt.Unix() - nowt
	timeout := now.Add(5 * time.Minute).Unix()
	dst := MagicKeyEncode([]byte(fmt.Sprintf("%v,%v,%v", code, renewal, timeout)))

	dst = []byte(base64.StdEncoding.EncodeToString(dst))
	fmt.Println(string(dst))
	dst, err = base64.StdEncoding.DecodeString(string(dst))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dst))
	src := MagicKeyDecode(dst)
	fmt.Println(string(src))
}
