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
		Clusterid    int64  `gorm:"column:clusterId" json:"clusterid" form:"clusterid"`
		PwdAt        uint64 `gorm:"column:pwdAt" json:"pwdAt" form:"pwdAt"` //密码修改时间
	}
	ListReq struct {
		Username string `json:"username"`
		PageNum  int    `json:"page_num"`
		PageSize int    `json:"page_size"`
	}
)

func GetList(req *ListReq) (list []*EdgeUsers, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Table("edgeUsers").Where("Ison=?", 1)
	if req != nil {
		if req.Username != "" {
			model = model.Where("username like ?", "%"+req.Username+"%")
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
	err = model.MysqlConn.Table("edgeUsers").Where("id=?", id).First(info).Error
	return
}

//通过用户名获取用户信息
func GetInfoByUsername(name string) (info *EdgeUsers, err error) {
	err = model.MysqlConn.Table("edgeUsers").Where("username=?", name).First(info).Error
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
