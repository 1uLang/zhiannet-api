package util

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/encrypt"
	"strconv"
	"strings"
	"time"
)

const DBNodePasswordEncodedPrefix = "EDGE_ENCODED:"

func Interface2Int(i interface{}) (int, error) {
	return strconv.Atoi(fmt.Sprintf("%v", i))
}

func Interface2Uint64(i interface{}) (uint64, error) {
	return strconv.ParseUint(fmt.Sprintf("%v", i), 10, 64)
}

// EncodePassword 加密密码
func EncodePassword(password string) string {
	if strings.HasPrefix(password, DBNodePasswordEncodedPrefix) {
		return password
	}
	encodedString := base64.StdEncoding.EncodeToString(encrypt.MagicKeyEncode([]byte(password)))
	return DBNodePasswordEncodedPrefix + encodedString
}

// DecodePassword 解密密码
func DecodePassword(password string) string {
	if !strings.HasPrefix(password, DBNodePasswordEncodedPrefix) {
		return password
	}
	dataString := password[len(DBNodePasswordEncodedPrefix):]
	data, err := base64.StdEncoding.DecodeString(dataString)
	if err != nil {
		return password
	}
	return string(encrypt.MagicKeyDecode(data))
}

//SHA256生成哈希值
func GetSHA256HashCode(message []byte) string {
	//方法一：
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(message)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode

	//方法二：
	//bytes2:=sha256.Sum256(message)//计算哈希值，返回一个长度为32的数组
	//hashcode2:=hex.EncodeToString(bytes2[:])//将数组转换成切片，转换成16进制，返回字符串
	//return hashcode2
}

func GetFirstDateOfWeek() (weekMonday time.Time, err error) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}

	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday = weekStartDate
	return
}
