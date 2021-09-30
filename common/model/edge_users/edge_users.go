package edge_users

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	EdgeUsers struct {
		ID           uint64 `gorm:"column:id" json:"id" form:"id"`
		Ison         int64  `gorm:"column:isOn" json:"ison" form:"ison"`
		Username     string `gorm:"column:username" json:"username" form:"username"`
		Password     string `gorm:"column:password" json:"password" form:"password"`
		Fullname     string `gorm:"column:fullname" json:"fullname" form:"fullname"`
		Mobile       string `gorm:"column:mobile" json:"mobile" form:"mobile"`
		Tel          string `gorm:"column:tel" json:"tel" form:"tel"`
		Remark       string `gorm:"column:remark" json:"remark" form:"remark"`
		Email        string `gorm:"column:email" json:"email" form:"email"`
		Avatarfileid int64  `gorm:"column:avatarFileid" json:"avatarfileid" form:"avatarfileid"`
		Createdat    int64  `gorm:"column:createdAt" json:"createdat" form:"createdat"`
		Updatedat    int64  `gorm:"column:updatedAt" json:"updatedat" form:"updatedat"`
		State        int64  `gorm:"column:state" json:"state" form:"state"`
		Source       string `gorm:"column:source" json:"source" form:"source"`
		Clusterid    uint64 `gorm:"column:clusterId" json:"clusterid" form:"clusterid"`
		ParentId     uint64 `gorm:"column:parentId" json:"parentid" form:"parentid"`
		PwdAt        uint64 `gorm:"column:pwdAt" json:"pwdat" form:"pwdat"`             //密码修改时间
		ChannelId    uint64 `gorm:"column:channelId" json:"channelId" form:"channelId"` //渠道ID
	}
	ListReq struct {
		Username  string `json:"username"`
		PageNum   int    `json:"page_num"`
		PageSize  int    `json:"page_size"`
		ParentId  uint64 `json:"parent_id"`
		ChannelId uint64 `json:"channel_id"`
	}

	ChanUserTotal struct {
		ChannelId uint64 `json:"channel_id"`
		Total     int64  `json:"total"`
	}

	SubUserTotal struct {
		ParentId uint64 `json:"parent_id"`
		Total    int64  `json:"total"`
	}
)

func GetList(req *ListReq) (list []*EdgeUsers, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Debug().Table("edgeUsers").Where("state=? ", 1)
	if req != nil {
		if req.Username != "" {
			model = model.Where("(username LIKE ? OR fullname LIKE ? OR mobile LIKE ? OR email LIKE ? OR tel LIKE ? OR remark LIKE ?)", "%"+req.Username+"%", "%"+req.Username+"%", "%"+req.Username+"%", "%"+req.Username+"%", "%"+req.Username+"%", "%"+req.Username+"%")
		}
		if req.ParentId > 0 {
			model = model.Where("parentId=?", req.ParentId)
		} else {
			model = model.Where(" parentId is null")
		}
		if req.ChannelId > 0 {
			model = model.Where("channelId=?", req.ChannelId)
		}

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
	return
}

func GetListByUid(req []uint64) (resMap map[uint64]*EdgeUsers, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("edgeUsers").Where("isOn=?", 1)
	if len(req) == 0 {
		return
	}
	model = model.Where("id in(?)", req)
	err = model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	var list []*EdgeUsers
	err = model.Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	resMap = make(map[uint64]*EdgeUsers)
	for _, v := range list {
		resMap[v.ID] = v
	}
	return
}

//通过id获取用户信息
func GetInfoById(id uint64) (info *EdgeUsers, err error) {
	err = model.MysqlConn.Debug().Table("edgeUsers").Where("id=?", id).First(&info).Error
	return
}

//通过用户名获取用户信息
func GetInfoByUsername(name string) (info *EdgeUsers, err error) {
	err = model.MysqlConn.Table("edgeUsers").Where("username=?", name).Order("state desc").First(&info).Error
	return
}

//更新账号密码修改时间
func UpdatePwdAt(id uint64) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeUsers").Where("id=?", id).Update("pwdAt", time.Now().Unix())
	row = tx.RowsAffected
	err = tx.Error
	return
}

//更新账号密码
func UpdatePwd(id uint64, pwd string) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeUsers").Where("id=?", id).Updates(map[string]interface{}{"pwdAt": time.Now().Unix(), "password": pwd})
	row = tx.RowsAffected
	err = tx.Error
	return
}

//更新渠道ID
func UpdateChannel(id, chanId uint64) (row int64, err error) {
	tx := model.MysqlConn.Table("edgeUsers").Where("id=?", id).Updates(map[string]interface{}{"updatedAt": time.Now().Unix(), "channelId": chanId})
	row = tx.RowsAffected
	err = tx.Error
	return
}

//统计渠道用户数 subUse 是否包含子账户
func GetChannelUserTotal(channelId []uint64, subUse bool) (total []ChanUserTotal, err error) {
	total = make([]ChanUserTotal, 0)
	model := model.MysqlConn.Table("edgeUsers").Where("state=? ", 1)
	if !subUse {
		model = model.Where(" parentId is null")
	}
	model = model.Where("channelId in(?)", channelId).Group("channelId").
		Select("count(id) total,channelId channel_id")
	err = model.Scan(&total).Error
	return
}

//统计子账号用户数
func GetSubUserTotal(parentId []uint64) (total []SubUserTotal, err error) {
	total = make([]SubUserTotal, 0)
	model := model.MysqlConn.Table("edgeUsers").Where("state=?", 1)
	model = model.Where("parentId in(?)", parentId).Group("parentId").
		Select("count(id) total,parentId parent_id")
	err = model.Scan(&total).Error
	return
}
