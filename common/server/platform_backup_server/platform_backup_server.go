package platform_backup_server

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/common/model"
	"github.com/1uLang/zhiannet-api/common/model/platform_backup"
	"time"
)

type PlatformBackUp struct {
}

func (s *PlatformBackUp) Run() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("平台数据凌晨备份失败----------------------------------------------", err)
		}
	}()
	Backup()
}

//初始化表
func InitTable() {
	platform_backup.InitTable()
}

//所有列表
func GetBackupList(req *platform_backup.ListReq) (list []*platform_backup.PlatformBackup, total int64, err error) {
	list, total, err = platform_backup.GetList(req)
	return
}

//备份
func Backup() (row int64, err error) {
	file, size, err := BackupMysqlDb(model.DSN, "", []string{
		"subassemblynode", "edgeSysSettings", "edgeIPLists", "edgeIPItems", //"edgeSysEvents",
	})
	if err != nil {
		return
	}

	backup := &platform_backup.PlatformBackup{
		Name: file,
		Size: size,
	}
	return platform_backup.Save(backup)
}

//恢复文件
func Recovery(id int64) (err error) {
	info, err := platform_backup.GetInfo(id)
	if err != nil {
		return
	}
	return ReductionMysqlDb(model.DSN, info.Name)
}

//删除
func Delete(id int64) (err error) {
	info, err := platform_backup.GetInfo(id)
	if err != nil {
		return
	}
	err = DelFile(info.Name)
	if err != nil {
		return err
	}
	err = platform_backup.Delete([]int64{info.Id})
	return err
}

//清除30天之外的备份
func Clean30DayData() (err error) {
	list, total, err := platform_backup.GetList(&platform_backup.ListReq{
		PageSize: 999,
		PageNum:  1,
		LastTime: time.Now().Add(-time.Hour * 24 * 30),
	})
	ids := []int64{}
	if total > 0 {
		for _, v := range list {
			err = DelFile(v.Name)
			if err != nil {
				return err
			}
			ids = append(ids, v.Id)
		}
	}

	return platform_backup.Delete(ids)
}
