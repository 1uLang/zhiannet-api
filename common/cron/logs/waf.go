package logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model/edge_db_nodes"
	"github.com/1uLang/zhiannet-api/common/model/logs_statistics"
	"gorm.io/gorm"
	"time"
)

type StatisticsWAFLogs struct {
	Conn *gorm.DB //日志数据库链接
}
type ServerCount struct {
	ServerId int64  `json:"server_id"`
	Total    uint64 `json:"total"`
}

func (s *StatisticsWAFLogs) Run() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	conn, err := edge_db_nodes.NewConn()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	s.Conn = conn

	var total []ServerCount
	timeH := time.Now().Add(-time.Hour)
	timeStr := timeH.Format("20060102")
	sTime := time.Date(timeH.Year(), timeH.Month(), timeH.Day(), timeH.Hour(), 0, 0, 0, time.Local)
	eTime := sTime.Add(time.Hour)
	err = s.Conn.Table("edgeHTTPAccessLogs_"+timeStr).Select("serverId server_id,count(id) total").
		Where("createdAt>=? and createdAt<?", sTime.Unix(), eTime.Unix()).
		Group("serverId").Scan(&total).Error
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(total) > 0 {
		s.Save(total, sTime)
	}

}

func (s *StatisticsWAFLogs) Save(req []ServerCount, t time.Time) {
	if len(req) > 0 {
		for _, v := range req {
			logs_statistics.Save(&logs_statistics.LogsStatistics{
				NodeId: v.ServerId,
				Type:   3,
				Time:   t.Format("2006-01-02 15:04:05"),
				Total:  v.Total,
			})
		}

	}

}
