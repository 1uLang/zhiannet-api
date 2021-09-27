package edge_messages

import (
	"crypto/md5"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
)

type (
	// 消息通知
	Edgemessages struct {
		Id        uint64 `gorm:"column:id" db:"id" json:"id" form:"id"`                             //ID
		Adminid   uint   `gorm:"column:adminId" db:"adminId" json:"adminId" form:"adminId"`         //管理员ID
		Userid    uint   `gorm:"column:userId" db:"userId" json:"userId" form:"userId"`             //用户ID
		Clusterid int    `gorm:"column:clusterId" db:"clusterId" json:"clusterId" form:"clusterId"` //集群ID
		Nodeid    int    `gorm:"column:nodeId" db:"nodeId" json:"nodeId" form:"nodeId"`             //节点ID
		Level     string `gorm:"column:level" db:"level" json:"level" form:"level"`                 //级别
		Subject   string `gorm:"column:subject" db:"subject" json:"subject" form:"subject"`         //标题
		Body      string `gorm:"column:body" db:"body" json:"body" form:"body"`                     //内容
		Type      string `gorm:"column:type" db:"type" json:"type" form:"type"`                     //消息类型
		Params    string `gorm:"column:params" db:"params" json:"params" form:"params"`             //额外的参数
		Isread    int    `gorm:"column:isRead" db:"isRead" json:"isRead" form:"isRead"`             //是否已读
		State     int    `gorm:"column:state" db:"state" json:"state" form:"state"`                 //状态
		Createdat uint64 `gorm:"column:createdAt" db:"createdAt" json:"createdAt" form:"createdAt"` //创建时间
		Day       string `gorm:"column:day" db:"day" json:"day" form:"day"`                         //日期YYYYMMDD
		Hash      string `gorm:"column:hash" db:"hash" json:"hash" form:"hash"`                     //消息内容的Hash
		Role      string `gorm:"column:role" db:"role" json:"role" form:"role"`                     //角色
	}
)

func check(req *Edgemessages) bool {
	var count int64
	_ = model.MysqlConn.Table("edgeMessages").Where("day=?", req.Day).Where("body=?", req.Body).Count(&count).Error
	return count > 0
}
func Add(req *Edgemessages) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}
	//if check(req) {
	//	return 0, nil
	//}
	req.Hash = calHash(req.Role, req.Clusterid, req.Nodeid, req.Subject, req.Body, []byte(req.Params))
	res := model.MysqlConn.Table("edgeMessages").Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

// 计算Hash
func calHash(role string, clusterId int, nodeId int, subject string, body string, paramsJSON []byte) string {
	h := md5.New()
	h.Write([]byte(role + "@" + fmt.Sprintf("%d", clusterId) + "@" + fmt.Sprintf("%d", nodeId)))
	h.Write([]byte(subject + "@"))
	h.Write([]byte(body + "@"))
	h.Write(paramsJSON)
	return fmt.Sprintf("%x", h.Sum(nil))
}
