package cache

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/tidwall/gjson"
)

var (
	Rdb *redis.Client
)

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.168.17:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = Rdb.Ping(ctx).Result()
	return err
}

/**
设置缓存
返回参数,,第一个数据,,第二个数据执行结果,第三个是否走缓存
*/
func CheckCache(key string, f interface{}, duration uint32, needCache bool) (interface{}, error, bool) {
	key = Md5Str(key)
	//Rdb.SetNX(context.Background(),key+"-setnx",1, time.Duration(10))
	s, err := GetCache(Rdb, key)
	if needCache && err == nil {
		return s, nil, true
	} else {
		var re interface{}
		Num, ok := f.(func() (interface{}, error))()
		if ok == nil {

			data := make(map[string]interface{})
			data["data"] = Num
			js, jsErr := json.Marshal(data)
			if jsErr != nil {
				//logs.Error("----json.Marshal--", jsErr)
			}
			SetCache(Rdb, key, data, duration)

			dom := gjson.ParseBytes(js)
			re = dom.Get("data").Value()
		} else {
			re = Num
		}

		return re, ok, false
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

func GetCache(Rdb *redis.Client, key string) (interface{}, error) {
	bg := context.Background()
	val, err := Rdb.Get(bg, key).Result()
	if err == nil && val != "" {
		dom := gjson.Parse(val)
		return dom.Get("data").Value(), err
	}

	return nil, err
}

func SetCache(Rdb *redis.Client, key string, data interface{}, duration uint32) (err error) {
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
	fmt.Println("a")
	time.Sleep(time.Second * 3)
	return "a", nil
}
