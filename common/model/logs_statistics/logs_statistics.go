package logs_statistics

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	LogsStatistics struct {
		Id         int64  `gorm:"column:id" json:"id" form:"id"`                            //id
		NodeId     int64  `gorm:"column:node_id" json:"node_id" form:"node_id"`             //节点id or server_id
		Time       string `gorm:"column:time" json:"time" form:"time"`                      //时间
		Type       int8   `gorm:"column:type" json:"type" form:"type"`                      //类型：1.ddos、2.下一代防火墙、3.云WAF 4主机
		Total      uint64 `gorm:"column:total" json:"total" form:"total"`                   //日志数量
		Event      string `gorm:"column:event" json:"event" form:"event"`                   //事件
		CreateTime int64  `gorm:"column:create_time" json:"create_time" form:"create_time"` //创建时间
	}

	LogReq struct {
		Type       int     `json:"type"`
		STime      string  `json:"s_time"`
		ETime      string  `json:"e_time"`
		NodeId     []int64 `json:node_id`
		PageNum    int     `json:page_num`
		PageSize   int     `json:page_size`
		ReportType int     `json:"report_type"`
	}

	LogResp struct {
		Times string `json:"times"`
		Total uint64 `json:"total"`
	}
	LogEventResp struct {
		Event string `json:"event"`
		Total uint64 `json:"total"`
	}
)

//获取节点
func GetList(req *LogReq) (list []*LogsStatistics, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&LogsStatistics{}).Where("type=?", req.Type)
	if req.STime != "" {
		model = model.Where("time>=?", req.STime)
	}
	if req.ETime != "" {
		model = model.Where("time<?", req.ETime)
	}
	if len(req.NodeId) > 0 {
		model = model.Where("node_id in (?)", req.NodeId)
	}

	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//修改菜单操作
func Save(req *LogsStatistics) (rows int64, err error) {
	var entity LogsStatistics

	err = model.MysqlConn.Where("node_id=? and type=? and time=? and event=?", req.NodeId, req.Type, req.Time, req.Event).Find(&entity).Error
	if err != nil {
		return
	}
	entity.NodeId = req.NodeId
	entity.Time = req.Time
	entity.Type = req.Type
	entity.Total = req.Total
	entity.Event = req.Event

	if entity.Id == 0 {
		entity.CreateTime = time.Now().Unix()
	}
	res := model.MysqlConn.Model(&LogsStatistics{}).Where("id=?", entity.Id).Save(&entity)
	if res.Error != nil {
		err = res.Error
		return
	}
	rows = res.RowsAffected

	return
}

//按时间统计计数
func GetStatistics(req *LogReq) (list []*LogResp, err error) {
	model := model.MysqlConn.Model(&LogsStatistics{}).Where("type=?", req.Type)
	if req.STime != "" {
		model = model.Where("time>=?", req.STime)
	}
	if req.ETime != "" {
		model = model.Where("time<?", req.ETime)
	}
	if len(req.NodeId) > 0 {
		model = model.Where("node_id in (?)", req.NodeId)
	} else {
		model = model.Where("node_id in (?)", 0)

	}
	if req.ReportType == 0 { //日报
		model = model.Debug().Select("DATE_FORMAT(time,'%Y-%m-%d %H') as times,sum(total) total").Group("times")
	} else {
		model = model.Debug().Select("DATE_FORMAT(time,'%Y-%m-%d') as times,sum(total) total").Group("times")

	}
	err = model.Scan(&list).Error

	return
}

//按时间和事件 统计计数 //目前只有ddos
func GetStatisticsEvent(req *LogReq) (list []*LogEventResp, err error) {
	model := model.MysqlConn.Model(&LogsStatistics{}).Where("type=?", req.Type)
	if req.STime != "" {
		model = model.Where("time>=?", req.STime)
	}
	if req.ETime != "" {
		model = model.Where("time<?", req.ETime)
	}
	if len(req.NodeId) > 0 {
		model = model.Where("node_id in (?)", req.NodeId)
	} else {
		model = model.Where("node_id in (?)", 0)
	}

	model = model.Debug().Select("event,sum(total) total").Group("event")
	err = model.Scan(&list).Error

	return
}
