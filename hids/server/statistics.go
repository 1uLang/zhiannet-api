package server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/hids/model/risk"
	"github.com/1uLang/zhiannet-api/hids/request"
)

/*
	safagod 主机入侵分析数据统计
*/
type StatisticsRequest struct{}

func (this *StatisticsRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("hids statistics -----------------------------------------------", err)
		}
	}()

	//init server
	info, err := GetHideInfo()
	if err != nil {
		return
	}
	err = SetUrl(info.Addr)
	if err != nil {
		return
	}
	err = SetAPIKeys(&request.APIKeys{info.AppId, info.Secret})
	if err != nil {
		return
	}

	risk.Statistics()
}
