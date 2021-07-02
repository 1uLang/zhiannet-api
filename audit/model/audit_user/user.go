package audit_user

import "github.com/1uLang/zhiannet-api/common/model"

type (
	// 用户表
	User struct {
		Id            uint64 `gorm:"column:id" json:"id" form:"id"`
		UserName      string `gorm:"column:user_name" json:"user_name" form:"user_name"`                   //用户名
		Mobile        string `gorm:"column:mobile" json:"mobile" form:"mobile"`                            //中国手机不带国家代码，国际手机号格式为：国家代码-手机号
		UserNickname  string `gorm:"column:user_nickname" json:"user_nickname" form:"user_nickname"`       //用户昵称
		Birthday      int    `gorm:"column:birthday" json:"birthday" form:"birthday"`                      //生日
		CreateTime    int    `gorm:"column:create_time" json:"create_time" form:"create_time"`             //注册时间
		UserPassword  string `gorm:"column:user_password" json:"user_password" form:"user_password"`       //登录密码;cmf_password加密
		UserStatus    uint8  `gorm:"column:user_status" json:"user_status" form:"user_status"`             //用户状态;0:禁用,1:正常,2:未验证
		UserEmail     string `gorm:"column:user_email" json:"user_email" form:"user_email"`                //用户登录邮箱
		Sex           int8   `gorm:"column:sex" json:"sex" form:"sex"`                                     //性别;0:保密,1:男,2:女
		Avatar        string `gorm:"column:avatar" json:"avatar" form:"avatar"`                            //用户头像
		LastLoginTime int    `gorm:"column:last_login_time" json:"last_login_time" form:"last_login_time"` //最后登录时间
		LastLoginIp   string `gorm:"column:last_login_ip" json:"last_login_ip" form:"last_login_ip"`       //最后登录ip
		CustomerId    uint64 `gorm:"column:customer_id" json:"customer_id" form:"customer_id"`             //客户id
		Remark        string `gorm:"column:remark" json:"remark" form:"remark"`                            //备注
		IsAdmin       int8   `gorm:"column:is_admin" json:"is_admin" form:"is_admin"`                      //是否后台管理员 1 是  0   否
		Opt           int8   `gorm:"column:opt" json:"opt" form:"opt"`                                     //是否开启opt双因子认证登录
		Status        int8   `gorm:"column:status" json:"status" form:"status"`                            //删除
		Pid           int8   `gorm:"column:pid" json:"pid" form:"pid"`                                     //创建该用户的id
		Pwd           string `gorm:"column:pwd" json:"pwd" form:"pwd"`                                     //密码
	}
	AuditReq struct {
		UserId uint64 `json:"user_id" ` //waf用户端 用户ID
	}
)

//获取关联信息
func GetInfo(req *AuditReq) (info *User, err error) {
	if req == nil {
		return info, err
	}
	//从数据库获取
	model := model.MysqlConn.Model(&User{}).Where("id=?", req.UserId)
	err = model.First(&info).Error

	return info, err
}
