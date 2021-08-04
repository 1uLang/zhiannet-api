package model

import (
	"errors"
	"fmt"

	"github.com/1uLang/zhiannet-api/common/model"
)

// StoreNCToken 保存nextcloud用户token
func StoreNCToken(name, token string, kind ...uint8) error {
	var kd uint8
	if len(kind) == 0 {
		kd = 0
	} else {
		kd = kind[0]
	}

	nct := NextCloudToken{}
	model.MysqlConn.First(&nct, "user = ?", name)
	nct.User = name
	nct.Token = token
	nct.Kind = kd

	err := model.MysqlConn.Save(&nct).Error
	if err != nil {
		return fmt.Errorf("创建数据备份账号错误：%w", err)
	}

	return nil
}

// BindNCTokenAndUID 绑定nextcloud和uid
func BindNCTokenAndUID(name string, uid int64, kind ...uint8) error {
	var kd uint8
	if len(kind) == 0 {
		kd = 0
	} else {
		kd = kind[0]
	}
	udb := model.MysqlConn.Model(&NextCloudToken{}).Where("user = ? AND uid = 0 AND kind = ?", name, kd).Update("uid", uid)
	if udb.RowsAffected == 0 {
		return fmt.Errorf("绑定token和主站用户错误：%w", udb.Error)
	}

	return nil
}

// QueryTokenByUID 通过主站用户id获取nextcloud token
func QueryTokenByUID(uid int64, kind ...uint8) (string, error) {
	var kd uint8
	if len(kind) == 0 {
		kd = 0
	} else {
		kd = kind[0]
	}

	nct := NextCloudToken{}
	model.MysqlConn.Model(&NextCloudToken{}).Where("uid = ? AND kind = ?", uid, kd).First(&nct)
	if nct.ID == 0 {
		return "", errors.New("获取nextcloud token错误")
	}

	return nct.Token, nil
}
