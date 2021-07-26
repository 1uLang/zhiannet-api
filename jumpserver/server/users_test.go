package server

import (
	"context"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	users_model "github.com/1uLang/zhiannet-api/jumpserver/model/users"
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
	var err error
	req, err = NewServerRequest("http://182.150.0.106:8080", "admin", "21ops.com")
	if err != nil {
		panic(err)
	}
}

func TestUserList(t *testing.T) {

	args := &users_model.ListReq{}
	list, err := req.Users.List(args)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	fmt.Println(list, len(list))
}
func TestUserCreate(t *testing.T) {
	args := &users_model.CreateReq{Name: "llu", Username: "lusir", Password: "123456", Email: "243971996@qq.com"}
	list, err := req.Users.Create(args)
	if err != nil {
		t.Fatal(err)
		t.Fail()
	}
	fmt.Println(list, len(list))
}
