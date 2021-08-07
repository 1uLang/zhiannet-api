package request

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

// ListFoldersWithPath 通过url获取文件列表
func ListFoldersWithPath(token string, filePath ...string) (*model.FolderList, error) {
	var lfr model.ListFoldersResp
	var fl model.FolderList
	if param.BASE_URL == "" || param.AdminUser == "" || param.AdminPasswd == "" {
		return nil, errors.New("该组件暂未添加，请添加后重试")
	}
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 拼接用户文件列表路径，默认根目录
	var uRL string
	if len(filePath) == 0 {
		uRL = fmt.Sprintf("%s/"+param.LIST_FOLDERS+"/", param.BASE_URL, user)
	} else {
		fp := filePath[0]
		if fp == "" {
			fp = fmt.Sprintf("/"+param.LIST_FOLDERS+"/", user)
		} else {
			if string([]rune(fp)[len([]rune(fp))-1]) != "/" {
				return nil, errors.New("请填写正确的文件夹路径")
			}
		}
		uRL = fmt.Sprintf("%s%s", param.BASE_URL, fp)
	}

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	reqBody, err := os.ReadFile("xml/list_file.xml")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PROPFIND", uRL, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	rsp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &lfr)
	if err != nil {
		return nil, err
	}

	for _, v := range lfr.Response[1:] {
		var fb model.FolderBody
		unescape, err := url.QueryUnescape(v.Href)
		if err != nil {
			continue
		}
		str := strings.Split(unescape, "/")
		if str[len(str)-1] == "" {
			fb.Name = str[len(str)-2] + "/"
			fb.UsedBytes = FormatBytes(v.Propstat.Prop.QuotaUsedBytes)
		} else {
			fb.Name = str[len(str)-1]
			fb.UsedBytes = FormatBytes(v.Propstat.Prop.Getcontentlength)
		}
		fb.FileID = v.Propstat.Prop.FileID
		fb.URL = unescape
		fb.ContentType = v.Propstat.Prop.Getcontenttype
		fb.LastModified = FormatTime(v.Propstat.Prop.Getlastmodified, "2006-01-02 15:04:05")

		fl.List = append(fl.List, fb)
	}
	return &fl, nil
}

// DownLoadFileWithPath 下载文件 method: GET
func DownLoadFileWithPath(token, path string) (*http.Response, error) {
	// 拼接下载路径，默认为根目录
	if string([]rune(path)[len([]rune(path))-1]) == "/" {
		return nil, errors.New("请输入正确的文件路径")
	}
	uRL := fmt.Sprintf("%s%s", param.BASE_URL, path)

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("GET", uRL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)

	return cli.Do(req)
}

// DownLoadFileURLWithPath 根据path获取文件url
func DownLoadFileURLWithPath(path string) (string, error) {
	if path == "" {
		return "", errors.New("请输入正确的文件路径")
	}
	// 拼接下载路径，默认为根目录
	if string([]rune(path)[len([]rune(path))-1]) == "/" {
		return "", errors.New("请输入正确的文件路径")
	}
	uRL := fmt.Sprintf("%s%s", param.BASE_URL, path)

	return uRL, nil
}

// DeleteFileWithPath 删除文件 Method: DELETE
func DeleteFileWithPath(token, path string) error {
	// 解析token获取用户名
	_, err := ParseToken(token)
	if err != nil {
		return err
	}

	if path == "" {
		return errors.New("文件或文件夹路径不能为空")
	}
	// 拼接删除路径，默认为根目录
	uRL := fmt.Sprintf("%s%s", param.BASE_URL, path)

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("DELETE", uRL, nil)
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
		return fmt.Errorf("读取删除文件响应错误：%w", err)
	}

	delErr := model.DeleteFileError{}
	err = xml.Unmarshal(body, &delErr)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return fmt.Errorf("解码删除文件响应错误：%w", err)
	}

	return fmt.Errorf("删除文件错误：%s", delErr.Message)
}

// UploadFileWithPath 上传文件，默认上传到根目录 method: PUT
func UploadFileWithPath(token, fileName string, f io.Reader, dirPath ...string) error {
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return err
	}

	// 拼接上传路径，默认为根目录
	var uRL string
	if len(dirPath) == 0 {
		uRL = fmt.Sprintf("%s/"+param.UPLOAD_FILES+"/%s", param.BASE_URL, user, fileName)
	} else {
		fp := dirPath[0]
		if fp == "" {
			fp = fmt.Sprintf("/"+param.UPLOAD_FILES+"/%s", user, fileName)
		} else {
			if string([]rune(fp)[len([]rune(fp))-1]) != "/" {
				return errors.New("请填写正确的文件夹路径")
			}
			fp += fileName
		}
		uRL = fmt.Sprintf("%s%s", param.BASE_URL, fp)
	}

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("PUT", uRL, f)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)

	_, err = cli.Do(req)
	return err
}
