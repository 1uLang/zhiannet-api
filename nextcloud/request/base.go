package request

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/1uLang/zhiannet-api/common/model/subassemblynode"
	"github.com/1uLang/zhiannet-api/utils"

	cm_model "github.com/1uLang/zhiannet-api/common/model"
	param "github.com/1uLang/zhiannet-api/nextcloud/const"
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

type (
	CheckRequest struct{}
)

// GenerateToken 根据用户名密码生成Auth basic
func GenerateToken(req *model.LoginReq) string {
	src := fmt.Sprintf("%s:%s", req.User, req.Password)
	dst := base64.StdEncoding.EncodeToString([]byte(src))

	return fmt.Sprintf("Basic %s", dst)
}

// GetAdminToken 获取nextcloud管理员的token
func GetAdminToken() string {
	req := &model.LoginReq{
		User:     param.AdminUser,
		Password: param.AdminPasswd,
	}

	return GenerateToken(req)
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

// CheckConf 校验配置是否可用
func CheckConf(name, passwd, url string) error {
	var lfr model.ListFoldersResp
	lreq := model.LoginReq{
		User:     name,
		Password: passwd,
	}

	token := GenerateToken(&lreq)
	lURL := fmt.Sprintf("%s/"+param.LIST_FOLDERS, url, name)
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("PROPFIND", lURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	rsp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(body, &lfr)
	if err != nil {
		return err
	}
	if lfr.Response == nil {
		return errors.New("配置错误")
	}
	return nil
}

func (this *CheckRequest) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("next-cloud-----------------------------------------------", err)
		}
	}()
	nodes, _, err := subassemblynode.GetList(&subassemblynode.NodeReq{
		//State:    "1",
		Type:     8,
		PageNum:  1,
		PageSize: 99,
	})
	if err != nil || len(nodes) == 0 {
		err = fmt.Errorf("获取数据备份节点信息失败")
		return
	}
	for _, v := range nodes {
		url := utils.CheckHttpUrl(v.Addr, v.IsSsl == 1)
		err := CheckConf(v.Key, v.Secret, url)
		var conn int = 1
		if err != nil {
			//登录失败 不可用
			conn = 0
		}
		if conn != v.ConnState {
			subassemblynode.UpdateConnState(v.Id, conn)
		}
	}

}

// ConnNextcloudWithAdmin admin与nextcloud建立连接
func ConnNextcloudWithAdmin(name, passwd string) error {
	req := model.LoginReq{
		User:     name,
		Password: passwd,
	}

	token := GenerateToken(&req)
	nct := model.NextCloudToken{}
	un := "admin_admin"
	cm_model.MysqlConn.First(&nct, "user = ?", un)
	nct.User = un
	nct.Token = token
	nct.UID = 1
	nct.Kind = 1

	err := cm_model.MysqlConn.Save(&nct).Error
	if err != nil {
		return fmt.Errorf("与admin关联失败：%w", err)
	}

	return nil
}

// InitialAdminUser 获取数据库中配置的用户名密码
func InitialAdminUser() {
	sn := model.Subassemblynode{}
	cm_model.MysqlConn.Model(&model.Subassemblynode{}).Where("type = 8 AND state = 1 AND is_delete = 0").First(&sn)
	// 判断addr是否发生了变更
	ss := strings.Split(param.BASE_URL, `//`)
	var flag = false
	if ss[1] != sn.Addr {
		flag = true
	}
	if sn.ID > 0 {
		param.AdminUser = sn.Key
		param.AdminPasswd = sn.Secret
		// param.BASE_URL = sn.Addr
		if sn.IsSSL == 1 {
			param.BASE_URL = fmt.Sprintf(`https://%s`, sn.Addr)
		} else {
			param.BASE_URL = fmt.Sprintf(`http://%s`, sn.Addr)
		}
	}
	if flag {
		// 更新用户数据库
		wg := sync.WaitGroup{}
		passwd := `adminAd#@2021`
		token := GetAdminToken()
		ncTokens := []model.NextCloudToken{}
		cm_model.MysqlConn.Model(&model.NextCloudToken{}).Find(&ncTokens)
		if len(ncTokens) <= 100 {
			wg.Add(len(ncTokens))
		} else {
			wg.Add(100)
		}
		for _, v := range ncTokens {
			go func(v model.NextCloudToken) {
				defer wg.Done()
				// 特殊处理admin账号
				if v.UID == 1 && v.Kind == 1 {
					nToken := GenerateToken(&model.LoginReq{
						User:     param.AdminUser,
						Password: param.AdminPasswd,
					})
					cm_model.MysqlConn.Model(&model.NextCloudToken{}).Where("id = ?", v.ID).UpdateColumn("token", nToken)
				} else {
					CreateUserV2(token, v.User, passwd)
				}
			}(v)
		}
		wg.Wait()
	}
}

// func updateNCToken(token, user, passwd string) {
// 	err := CreateUser(token, user, passwd)
// 	nToken := GenerateToken(&model.LoginReq{
// 		User: user,
// 		Password: passwd,
// 	})
// 	if err == nil {
// 		cm_model.MysqlConn.Model(&model.NextCloudToken{}).Where("user = ?",user).UpdateColumn("token",nToken)
// 	}
// }
