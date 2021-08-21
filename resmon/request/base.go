package request

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	param "github.com/1uLang/zhiannet-api/resmon/const"
	"github.com/1uLang/zhiannet-api/resmon/model"
)

const (
	// B byte
	B float64 = 1
	// KB k byte
	KB float64 = 1024
	// MB M byte
	MB float64 = 1024 * 1024
	// GB G byte
	GB float64 = 1024 * 1024 * 1024
)

func createReq(reqURL string, aid string) ([]byte, error) {
	if reqURL == param.AGENT_LIST {
		reqURL = fmt.Sprintf("%s/%s", param.BASE_URL, reqURL)
	} else {
		reqURL = fmt.Sprintf("%s/"+reqURL, param.BASE_URL, aid)
	}

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败：%w", err)
	}
	reqQuery := req.URL.Query()
	reqQuery.Add("TeaKey", param.TEA_KEY)
	req.URL.RawQuery = reqQuery.Encode()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	rsp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("获取请求响应失败：%w", err)
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取请求响应体失败：%w", err)
	}

	return body, nil
}

func formatByte(bytes int64) string {
	fb := float64(bytes)
	switch {
	case fb >= GB:
		return fmt.Sprintf("%.2fGB", fb/GB)
	case fb >= MB:
		return fmt.Sprintf("%.2fMB", fb/MB)
	case fb >= KB:
		return fmt.Sprintf("%.2fKB", fb/KB)
	default:
		return fmt.Sprintf("%.2fB", fb)
	}
}

func CheckNodeConn() error {
	body, err := createReq(param.AGENT_LIST, "")
	if err != nil {
		return err
	}

	agents := model.Agents{}
	err = json.Unmarshal(body, &agents)
	if err != nil {
		return errors.New("节点配置信息错误")
	}

	return nil
}
