package util

import (
	"encoding/base64"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/encrypt"
	"strconv"
	"strings"
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
