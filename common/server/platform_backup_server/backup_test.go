package platform_backup_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/cache"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/util"
	"testing"
)

func init() {
	model.InitMysqlLink()
	cache.InitClient()
}

//备份文件
func TestBackup(t *testing.T) {
	res, err := Backup()
	fmt.Println(res, err)

}
func TestBackupMySqlDb(t *testing.T) {
	file, size, err := BackupMysqlDb(model.DSN, "edges", []string{
		"subassemblynode", "edgeSysSettings", "edgeIPLists", "edgeIPItems", "edgeSysEvents",
	})
	fmt.Println(file, size, err)

}

//测试解压
func TestUnZip(t *testing.T) {
	filename := "./edges-202109221.zip"
	err := util.Unzip(filename, "./")
	fmt.Println(err)
}

//测试恢复
func TestRecover(t *testing.T) {
	err := Recovery(1)
	fmt.Println(err)

}
func TestReductionMySqlDb(t *testing.T) {
	dsn := "root:mysql8@tcp(127.0.0.1:3306)/deges-1?charset=utf8mb4&parseTime=True&loc=Local"
	err := ReductionMysqlDb(dsn, "edges-20210922.zip")
	fmt.Println(err)

}

//获取文件大小
func Test_file_size(t *testing.T) {
	GetFileSize("./edges-20210922.zip")
}
