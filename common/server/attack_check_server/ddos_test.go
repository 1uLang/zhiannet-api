package attack_check_server

import (
	"context"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/go-redis/redis/v8"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func InitTestDB() {

	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/edge?charset=utf8mb4&parseTime=True&loc=Local"
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

	dsn = "root:mysql8@tcp(45.195.61.132:3306)/gfast_open_test?charset=utf8mb4&parseTime=True&loc=Local"
	model.AuditMysqlConn, err = gorm.Open(gmysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表前缀
			SingularTable: true, //表名复数形式
		},
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("审计系统 mysql link err ")
	}

	cache.Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.137.8:6379",
		//Password: "1232345342675", // no password set
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
func TestDdosAttackCheck(t *testing.T)  {
	InitTestDB()
	err := ddos{}.AttackCheck()
	if err != nil {
		t.Fatal(err)
	}
}
