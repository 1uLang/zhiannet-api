package host_status

import (
	"context"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/go-redis/redis/v8"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"testing"
	"time"
)

func InitDB() {
	model.InitMysqlLink()
	cache.InitClient()
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
func Test_ddos_list(t *testing.T) {
	model.InitMysqlLink()
	list, _, err := GetDdosNodeList()
	fmt.Println("data ====", list[0])
	fmt.Println(list, err)
}
func Test_ddos_info(t *testing.T) {
	model.InitMysqlLink()
	info, err := GetDDoSNodeInfo(11)
	fmt.Println(info, err)
}

//主机状态
func Test_host_status(t *testing.T) {
	//GetHostStatus()
}

//主机列表
func Test_host_list(t *testing.T) {
	InitDB()
	list, total, err := GetHostList(&ddos_host_ip.HostReq{NodeId: 1})
	fmt.Println(list[0])
	fmt.Println(total)
	fmt.Println(err)
}

//添加ip
func Test_add_addr(t *testing.T) {
	InitDB()
	list, err := AddAddr(&ddos_host_ip.AddHost{NodeId: 11, Addr: "118.112.240.13", UserId: 1})
	fmt.Println(list)
	fmt.Println(err)
}
func TestDeleteAddr(t *testing.T) {
	InitTestDB()

	err := DeleteAddr([]uint64{31})

	fmt.Println(err)
}

//屏蔽列表
func Test_get_host_shield_list(t *testing.T) {
	InitDB()
	req := &ShieldReq{NodeId: 11, Addr: "118.112.250.80", Page: 1}
	list, err := GetHostShieldList(req)
	fmt.Println(list)
	fmt.Println(err)
}

//释放屏蔽列表
func Test_ReleaseShield(t *testing.T) {
	InitDB()
	req := &ReleaseShieldReq{NodeId: 11, Addr: []string{"118.112.250"}}
	err := ReleaseShield(req)
	fmt.Println(err)
}

//链接列表
func Test_GetLinkList(t *testing.T) {
	InitDB()
	req := &LinkReq{NodeId: 1, Addr: "182.150.0.36:50451-103.80.27.105:6051"}
	list, err := GetLinkList(req)
	fmt.Println(list)
	fmt.Println(err)
}

//主机详细信息
func Test_host_info(t *testing.T) {
	InitDB()
	req := &HostGetReq{
		NodeId: 1,
		Addr:   "118.112.240.1",
	}
	res, err := GetHostInfo(req)

	fmt.Println(res.Ignore)
	fmt.Println(res.ParamSet)
	fmt.Println(res.FilterSet)
	fmt.Println(res.PortproSetTcp)
	fmt.Println(res.PortproSetUdp)
	fmt.Println(err)
}

//主机设置
func Test_HostSet(t *testing.T) {
	InitDB()
	req := &HostSetReq{
		NodeId:     1,
		Addr:       "118.112.240.1",
		Ignore:     true,
		ProtectSet: 0,
		FilterSet:  0,
		SetTcp:     0,
		SetUdp:     0,
	}
	list, err := SetHost(req)
	fmt.Println(list)
	fmt.Println(err)
}
