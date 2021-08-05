package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/sync/singleflight"
)

var (
	Rdb   *redis.Client
	lockG = &singleflight.Group{}
)

type (
	RdbConfig struct {
		Redis struct {
			Addr     string `yaml:"addr"`
			Password string `yaml:"password"`
		} `yaml:"redis"`
	}
)

var ApiDbPath = "./build/configs/api_db.yaml"

//func init() {
//InitClient()
//}

// 初始化连接
func InitClient() (err error) {
	var yamlFile []byte
	conf := new(RdbConfig)
	yamlFile, err = ioutil.ReadFile(ApiDbPath)
	//yamlFile, err = ioutil.ReadFile("./build/configs/api_db.yaml")
	//yamlFile, err = ioutil.ReadFile("/Users/yons/zhian/zhiannet-edge-line/EdgeAdmin/build/configs/api_db.yaml")

	if err != nil {
		return err
		//panic(fmt.Errorf("zhiannet package redis link yamlFile.Get err #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &conf)

	if err != nil {
		//panic(fmt.Errorf("zhiannet package redis link yaml.Unmarshal err %v", err))
		return err
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,     //"45.195.61.132:6379",
		Password: conf.Redis.Password, //"1232345342675", // no password set
		DB:       0,                   // use default DB
		PoolSize: 100,                 // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		//panic(fmt.Errorf("zhiannet-api package link redis err %v", err))
	}
	return err
}

/**
设置缓存
返回参数,,第一个数据,,第二个数据执行结果
*/
func CheckCache(key string, fn func() (interface{}, error), duration uint32, needCache bool) (interface{}, error) {
	key = Md5Str(key)
	s, err := GetCache(key)
	if needCache && err == nil {
		return s, nil
	} else {
		var re interface{}
		//Num, ok := fn()
		//同一时间只有一个带相同key的函数执行 防击穿
		Num, ok, _ := lockG.Do(key, fn)
		if ok == nil {
			SetCache(key, Num, duration)
			re = Num
		} else {
			re = Num
		}

		return re, ok
	}

}

/**
md5
*/
func Md5Str(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func GetCache(key string) (interface{}, error) {
	bg := context.Background()
	val, err := Rdb.Get(bg, key).Result()
	if err == nil && val != "" {
		dom := gjson.Parse(val)
		return dom.Get("data").Value(), err
	}

	return "", err
}

func SetCache(key string, data interface{}, duration uint32) (err error) {
	bg := context.Background()
	dataMap := make(map[string]interface{})
	dataMap["data"] = data
	var js []byte
	js, err = json.Marshal(dataMap)
	if err != nil {
		return err
	} else {
		err = Rdb.Set(bg, key, js, time.Duration(duration)*time.Second).Err()
		if err != nil {
			return err
		}
	}
	return err
}

func getA() (a interface{}, err error) {
	fmt.Println("audit_db a")
	time.Sleep(time.Second * 3)
	return "a", nil
}

//加锁key 值为1
func SetNx(key string, t time.Duration) (res bool, err error) {
	//key = Md5Str(key)
	ctx := context.Background()
	res, err = Rdb.SetNX(ctx, key, 1, t).Result()
	return
}

//key 值+1
func Incr(key string, t time.Duration) (res int64, err error) {
	//key = Md5Str(key)
	ctx := context.Background()
	//return Rdb.Incr(ctx, key).Result()

	pipe := Rdb.TxPipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, t)

	// Execute
	//
	//     MULTI
	//     INCR pipeline_counter
	//     EXPIRE pipeline_counts 3600
	//     EXEC
	//
	// using one rdb-server roundtrip.
	_, err = pipe.Exec(ctx)
	//fmt.Println(incr.Val(), err)
	res = incr.Val()
	return
}

//获取key的int值
func GetInt(key string) (res int, err error) {
	//key = Md5Str(key)
	ctx := context.Background()
	var result string
	result, err = Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		res, err = 0, nil
	} else {
		res, _ = strconv.Atoi(result)
	}

	return
}

func DelKey(key string) (err error) {
	ctx := context.Background()
	_, err = Rdb.Del(ctx, key).Result()
	return
}
