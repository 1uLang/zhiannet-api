package request

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	param "github.com/1uLang/zhiannet-api/resmon/const"
	"github.com/1uLang/zhiannet-api/resmon/model"
	"github.com/1uLang/zhiannet-api/resmon/server"
	"io"
	"net/http"
	"time"
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
	server.GetNodeInfo()
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

type CheckRequest struct {
}

func (*CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("teaweb-节点监控----------------------------------------------", err)
		}
	}()
	info, _ := server.GetNodeInfo()
	if info.ID > 0 {
		var conn int = 1
		body, err := createReq(param.AGENT_LIST, "")
		fmt.Println("body", string(body), "err", err)
		if err != nil {
			//err   {"data":{},"errors":null,"code":400,"message":"Authenticate Failed 002"}
			conn = 0
			if info.ConnState == 1 {
				//修改
				subassemblynode.UpdateConnState(uint64(info.ID), conn)
			}
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "节点监控状态不可用",
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})
			return
		}

		agents := &model.Agents{}
		err = json.Unmarshal(body, &agents)
		if err != nil || agents == nil {
			conn = 0
			if info.ConnState == 1 {
				//修改
				subassemblynode.UpdateConnState(uint64(info.ID), conn)
			}
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      "节点监控状态不可用",
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})
			return
		}

		if info.ConnState == 0 && conn == 1 {
			edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "success",
				Subject:   "组件状态恢复正常",
				Body:      "节点监控恢复可用状态",
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})

		}
		if conn != info.ConnState {
			subassemblynode.UpdateConnState(uint64(info.ID), conn)

		}
	}

}
