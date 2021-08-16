package model

import (
	//"fmt"
	"fmt"

	"github.com/1uLang/zhiannet-api/common/model"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	//"github.com/spf13/viper"
	//"gorm.io/driver/mysql"
)

// var (
// 	db  *gorm.DB
// 	dsn string
// )

// func init() {
// 加载配置文件
//readConfig()

// 初始化数据库
// var err error
// db, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
// if err != nil {
// 	panic(fmt.Errorf("连接数据库失败：%w", err))
// }

// 读取数据库中的节点配置
// getAdminUser()
// }

//func readConfig() {
//	viper.SetConfigFile(param.DB_CONFIG_PATH)
//	err := viper.ReadInConfig()
//	if err != nil {
//		panic(fmt.Errorf("加载配置文件错误：%w", err))
//	}
//
//	dsn = viper.GetString("dbs.prod.dsn")
//}

// InitialAdminUser 获取数据库中配置的用户名密码
func InitialAdminUser() {
	sn := Subassemblynode{}
	model.MysqlConn.Model(&Subassemblynode{}).Where("type = 8 AND state = 1 AND is_delete = 0").First(&sn)
	if sn.ID > 0 {
		param.AdminUser = sn.Key
		param.AdminPasswd = sn.Secret
		// param.BASE_URL = sn.Addr
		if sn.IsSSL == 1 {
			param.BASE_URL = fmt.Sprintf(`https://%s`, sn.Addr)
		} else {
			param.BASE_URL = fmt.Sprintf(`http://%s`, sn.Addr)
		}
	}
}
