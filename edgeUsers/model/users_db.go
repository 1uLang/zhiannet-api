package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"time"
)

type (
	Edgeusers struct {
		Id           uint64 `gorm:"column:id" json:"id" form:"id"`                               // id
		IsOn         uint8  `gorm:"column:isOn" json:"isOn" form:"isOn"`                         // 是否启用
		Username     string `gorm:"column:username" json:"username" form:"username"`             // 用户名
		Password     string `gorm:"column:password" json:"password" form:"password"`             // 密码
		Name         string `gorm:"column:fullname" json:"fullname" form:"fullname"`             // 真实姓名
		Mobile       string `gorm:"column:mobile" json:"mobile" form:"mobile"`                   // 手机号
		Tel          string `gorm:"column:tel" json:"tel" form:"tel"`                            // 联系电话
		Remark       string `gorm:"column:remark" json:"remark" form:"remark"`                   // 备注
		Email        string `gorm:"column:email" json:"email" form:"email"`                      // 邮箱地址
		AvatarFileId uint64 `gorm:"column:avatarFileId" json:"avatarFileId" form:"avatarFileId"` // 头像文件ID
		CreatedAt    int64  `gorm:"column:createdAt" json:"createdAt" form:"createdAt"`          // 创建时间
		Day          string `gorm:"column:day" json:"day" form:"day"`                            // YYMMDD
		UpdatedAt    int64  `gorm:"column:updatedAt" json:"updatedAt" form:"updatedAt"`          // 修改时间
		State        uint8  `gorm:"column:state" json:"state" form:"state"`                      // 状态 删除
		Source       string `gorm:"column:source" json:"source" form:"source"`                   // 来源
		ClusterId    uint32 `gorm:"column:clusterId" json:"clusterId" form:"clusterId"`          // 集群ID
		Features     string `gorm:"column:features" json:"features" form:"features"`             // 允许操作的特征
		ParentId     uint64 `gorm:"column:parentId" json:"parentId" form:"parentId"`             // 创建该用户的父级用户iD
		PwdAt        uint64 `gorm:"column:pwdAt" json:"pwdAt" form:"pwdAt"`                      // 创建该用户的父级用户iD
	}
	EdgeusersResp struct {
		Edgeusers
		OtpOn     int8   `gorm:"column:otpIsOn" json:"otpIsOn" form:"otpIsOn"`
		OtpParams string `gorm:"column:otpParams" json:"otpParams" form:"otpParams"`
	}
	GetNumReq struct {
		UserId uint64
	}
	ListReq struct {
		UserId uint64
		Offset int
		Size   int
		All    bool
	}
	CheckUserNameReq struct {
		UserId   uint64
		Username string
	}
	UpdateUserReq struct {
		Id       uint64
		IsOn     uint8
		Password string
		Name     string
		Mobile   string
		Tel      string
		Remark   string
		Email    string
	}
	CreateUserReq struct {
		UserId   uint64
		IsOn     uint8
		Password string
		Name     string
		Mobile   string
		Tel      string
		Remark   string
		Email    string
		Source   string
		Fullname string
		Username string
	}
	DeleteUserReq struct {
		UserId uint64
	}
	FindUserFeaturesReq struct {
		UserId int64
	}
	UpdateUserFeaturesReq struct {
		UserId   int64
		Features string
	}
	GetParentIdReq struct {
		UserId uint64
	}
)

var edgeUserTableName = "edgeUsers"

func InitTable() {
	err := model.MysqlConn.Exec("alter table edgeUsers add parentId  bigint DEFAULT null COMMENT '父级id';").Error
	if err != nil {
		fmt.Println("更新edgeUsers表字段 parentId，失败：", err.Error())
		return
	}
}

