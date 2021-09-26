package logs

import (
	"github.com/1uLang/zhiannet-api/common/model/logs_statistics"
	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"

	"github.com/1uLang/zhiannet-api/ddos/request/logs"
	"github.com/1uLang/zhiannet-api/ddos/server"
	"time"
)

type StatisticsDDOSLogs struct{}

func (s *StatisticsDDOSLogs) Run() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	//获取ddos节点
	list, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		Type:     1, //ddos
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
		//查询节点下添加的高防ip
		iplist, _, err := ddos_host_ip.GetList(&ddos_host_ip.HostReq{
			NodeId:  v.Id,
			PageNum: 1, PageSize: 999,
		})
		if err != nil || len(iplist) == 0 {
			continue
		}
		//total := uint64(0)
		logReq, err := server.GetLoginInfo(server.NodeReq{NodeId: v.Id})
		if err != nil {
			continue
		}

		eventMap := map[string]int{}

		for _, ipv := range iplist {
			//查询每个IP的日志
			args := &logs.AttackLogReq{
				Addr:      ipv.Addr,
				StartTime: sTime,
				EndTime:   eTime,
			}
			ls, err := logs.AttackLogList(args, logReq, true)
			if err != nil || (ls == nil) {
				continue
			}
			//total += uint64(len(ls.Report))

			for _, v1 := range ls.Report {
				if num, ok := eventMap[v1.Flags]; ok {
					eventMap[v1.Flags] = num + 1
				} else {
					eventMap[v1.Flags] = 1
				}
			}
		}

		if len(eventMap) > 0 {
			for k, v2 := range eventMap {
				nodeTotal = append(nodeTotal, ServerCount{
					Total:    uint64(v2),
					ServerId: int64(v.Id),
					Event:    k,
				})
			}
		}
		//nodeTotal = append(nodeTotal, ServerCount{
		//	Total:    total,
		//	ServerId: int64(v.Id),
		//})

	}
	s.Save(nodeTotal, sTime)

}

func (s *StatisticsDDOSLogs) Save(req []ServerCount, t time.Time) {
	if len(req) > 0 {
		for _, v := range req {
			logs_statistics.Save(&logs_statistics.LogsStatistics{
				NodeId: v.ServerId,
				Type:   1,
				Time:   t.Format("2006-01-02 15:04:05"),
				Total:  v.Total,
				Event:  v.Event,
			})
		}

	}

}
