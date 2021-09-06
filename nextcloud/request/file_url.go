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
	"strconv"
	"strings"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

// ListFoldersWithPath 通过url获取文件列表
func ListFoldersWithPath(token string, filePath ...string) (*model.FolderList, error) {
	getNCInfo()
	var lfr model.ListFoldersResp
	var fl model.FolderList
	if param.BASE_URL == "" || param.AdminUser == "" || param.AdminPasswd == "" {
		return nil, errors.New("该组件暂未添加，请添加后重试")
	}
	// 解析token获取用户名
	user, _, err := ParseToken(token)
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

	req, err := http.NewRequest("PROPFIND", uRL, bytes.NewReader([]byte(listFileXML)))
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

	fl.Quota, fl.Used, fl.Percent = GetNCUserInfo(token, user)

	return &fl, nil
}

// DownLoadFileWithPath 下载文件 method: GET
func DownLoadFileWithPath(token, path string) (*http.Response, error) {
	getNCInfo()
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
	getNCInfo()
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
	getNCInfo()
	// 解析token获取用户名
	_, _, err := ParseToken(token)
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
	getNCInfo()
	// 解析token获取用户名
	user, _, err := ParseToken(token)
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

// GetDirectDownloadURL 获取下载直链
func GetDirectDownloadURL(fileID int64, token string) (string, error) {
	getNCInfo()
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}

	uRL := fmt.Sprintf("%s/%s", param.BASE_URL, param.Direct_Download)
	req, err := http.NewRequest("POST", uRL, nil)
	if err != nil {
		return "", fmt.Errorf("新建请求失败：%w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("OCS-APIRequest", "true")
	reqQuery := req.URL.Query()
	reqQuery.Add("fileId", strconv.FormatInt(fileID, 10))
	req.URL.RawQuery = reqQuery.Encode()
	rsp, err := cli.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求执行失败：%w", err)
	}
	defer rsp.Body.Close()

	dRsp := model.DirectResp{}
	rBody, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", fmt.Errorf("获取响应体失败：%w", err)
	}
	err = xml.Unmarshal(rBody, &dRsp)
	if err != nil {
		return "", fmt.Errorf("xml解析错误：%w", err)
	}

	if dRsp.Meta.Statuscode != 200 {
		return "", errors.New(dRsp.Meta.Message)
	}

	dURL, err := url.QueryUnescape(dRsp.Data.URL)
	if err != nil {
		return "", fmt.Errorf("url解析失败：%w", err)
	}

	return dURL, nil
}

// CreateFoler 创建文件夹，pfURL：父级目录url
func CreateFoler(token, pfURL, folerName string) error {
	getNCInfo()
	// 解析token获取用户名
	user, _, err := ParseToken(token)
	if err != nil {
		return err
	}
	pfURL = strings.TrimSpace(pfURL)
	if pfURL == "" {
		pfURL = fmt.Sprintf("%s/"+param.LIST_FOLDERS+"/", param.BASE_URL, user)
	}
	if string([]rune(pfURL)[len([]rune(pfURL))-1]) != "/" {
		return errors.New("传入的父级url不是目录")
	}
	nfURL := fmt.Sprintf("%s%s", pfURL, folerName)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("MKCOL", nfURL, nil)
	if err != nil {
		return fmt.Errorf("新建请求失败：%w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("OCS-APIRequest", "true")
	resp, err := cli.Do(req)
	if err != nil {
		return fmt.Errorf("请求执行失败：%w", err)
	}
	defer resp.Body.Close()

	return nil
}
