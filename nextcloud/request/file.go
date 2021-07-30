package request

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

// ListFolders 列举用户文件 method: PROPFIND
func ListFolders(token string, filePath ...string) (*model.FolderList, error) {
	var lfr model.ListFoldersResp
	var fl model.FolderList
	var fp string
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 拼接用户文件列表路径，默认根目录
	if len(filePath) == 0 {
		fp = "/"
	} else {
		fp = filePath[0]
	}
	lf := fmt.Sprintf(param.LIST_FOLDERS, user)
	uRL := fmt.Sprintf("%s/%s%s", param.BASE_URL, lf, fp)

	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	cli := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequest("PROPFIND", uRL, nil)
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

	for _, v := range lfr.Response {
		var fb model.FolderBody
		unescape, err := url.QueryUnescape(v.Href)
		if err != nil {
			continue
		}
		str := strings.Split(unescape, "/")
		if str[len(str)-1] == "" {
			// fb.Name = str[len(str)-2] + "/"
			// fb.UsedBytes = FormatBytes(v.Propstat[0].Prop.QuotaUsedBytes)
			// 如果是文件夹则不展示，后续不会使用文件夹
			// 默认全都存储在根目录下
			continue
		} else {
			fb.Name = str[len(str)-1]
			fb.UsedBytes = FormatBytes(v.Propstat[0].Prop.Getcontentlength)
		}
		fb.URL = unescape
		fb.ContentType = v.Propstat[0].Prop.Getcontenttype
		fb.LastModified = FormatTime(v.Propstat[0].Prop.Getlastmodified, "2006-01-02 15:04:05")

		fl.List = append(fl.List, fb)
	}
	return &fl, nil
}

// DownLoadFile 下载文件 method: GET
func DownLoadFile(token, fileName string) (*http.Response, error) {
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return nil, err
	}

	// 拼接下载路径，默认为根目录
	df := fmt.Sprintf(param.DOWNLOAD_FILES, user, fileName)
	uRL := fmt.Sprintf("%s/%s", param.BASE_URL, df)

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

// DownLoadFileWithURL 获取文件下载链接
func DownLoadFileWithURL(token, fileName string) (string, error) {
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return "", err
	}

	// 拼接下载路径，默认为根目录
	df := fmt.Sprintf(param.DOWNLOAD_FILES, user, fileName)
	uRL := fmt.Sprintf("%s/%s", param.BASE_URL, df)

	return uRL, nil
}

// UploadFile 上传文件，默认上传到根目录 method: PUT
func UploadFile(token, fileName string, f io.Reader) error {
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return err
	}

	// 拼接上传路径，默认为根目录
	uf := fmt.Sprintf(param.UPLOAD_FILES, user)
	uRL := fmt.Sprintf("%s/%s/%s", param.BASE_URL, uf, fileName)

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

// DeleteFile 删除文件 Method: DELETE
func DeleteFile(token, fileName string) error {
	// 解析token获取用户名
	user, err := ParseToken(token)
	if err != nil {
		return err
	}

	// 拼接删除路径，默认为根目录
	df := fmt.Sprintf(param.DOWNLOAD_FILES, user, fileName)
	uRL := fmt.Sprintf("%s/%s", param.BASE_URL, df)

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
