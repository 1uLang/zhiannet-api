package edge_db_nodes

import "github.com/1uLang/zhiannet-api/common/model"

type EdgeDBNodes struct {
	Id          uint   `gorm:"column:id" json:"id" form:"id"`                            //ID
	Ison        uint8  `gorm:"column:isOn" json:"isOn" form:"isOn"`                      //是否启用
	Role        string `gorm:"column:role" json:"role" form:"role"`                      //数据库角色
	Name        string `gorm:"column:name" json:"name" form:"name"`                      //名称
	Description string `gorm:"column:description" json:"description" form:"description"` //描述
	Host        string `gorm:"column:host" json:"host" form:"host"`                      //主机
	Port        uint   `gorm:"column:port" json:"port" form:"port"`                      //端口
	Database    string `gorm:"column:database" json:"database" form:"database"`          //数据库名称
	Username    string `gorm:"column:username" json:"username" form:"username"`          //用户名
	Password    string `gorm:"column:password" json:"password" form:"password"`          //密码
	Charset     string `gorm:"column:charset" json:"charset" form:"charset"`             //通讯字符集
	Conntimeout uint   `gorm:"column:connTimeout" json:"connTimeout" form:"connTimeout"` //连接超时时间（秒）
	State       uint8  `gorm:"column:state" json:"state" form:"state"`                   //状态
	Createdat   uint64 `gorm:"column:createdAt" json:"createdAt" form:"createdAt"`       //创建时间
	Weight      uint   `gorm:"column:weight" json:"weight" form:"weight"`                //权重
	Order       uint   `gorm:"column:order" json:"order" form:"order"`                   //排序
	Adminid     uint   `gorm:"column:adminId" json:"adminId" form:"adminId"`             //管理员ID
}

//获取节点详细信息
func GetNodeInfo() (info *EdgeDBNodes, err error) {
	err = model.MysqlConn.Table("edgeDBNodes").Where("isOn=?", 1).First(&info).Error
	return
}
