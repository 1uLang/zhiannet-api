package server

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	param "github.com/1uLang/zhiannet-api/agent/const"
	"github.com/1uLang/zhiannet-api/agent/model"
	cm "github.com/1uLang/zhiannet-api/common/model"
)

// UploadFile 上传文件
func UploadFile(name, format, describe string, body []byte) error {
	// 上传文件流到本地
	sn := generateSN(name)
	var sp string
	fsp := []rune(param.FILE_STORE_PATH)
	if string(fsp[len(fsp)-1]) != "/" {
		sp = param.FILE_STORE_PATH + "/" + sn
	} else {
		sp = param.FILE_STORE_PATH + sn
	}
	fp := filepath.Dir(sp)
	err := os.MkdirAll(fp, 0666)
	if err != nil {
		return fmt.Errorf("创建文件目录失败：%w", err)
	}
	sf, err := os.Create(sp)
	if err != nil {
		return fmt.Errorf("创建存储文件失败：%w", err)
	}
	_, err = sf.Write(body)
	if err != nil {
		return fmt.Errorf("存储文件失败：%w", err)
	}

	// 增加数据记录
	af := model.AgentFile{
		Name:      name,
		Describe:  describe,
		Size:      len(body),
		Format:    format,
		State:     1,
		Path:      sp,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	cdb := cm.MysqlConn.Create(&af)
	if cdb.RowsAffected == 0 {
		return fmt.Errorf("创建文件保存记录失败：%w", cdb.Error)
	}

	return nil
}

// GetFileStream 获取文件流
func GetFileStream(id int64) ([]byte, error) {
	// 通过id获取文件路径
	sp, err := getFPWithID(id)
	if err != nil {
		return nil, err
	}

	// 读取文件获取文件流
	body, err := os.ReadFile(sp)
	if err != nil {
		return nil, fmt.Errorf("获取文件流失败：%w", err)
	}

	return body, err
}

// DeleteFile 删除文件，这里采用伪删除，注释真删除代码
func DeleteFile(id int64) error {
	// 从物理机上删除文件
	// 通过id获取文件路径
	// sp, err := getFPWithID(id)
	// if err != nil {
	// 	return err
	// }
	// err = os.Remove(sp)
	// if err != nil {
	// 	return fmt.Errorf("从物理机上删除文件失败：%w", err)
	// }

	// 变更数据库状态
	udb := cm.MysqlConn.Model(&model.AgentFile{}).Where("id = ?", id).
		Updates(map[string]interface{}{"state": 0, "updated_at": time.Now().Unix()})
	if udb.RowsAffected == 0 {
		return fmt.Errorf("删除文件失败：%w", udb.Error)
	}

	return nil
}

// UpdateFileInfo 更新文件信息
func UpdateFileInfo(name, describe string, id int64) error {
	af := model.AgentFile{
		Name:      name,
		Describe:  describe,
		UpdatedAt: time.Now().Unix(),
	}

	udb := cm.MysqlConn.Model(&model.AgentFile{}).Where("id = ?", id).Updates(af)
	if udb.RowsAffected == 0 {
		return fmt.Errorf("变更文件信息失败：%w", udb.Error)
	}

	return nil
}

// ListFile 文件列表
func ListFile(paging ...int) (*model.FileListRsp, error) {
	var page, size int
	var count int64
	var list []model.AgentFile
	var fs []model.FileInfo
	var rsp model.FileListRsp
	if len(paging) >= 2 {
		page = paging[0]
		size = paging[1]
	}

	// 获取总条数
	err := cm.MysqlConn.Model(&model.AgentFile{}).Where("state = 1").Count(&count).Error
	if err != nil {
		return nil, fmt.Errorf("获取文件总数失败：%w", err)
	}

	if page == 0 || size == 0 {
		err = cm.MysqlConn.Model(&model.AgentFile{}).Where("state = 1").Order("created_at DESC").Find(&list).Error
	} else {
		offset := (page - 1) * size
		err = cm.MysqlConn.Model(&model.AgentFile{}).Where("state = 1").Offset(offset).Limit(size).Order("created_at DESC").Find(&list).Error
	}

	if err != nil {
		return nil, fmt.Errorf("获取文件信息失败：%w", err)
	}

	for _, af := range list {
		fi := model.FileInfo{
			ID:        af.ID,
			Name:      af.Name,
			Describe:  af.Describe,
			Size:      formatBytes(af.Size),
			Format:    af.Format,
			CreatedAt: af.CreatedAt,
		}

		fs = append(fs, fi)
	}

	rsp.Total = count
	rsp.List = fs

	return &rsp, nil
}