//获取节点
func GetList(req *ListReq) (list []*EdgeusersResp, err error) {
	if req.UserId == 0 {
		return
	}
	//判断是否为子账号
	all := req.All
	uid := req.UserId
	parent := Edgeusers{}
	err = model.MysqlConn.Debug().Table(edgeUserTableName).Where("edgeUsers.id=?", req.UserId).Find(&parent).Error
	if err != nil {
		return nil, err
	}
	if parent.ParentId != 0 {
		req.UserId = parent.ParentId
	}
	//从数据库获取
	db_model := model.MysqlConn.Debug().Table(edgeUserTableName).Where("edgeUsers.state=?", 1)
	if req != nil && req.UserId > 0 {
		if all {
			db_model = db_model.Where("edgeUsers.parentId=? or edgeUsers.id=?", req.UserId, uid)
		} else {
			db_model = db_model.Where("edgeUsers.parentId=?", req.UserId)
		}
	} else {
		return
	}
	var total int64

	err = db_model.Count(&total).Error
	if err != nil || total == 0 {
		return
	}
	if req.Offset <= 0 {
		req.Offset = 0
	}
	if req.Size <= 0 {
		req.Size = 20
	}
	db_model = db_model.Joins("left join edgeLogins on edgeUsers.id=edgeLogins.userId").Select("edgeUsers.*,edgeLogins.isOn as otpIsOn,edgeLogins.params as otpParams")
	err = db_model.Debug().Offset(req.Offset).Limit(req.Size).Order("edgeUsers.id asc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//获取数量
func GetNum(req *GetNumReq) (total int64, err error) {
	//从数据库获取
	db_model := model.MysqlConn.Table(edgeUserTableName).Where("state=?", 1)
	if req != nil && req.UserId > 0 {
		db_model = db_model.Where("parentId=?", req.UserId)
	} else {
		return
	}
	err = db_model.Count(&total).Error

	return
}
func CheckUserUsername(req *CheckUserNameReq) (bool, error) {
	if req.Username == "" {
		return false, fmt.Errorf("参数错误")
	}
	user_model := model.MysqlConn.Table(edgeUserTableName).Where("username=?", req.Username).Where("state=1")
	if req.UserId != 0 {
		user_model = user_model.Where("id!=?", req.UserId)
	}
	var total int64
	err := user_model.Count(&total).Error
	return err == nil && total > 0, err
}

func createMd5(source string) string {
	hash := md5.New()
	hash.Write([]byte(source))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func UpdateUser(req *UpdateUserReq) error {
	if req == nil || req.Id == 0 {
		return fmt.Errorf("参数错误")
	}
	ent := Edgeusers{}
	err := model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.Id).Find(&ent).Error
	if err != nil {
		return err
	}
	ent.Name = req.Name
	ent.IsOn = req.IsOn
	if req.Password != "" {
		ent.Password = createMd5(req.Password)
		ent.PwdAt = uint64(time.Now().Unix())
	}
	ent.Mobile = req.Mobile
	ent.Remark = req.Remark
	ent.Email = req.Email
	ent.UpdatedAt = time.Now().Unix()
	if ent.Features == "" {
		ent.Features = "[]"
	}
	return model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.Id).Save(&ent).Error
}
func GetParentId(req *GetParentIdReq) (uint64, error) {
	if req.UserId == 0 {
		return 0, fmt.Errorf("参数错误")
	}
	user := Edgeusers{}
	err := model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.UserId).Find(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ParentId, nil
}
func CreateUser(req *CreateUserReq) (uint64, error) {

	if req.UserId == 0 {
		return 0, fmt.Errorf("参数错误")
	}
	parent := Edgeusers{}
	err := model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.UserId).Find(&parent).Error
	if err != nil {
		return 0, err
	}

	op := Edgeusers{}
	if parent.ParentId == 0 {
		op.ParentId = req.UserId
	} else {
		op.ParentId = parent.ParentId
	}
	op.Username = req.Username
	op.Password = createMd5(req.Password)
	op.Name = req.Fullname
	op.Mobile = req.Mobile
	op.Tel = req.Tel
	op.Email = req.Email
	op.Remark = req.Remark
	op.Source = req.Source
	//默认使用上级用户的节点
	op.ClusterId = parent.ClusterId
	op.CreatedAt = time.Now().Unix()
	op.Day = time.Now().Format("20060102")
	op.IsOn = 1
	op.State = 1
	op.Features = "[]"
	op.PwdAt = uint64(time.Now().Unix())
	res := model.MysqlConn.Table(edgeUserTableName).Create(&op)

	if res.Error != nil {
		return 0, res.Error
	}
	return op.Id, nil
}
func DeleteUser(req *DeleteUserReq) error {
	res := model.MysqlConn.Table(edgeUserTableName).Where("id in (?)", req.UserId).Update("state", 0)
	return res.Error
}

func FindUserFeatures(req *FindUserFeaturesReq) ([]string, error) {
	if req.UserId == 0 {
		return nil, fmt.Errorf("参数错误")
	}
	ent := Edgeusers{}
	err := model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.UserId).Find(&ent).Error
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0)
	if ent.Features != "" {
		err = json.Unmarshal([]byte(ent.Features), &ret)
	}
	return ret, err
}
func UpdateUserFeatures(req *UpdateUserFeaturesReq) error {

	if req.UserId == 0 || !json.Valid([]byte(req.Features)) {
		return fmt.Errorf("参数错误")
	}

	return model.MysqlConn.Table(edgeUserTableName).Where("id=?", req.UserId).Update("features", req.Features).Error
}
