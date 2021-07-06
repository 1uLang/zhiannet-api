package nat

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

func init() {
	model.InitMysqlLink()
	cache.InitClient()
	//InitTestDB()
}
func InitTestDB() {

	var err error
	dsn := "root:mysql8@tcp(45.195.61.132:3306)/zhian-edges?charset=utf8mb4&parseTime=True&loc=Local"
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
func Test_nat_list(t *testing.T) {
	list, err := GetNat1To1List(&ListReq{NodeId: 2})
	fmt.Println(len(list))
	//fmt.Println(list[1].ID)
	fmt.Println(err)
}

func Test_nat_info(t *testing.T) {
	res, err := GetNat1To1Info(&InfoReq{NodeId: 12, Id: "2"})
	fmt.Println("src=", res.External, ";")
	fmt.Println("src=", "182.12.12.1/24", ";")
	fmt.Println(err)
}

//测试新增
func Test_nat_add(t *testing.T) {
	res, err := SaveNat1To1(&SaveNat1To1Req{
		NodeId:    12,
		ID:        "3",
		Interface: "lan",
		Type:      "nat",
		External:  "1.1.1.1/24",
		Src:       "1.2.3.4",
		Srcmask:   "24",
		Dst:       "12.1.1.1",
		Dstmask:   "24",
		Descr:     "api ",
	})
	fmt.Println(res)
	fmt.Println(err)
}

//启动停止
func Test_start_up(t *testing.T) {
	res, err := StartUpNat1To1(&StartNat1To1Req{NodeId: 12, Id: "3"})
	fmt.Println(res)
	fmt.Println(err)
}

//删除
func Test_del(t *testing.T) {
	res, err := DelNat1To1(&DelNat1To1Req{NodeId: 2, Id: "1"})
	fmt.Println(res)
	fmt.Println(err)
}

//应用
func Test_apply(t *testing.T) {
	res, err := ApplyNat1To1(12)
	fmt.Println(res)
	fmt.Println(err)
}
