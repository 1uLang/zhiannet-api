package scans

import (
	"context"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/nessus/model/scans"
	"github.com/1uLang/zhiannet-api/nessus/request"
	"github.com/1uLang/zhiannet-api/nessus/server"
	"github.com/go-redis/redis/v8"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func InitMysql() {
	var err error
	dsn := "root:123456@tcp(192.168.137.1:3306)/edge?charset=utf8mb4&parseTime=True&loc=Local"
	model.MysqlConn, err = gorm.Open(gmysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表前缀
			SingularTable: true, //表名复数形式
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Errorf("zhiannet-api package link mysql err %v", err))
	}

	cache.Rdb = redis.NewClient(&redis.Options{
		Addr:     "45.195.61.132:6379",
		Password: "1232345342675", // no password set
		DB:       0,               // use default DB
		PoolSize: 100,             // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = cache.Rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("zhiannet-api package link redis err %v", err))
	}

}
func init() {
	InitMysql()
	//初始化 nessus 服务器地址
	err := server.SetUrl("https://156.240.95.239:8834")
	if err != nil {
		panic(err)
	}
	//初始化 nessus 系统管理员账号apikeys
	err = server.SetAPIKeys(&request.APIKeys{
		Access: "4caa54e1df36556950450ee00d5e6e22b55fdcb81d940e3999c51c743782288c",
		Secret: "c96b82f6507e2249647ef5e50f32642504fbaad37d5cfee28f58c630545f9ebd",
	})
	if err != nil {
		panic(err)
	}
}

func TestCreate(t *testing.T) {

	args := &scans.AddReq{}
	args.UserId = 2
	args.AdminUserId = 0
	args.Settings.Name = "cess"
	args.Settings.Text_targets = "https://www.zhianmy.com"
	args.Settings.Description = "成都官网"

	err := Create(args)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
