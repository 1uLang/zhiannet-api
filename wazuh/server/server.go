package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_messages"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/wazuh/model"
	"github.com/1uLang/zhiannet-api/wazuh/model/groups"
	"github.com/1uLang/zhiannet-api/wazuh/request"
	"time"
)

type CheckRequest struct{}

// SetUrl 初始化 Nessus APIKeys
func SetUrl(url string) error {
	return request.InitServerUrl(url)
}

// SetAPIKeys 初始化 Nessus APIKeys
func InitToken(username, password string) error {
	return request.InitToken(username, password)
}

func GetWazuhInfo() (resp *model.WazuhInfoResp, err error) {
	return &model.WazuhInfoResp{
		Addr:     "https://156.240.95.168",
		Username: "wazuh",
		Password: "AgI_kwQ2GQ8v354EQtd6pSpT7bDjdaNJ",
	}, nil

	//return model.GetWazuhInfo()
}

//检测nessus 配置访问是否异常

func Check() (bool, uint64, int, error) {
	info, err := GetWazuhInfo()
	if err != nil {
		return false, 0, 0, err
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	err = InitToken(info.Username, info.Password)
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	req, err := request.NewRequest()
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	_, err = groups.List(req)
	if err != nil {
		return false, info.Id, info.ConnState, err
	}
	return true, info.Id, info.ConnState, nil
}
func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("wazuh-----------------------------------------------", err)
		}
	}()
	var conn int = 1
	res, id, oldConn, err := Check()
	if id > 0 {
		if !res {
			conn = 0
			_, _ = edge_messages.Add(&edge_messages.Edgemessages{
				Level:     "error",
				Subject:   "组件状态异常",
				Body:      fmt.Sprintf("主机防护状态不可用:%v", err),
				Type:      "AdminAssembly",
				Params:    "{}",
				Createdat: uint64(time.Now().Unix()),
				Day:       time.Now().Format("20060102"),
				Hash:      "",
				Role:      "admin",
			})
		}
		if oldConn != conn {
			subassemblynode.UpdateConnState(id, conn)
			if conn == 1 {
				_, _ = edge_messages.Add(&edge_messages.Edgemessages{
					Level:     "success",
					Subject:   "组件状态恢复正常",
					Body:      "主机防护恢复可用状态",
					Type:      "AdminAssembly",
					Params:    "{}",
					Createdat: uint64(time.Now().Unix()),
					Day:       time.Now().Format("20060102"),
					Hash:      "",
					Role:      "admin",
				})
			}
		}
	}

}
