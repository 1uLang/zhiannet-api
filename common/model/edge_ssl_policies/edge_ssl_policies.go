package edge_ssl_policies

import (
	"github.com/1uLang/zhiannet-api/common/model"
	"strings"
)

type (
	EdgeSSLPolicies struct {
		ID         uint64 `gorm:"column:id" json:"id" form:"id"`
		MinVersion string `gorm:"column:minVersion" json:"minVersion" form:"minVersion"`
	}
)

func CheckAndUpdate(id uint64, check []string, update string) error {
	ssl := EdgeSSLPolicies{}
	err := model.MysqlConn.Table("edgeSSLPolicies").Where("id=?", id).First(&ssl).Error
	if err != nil {
		return err
	}
	if strings.Contains(strings.Join(check," "),ssl.MinVersion) {
		tx := model.MysqlConn.Table("edgeSSLPolicies").Where("id=?", id).Update("minVersion", update)
		return tx.Error
	}
	return nil
}
