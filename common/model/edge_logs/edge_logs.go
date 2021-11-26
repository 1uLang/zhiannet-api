package edge_logs

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
	edge_users_model "github.com/1uLang/zhiannet-api/edgeUsers/model"
)

type (
	EdgeLogs struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                            //ID
		Level       string `gorm:"column:level" json:"level" form:"level"`                   //级别
		Description string `gorm:"column:description" json:"description" form:"description"` //描述
		CreatedAt   uint64 `gorm:"column:createdAt" json:"createdAt" form:"createdAt"`       //创建时间
		Action      string `gorm:"column:action" json:"action" form:"action"`                //动作
		UserId      uint64 `gorm:"column:userId" json:"userId" form:"userId"`                //用户ID
		AdminId     uint64 `gorm:"column:adminId" json:"adminId" form:"adminId"`             //管理员ID
		ProviderId  uint   `gorm:"column:providerId" json:"providerId" form:"providerId"`    //供应商ID
		Ip          string `gorm:"column:ip" json:"ip" form:"ip"`                            //IP地址
		Type        string `gorm:"column:type" json:"type" form:"type"`                      //类型：admin, user
		Day         string `gorm:"column:day" json:"day" form:"day"`                         //日期
		BillId      uint   `gorm:"column:billId" json:"billId" form:"billId"`                //账单ID
	}
	UserLogReq struct {
		UserId    uint64 `json:"user_id"`
		StartTime uint64 `json:"start_time"`
		EndTime   uint64 `json:"end_time"`
		Keyword   string `json:"keyword"`
		PageNum   int    `json:"page_num"`
		PageSize  int    `json:"page_size"`
	}
	UserLogResp struct {
		EdgeLogs
		UserName string `json:"username" gorm:"column:username"`
	}
)

func GetAll() (list []*UserLogResp, err error) {

	err = model.MysqlConn.Table("edgeLogs").Where("edgeLogs.userId != 0").Where("edgeLogs.level ='info'").
		Joins("left join edgeUsers on edgeUsers.id=edgeLogs.userId").Select("edgeLogs.*,edgeUsers.username").
		Order("edgeLogs.id desc").Find(&list).Error
	return
}
func GetList(req *UserLogReq) (list []*UserLogResp, total int64, err error) {
	parentIds := []uint64{}
	info := &edge_users.EdgeUsers{}
	err = model.MysqlConn.Table("edgeUsers").Where("edgeUsers.id=?", req.UserId).Find(&info).Error
	if err != nil {
		return
	}
	list = make([]*UserLogResp, 0)
	parentId := []*edge_users.EdgeUsers{}
	pid, err := edge_users_model.GetParentId(&edge_users_model.GetParentIdReq{UserId: req.UserId})
	if err != nil {
		return nil, 0, err
	}
	sqlStr := fmt.Sprintf("edgeUsers.id=%v or edgeUsers.parentId= '%v'", req.UserId, req.UserId)
	if pid > 0 {
		sqlStr = fmt.Sprintf("edgeUsers.id=%v or edgeUsers.parentId= '%v'", pid, pid)
	}
	err = model.MysqlConn.Table("edgeUsers").Where(sqlStr).Find(&parentId).Error
	if err != nil {
		return
	}
	for _, v := range parentId {
		parentIds = append(parentIds, v.ID)
	}
	//从数据库获取
	model := model.MysqlConn.Table("edgeLogs").Where("edgeLogs.userId in(?)", parentIds)
	if req.Keyword != "" {
		model = model.Where("edgeLogs.description like ?", "%"+req.Keyword+"%")
	}
	if req.StartTime > 0 {
		model = model.Where("edgeLogs.createdAt>=?", req.StartTime)
	}
	if req.EndTime > 0 {
		model = model.Where("edgeLogs.createdAt<=?", req.EndTime)
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
	model = model.Joins("left join edgeUsers on edgeUsers.id=edgeLogs.userId").Select("edgeLogs.*,edgeUsers.username")
	err = model.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("edgeLogs.id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

func GetNum(req *UserLogReq) (total int64, err error) {
	parentIds := []uint64{}
	parentId := []*edge_users.EdgeUsers{}
	pid, err := edge_users_model.GetParentId(&edge_users_model.GetParentIdReq{UserId: req.UserId})
	sqlStr := fmt.Sprintf("edgeUsers.id=%v or edgeUsers.parentId= '%v'", req.UserId, req.UserId)
	if pid > 0 {
		sqlStr = fmt.Sprintf("edgeUsers.id=%v or edgeUsers.parentId= '%v'", pid, pid)
	}
	err = model.MysqlConn.Table("edgeUsers").Where(sqlStr).Find(&parentId).Error

	for _, v := range parentId {
		parentIds = append(parentIds, v.ID)
	}
	//从数据库获取
	model := model.MysqlConn.Table("edgeLogs").Where("userId in(?)", parentIds)
	if req.Keyword != "" {
		model = model.Where("description like ?", "%"+req.Keyword+"%")
	}
	if req.StartTime > 0 {
		model = model.Where("createdAt>=?", req.StartTime)
	}
	if req.EndTime > 0 {
		model = model.Where("createdAt<=?", req.EndTime)
	}
	err = model.Count(&total).Error

	return
}
