package request

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

const (
	// B byte
	B float64 = 1
	// KB k byte
	KB float64 = 1024
	// MB M byte
	MB float64 = 1024 * 1024
	// GB G byte
	GB float64 = 1024 * 1024 * 1024
)

// GenerateToken 根据用户名密码生成Auth basic
func GenerateToken(req *model.LoginReq) string {
	src := fmt.Sprintf("%s:%s", req.User, req.Password)
	dst := base64.StdEncoding.EncodeToString([]byte(src))

	return fmt.Sprintf("Basic %s", dst)
}

// ParseToken 根据token获取用户名密码
func ParseToken(token string) (string, error) {
	token = strings.TrimSpace(token)
	ts := strings.Split(token, " ")
	if len(ts) < 2 {
		return "", errors.New("token错误")
	}
	token = ts[1]

	src, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	userInfo := strings.Split(string(src), ":")
	if len(userInfo) == 0 {
		return "", errors.New("token解析错误")
	}

	return userInfo[0], nil
}

// FormatTime 解析并格式化时间戳
func FormatTime(timeStr, format string) string {
	timestamp, err := time.Parse(time.RFC1123, timeStr)
	if err != nil {
		return timeStr
	}

	formatTime := timestamp.Format(format)
	return formatTime
}

// FormatBytes 格式化字节大小
func FormatBytes(bytes string) string {
	fb, err := strconv.ParseFloat(bytes, 64)
	if err != nil {
		return ""
	}

	switch {
	case fb > GB:
		return fmt.Sprintf("%.1fGB", fb/GB)
	case fb > MB:
		return fmt.Sprintf("%.1fMB", fb/MB)
	case fb > KB:
		return fmt.Sprintf("%.1fKB", fb/KB)
	default:
		return fmt.Sprintf("%.1fB", fb)
	}
}
