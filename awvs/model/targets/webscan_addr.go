package targets

import (
	"fmt"
	_const "github.com/1uLang/zhiannet-api/awvs/const"
	awvs_model "github.com/1uLang/zhiannet-api/awvs/model"
	"github.com/1uLang/zhiannet-api/awvs/request"
	"github.com/1uLang/zhiannet-api/common/model"
	db_model "github.com/1uLang/zhiannet-api/common/model"
	"github.com/tidwall/gjson"
)

type (
	WebscanAddr struct {
		Id          uint64 `gorm:"column:id" json:"id" form:"id"`                                  //id
		TargetId    string `gorm:"column:target_id" json:"target_id" form:"target_id"`             //扫描ID
		UserId      uint64 `gorm:"column:user_id" json:"user_id" form:"user_id"`                   //用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		IsDelete    uint8  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`             //1删除
		CreateTime  int    `gorm:"column:create_time" json:"create_time" form:"create_time"`       //创建时间
	}
	AddrListReq struct {
		UserId      uint64 `json:"user_id" gorm:"column:user_id"`                                  // 用户ID
		AdminUserId uint64 `gorm:"column:admin_user_id" json:"admin_user_id" form:"admin_user_id"` //admin用户ID
		PageNum     int    `json:"page_num" `                                                      //页数
		PageSize    int    `json:"page_size" `                                                     //每页条数
		TargetId    string `gorm:"column:target_id" json:"target_id" form:"target_id"`             //扫描ID
	}
	CheckAddrReq struct {
		Addr        string `json:"-"`
		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
)

//初始化建表
func InitTable() {
	err := db_model.MysqlConn.AutoMigrate(&WebscanAddr{})
	if err != nil {
		fmt.Println("初始化建表，失败：", err.Error())
		return
	}
}

//获取列表信息

//获取节点
func GetList(req *AddrListReq) (list []*WebscanAddr, total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&WebscanAddr{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.TargetId != "" {
			model = model.Where("target_id=?", req.TargetId)
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
	err = model.Debug().Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Order("id desc").Find(&list).Error
	if err != nil {
		return
	}
	return
}

//获取数量
func GetNum(req *AddrListReq) (total int64, err error) {
	//从数据库获取
	model := model.MysqlConn.Model(&WebscanAddr{}).Where("is_delete=?", 0)
	if req != nil {
		if req.UserId > 0 {
			model = model.Where("user_id=?", req.UserId)
		}
		if req.AdminUserId > 0 {
			model = model.Where("admin_user_id=?", req.AdminUserId)
		}
		if req.TargetId != "" {
			model = model.Where("target_id=?", req.TargetId)
		}
	}
	err = model.Count(&total).Error

	return
}

func CheckAddr(arg *CheckAddrReq) (id string,isExist bool, err error) {

	req, err := request.NewRequest()
	if err != nil {
		return "",false, err
	}

	req.Method = "GET"
	req.Url += _const.Targets_api_url
	args := &ListReq{UserId: arg.UserId, AdminUserId: arg.AdminUserId}
	args.Limit = 100
	req.Params = awvs_model.ToMap(args)

	resp, err := req.Do()
	if err != nil {
		return "",false, err
	}
	if args.UserId == 0 && args.AdminUserId == 0 {
		return "",false, err
	}
	//获取数据库 当前用户的扫描用户
	targetList, total, err := GetList(&AddrListReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		PageSize:    999,
		PageNum:     1,
	})
	if total == 0 || err != nil {
		return "",false, err
	}
	tarMap := map[string]int{}
	for _, v := range targetList {
		tarMap[v.TargetId] = 0
	}
	resList := gjson.ParseBytes(resp)
	if resList.Get("targets").Exists() {
		for _, v := range resList.Get("targets").Array() {
			if _, ok := tarMap[v.Get("target_id").String()]; ok && v.Get("address").String() == arg.Addr{
				return v.Get("target_id").String(),true,nil
			}
		}
	}
	return "",false,nil
}

//添加数据
func AddAddr(req *WebscanAddr) (insertId uint64, err error) {
	if req == nil {
		err = fmt.Errorf("参数错误")
		return
	}

	res := model.MysqlConn.Create(&req)
	if res.Error != nil {
		return 0, res.Error
	}
	insertId = req.Id
	return
}

//用target ID 删除
func DeleteByTargetIds(ids []string) (err error) {
	res := model.MysqlConn.Model(&WebscanAddr{}).Where("target_id in (?)", ids).Update("is_delete", 1)
	return res.Error
}

func DeleteByIds(ids []uint64) (err error) {
	res := model.MysqlConn.Model(&WebscanAddr{}).Where("id in (?)", ids).Update("is_delete", 1)
	return res.Error
}
