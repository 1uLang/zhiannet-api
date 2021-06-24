package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/tidwall/gjson"
	"golang.org/x/sync/singleflight"
)

var (
	Rdb   *redis.Client
	lockG = &singleflight.Group{}
)

func init() {
	InitClient()
}

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "45.195.61.132:6379",
		Password: "1232345342675", // no password set
		DB:       0,               // use default DB
		PoolSize: 100,             // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
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
	fmt.Println("request a")
	time.Sleep(time.Second * 3)
	return "a", nil
}
