package request

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
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
