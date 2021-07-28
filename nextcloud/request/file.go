package request

import (
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

	cli := &http.Client{}
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
		unescape, err := url.QueryUnescape(v.Href.Text)
		if err != nil {
			continue
		}
		str := strings.Split(unescape, "/")
		if str[len(str)-1] == "" {
			fb.Name = str[len(str)-2] + "/"
		} else {
			fb.Name = str[len(str)-1]
		}
		fb.URL = unescape
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

	cli := &http.Client{}
	req, err := http.NewRequest("GET", uRL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	return cli.Do(req)
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

	cli := &http.Client{}
	req, err := http.NewRequest("PUT", uRL, f)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)

	_, err = cli.Do(req)
	return err
}
