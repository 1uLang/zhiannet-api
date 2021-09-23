package platform_backup_server

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
	"github.com/1uLang/zhiannet-api/nextcloud/request"
	"io"
	"io/ioutil"
	"os"
	"strings"

	//"github.com/JamesStewy/go-mysqldump"
	mysqldump "github.com/1uLang/zhiannet-api/common/util"
	_ "github.com/go-sql-driver/mysql"
)

//生成备份文件
func BackupMysqlDb(dsn, dbname string, tableName []string) (filename, size string, err error) {

	dumpDir := "./"                                         // you should create this directory
	dumpFilenameFormat := fmt.Sprintf("%s20060102", dbname) // accepts time layout string and add .sql at the end of file

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return filename, size, err
	}

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return filename, size, err
	}
	// Dump database to file
	resultFilename, err := dumper.Dump(tableName)
	if err != nil {
		fmt.Println("Error dumping:", err)
		return filename, size, err
	}
	//fmt.Printf("File is saved to %s", resultFilename)
	//删除原文件
	defer os.Remove(resultFilename)

	// Close dumper and connected database
	dumper.Close()
	//压缩文件
	filename = strings.TrimRight(resultFilename, ".sql") + ".zip"
	err = mysqldump.ZipFile(resultFilename, filename)
	if err != nil {
		err = errors.New("创建备份文件出错")
		return filename, size, err
	}
	defer os.Remove(filename)

	//计算文件大小
	size, _ = GetFileSize(filename)

	//上传文件
	err = UpFile(filename)
	if err != nil {
		fmt.Println(err)
		err = errors.New("上传文件出错2")
		return filename, size, err
	}

	return filename, size, err
}

//按照文件内容还原
func ReductionMysqlDb(dsn, filename string) (err error) {
	//一些 mysql 驱动默认是不支持multi statements的需要进行配置，因为 multi statements 可能会增加sql注入的风险
	//解决办法
	//需要加入参数 multiStatements=true
	if strings.Index(dsn, "multiStatements=true") <= 0 {
		dsn = dsn + "&multiStatements=true"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return err
	}
	if db.Ping() != nil {
		err = fmt.Errorf("链接数据库错误，恢复失败")
		return err
	}

	//下载到本地
	err = DowFile(filename)
	if err != nil {
		return err
	}
	defer os.Remove(filename)

	//解压文件
	err = mysqldump.Unzip(filename, "./")
	if err != nil {
		return err
	}

	name := strings.TrimRight(filename, ".zip") + ".sql"
	fi, err := os.Open(name)
	if err != nil {
		return err
	}
	defer fi.Close()
	defer os.Remove(name)
	content, err := ioutil.ReadAll(fi)
	cont := string(content)

	//fmt.Println(cont)
	//return
	return mysqldump.Reduction(db, &cont)

}

//上传备份文件
func UpFile(file string) (err error) {
	//上传文件 使用admin账号
	token, err := model.QueryTokenByUID(1, 1)
	if err != nil {
		err = errors.New("上传文件出错1")
		return err
	}
	dirPath := "/remote.php/dav/files/admin/平台数据/"
	fileName := strings.TrimLeft(file, "./")
	//先创建文件夹
	err = request.CreateFoler(token, "", "平台数据")
	if err != nil {
		if err.Error() == "存在同名文件夹" {
			err = nil
		} else {
			return err
		}

	}

	//读取文件
	by, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = request.UploadFileWithPath(token, fileName, bytes.NewReader(by), dirPath)
	return err
}

//下载文件到本地
func DowFile(filename string) (err error) {
	//上传文件 使用admin账号
	token, err := model.QueryTokenByUID(1, 1)
	if err != nil {
		err = errors.New("下载文件出错1")
		return err
	}

	uRL := fmt.Sprintf(`/remote.php/dav/files/admin/平台数据/%v`, filename)
	rsp, err := request.DownLoadFileWithPath(token, uRL)
	if err != nil {
		return
	}

	bb, err := io.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(filename, bb, 0644)
	return err
}

//删除文件
func DelFile(file string) (err error) {
	//删除文件 使用admin账号
	token, err := model.QueryTokenByUID(1, 1)
	if err != nil {
		err = errors.New("删除文件出错1")
		return err
	}
	url := fmt.Sprintf(`/remote.php/dav/files/admin/平台数据/%v`, file)

	err = request.DeleteFileWithPath(token, url)
	return err
}

func GetFileSize(filename string) (size string, err error) {
	//计算大小
	fi, err := os.Stat(filename)
	if err != nil {
		return size, err
	}
	fb := fi.Size()
	fbn := float64(fb)
	GB, MB, KB := float64(1024*1024*1024), float64(1024*1024), float64(1024)
	switch {
	case fbn >= GB:
		size = fmt.Sprintf("%.1fGB", fbn/GB)
	case fbn >= MB:
		size = fmt.Sprintf("%.1fMB", fbn/MB)
	case fbn >= KB:
		size = fmt.Sprintf("%.1fKB", fbn/KB)
	default:
		size = fmt.Sprintf("%.1fB", fbn)
	}
	return size, nil

}
