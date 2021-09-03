package logs_statistics_server

import (
	"github.com/1uLang/zhiannet-api/common/model/logs_statistics"
	"time"
)

type (
	Statistics struct {
		Time  string `json:"time"`
		Value uint64 `json:"value"`
	}
)

//report 0日报  1周报
func GetWafStatistics(serverIds []int64, report, logsType int) (res []*Statistics, err error) {
	res = make([]*Statistics, 0)
	var sTime, eTime time.Time
	NowTime := time.Now()
	if report == 0 {
		toDay := NowTime.Add(-24 * time.Hour)
		sTime = time.Date(toDay.Year(), toDay.Month(), toDay.Day(), 0, 0, 0, 0, time.Local)
		eTime = sTime.Add(24 * time.Hour)
	} else {
		Day7 := NowTime.Add(-24 * 7 * time.Hour)
		sTime = time.Date(Day7.Year(), Day7.Month(), Day7.Day(), 0, 0, 0, 0, time.Local)
		eTime = sTime.Add(24 * 7 * time.Hour)
	}
	list, err := logs_statistics.GetStatistics(&logs_statistics.LogReq{
		Type:       logsType,
		NodeId:     serverIds,
		ReportType: report,
		STime:      sTime.Format("2006-01-02 15:04:05"),
		ETime:      eTime.Format("2006-01-02 15:04:05"),
	})

	if list == nil {
		list = make([]*logs_statistics.LogResp, 0)
	}

	//处理缺失数据默认值
	//timeMap := map[string]string{}
	//timeReport := make([]*Statistics,0)
	if report == 0 { //日报 每小时一份

		for i := 0; i < 24; i++ {
			//timeMap[sTime.Format("2006-01-02 15")] = sTime.Format("2006-01-02 15")
			value := uint64(0)
			for _, v := range list {
				if v.Times == sTime.Format("2006-01-02 15") {
					value = v.Total
					break
				}
			}
			res = append(res, &Statistics{
				Time:  sTime.Format("2006-01-02 15:04:05"),
				Value: value,
			})
			sTime = sTime.Add(time.Hour)
		}

	} else {
		for i := 0; i < 7; i++ {
			//timeMap[sTime.Format("2006-01-02")] = sTime.Format("2006-01-02")
			value := uint64(0)
			for _, v := range list {
				if v.Times == sTime.Format("2006-01-02") {
					value = v.Total
					break
				}
			}
			res = append(res, &Statistics{
				Time:  sTime.Format("2006-01-02"),
				Value: value,
			})
			sTime = sTime.Add(time.Hour * 24)
		}
	}
	return
}
