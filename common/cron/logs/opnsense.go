package logs

import (
	"github.com/1uLang/zhiannet-api/common/model/logs_statistics"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/opnsense/request"

	reqips "github.com/1uLang/zhiannet-api/opnsense/request/ips"
	"github.com/1uLang/zhiannet-api/opnsense/server"
	"time"
)

type StatisticsNFWLogs struct{}

func (s *StatisticsNFWLogs) Run() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	//获取 下一代云防火墙 节点
	list, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     2, //opnsense
		State:    "1",
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(list) == 0 {
		return
	}
	var nodeTotal []ServerCount
	timeH := time.Now().Add(-time.Hour)
	sTime := time.Date(timeH.Year(), timeH.Month(), timeH.Day(), timeH.Hour(), 0, 0, 0, time.Local)
	eTime := sTime.Add(time.Hour)
	for _, v := range list {
		total := uint64(0)
		var loginInfo *request.ApiKey
		loginInfo, err = server.GetLoginInfo(server.NodeReq{NodeId: v.Id})
		if err != nil || loginInfo == nil {
			continue
		}
		//设置请求接口必须的cookie
		err = request.SetCookie(loginInfo)
		if err != nil {
			continue
		}
		logs, err := reqips.GetIpsAlarmList(&reqips.IpsAlarmReq{
			IpsReq: reqips.IpsReq{
				RowCount: "-1",
				Current:  "1",
			},
		}, loginInfo)
		if err != nil || logs == nil {
			continue
		}
		for _, lv := range logs.Rows {
			ltime, _ := time.ParseInLocation("2006-01-02T15:04:05.000000+0800", lv.Timestamp, time.Local)
			if ltime.After(sTime) && ltime.Before(eTime) {
				total++
			}
		}
		nodeTotal = append(nodeTotal, ServerCount{
			Total:    total,
			ServerId: int64(v.Id),
		})

	}
	s.Save(nodeTotal, sTime)

}

func (s *StatisticsNFWLogs) Save(req []ServerCount, t time.Time) {
	if len(req) > 0 {
		for _, v := range req {
			logs_statistics.Save(&logs_statistics.LogsStatistics{
				NodeId: v.ServerId,
				Type:   2,
				Time:   t.Format("2006-01-02 15:04:05"),
				Total:  v.Total,
			})
		}

	}

}
