package request

import (
	"bytes"
	"crypto/tls"
	_ "embed"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"os"
	"testing"
	"time"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

var (
	req = &model.LoginReq{
		User: "admin",
		// User: "admin_zhoumj",
		// Password: "Dengbao123!@#",
		// Password: "admin",
		Password: "21ops.com@",
		// Password: "adminAd#@2021",
	}
	fileName = "Nextcloud.png"
	//go:embed Nextcloud.png
	nexcloud string
)

func TestFormatTime(t *testing.T) {
	ts := FormatTime("Sat, 24 Jul 2021 13:53:13 GMT", "2006-01-02 15:04:05")

	t.Log(ts)
}

func TestFormatBytes(t *testing.T) {
	str := FormatBytes("5656463")
	t.Log(str)
}

func TestToken(t *testing.T) {
	token := GenerateToken(req)
	user, _, err := ParseToken(token)
	if err != nil {
		t.Fatal(err)
	}
	if user != "admin" {
		t.Fail()
	}
	t.Log(token)
}

func TestDownLoadFile(t *testing.T) {
	token := GenerateToken(req)
	rsp, err := DownLoadFile(token, fileName)
	if err != nil {
		t.Fatal(err)
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != nexcloud {
		t.Fail()
	}
}

func TestDownLoadFileWithURL(t *testing.T) {
	token := GenerateToken(req)
	url, err := DownLoadFileWithURL(token, fileName)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(url)
}

func TestUploadFile(t *testing.T) {
	// param.BASE_URL = `https://bptest.dengbao.cloud`
	token := GenerateToken(req)
	by, err := os.ReadFile("golang.png")
	if err != nil {
		t.Fatal(err)
	}

	err = UploadFile(token, "golang.png", bytes.NewReader(by))
	if err != nil {
		t.Fatal(err)
	}
}

func TestListFolders(t *testing.T) {
	token := GenerateToken(req)
	// param.BASE_URL = "http://localhost:8088"
	// param.AdminUser = "admin"
	// param.AdminPasswd = "admin"
	lf, err := ListFolders(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range lf.List {
		t.Logf("%d %s\n", v.FileID, v.URL)
		t.Log(v.LastModified)
	}
}

func TestDeleteFile(t *testing.T) {
	token := GenerateToken(req)
	fileName := "golang.png"

	err := DeleteFile(token, fileName)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	token := GenerateToken(req)
	// 用户名只能是：“a-z”，“A-Z”，“0-9”和"_.@-'"
	// 线上可用用户的主键或sn编码作为用户名
	userNamer := "hanchan"
	passwd := "123456"

	err := CreateUser(token, userNamer, passwd)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListFoldersWithPath(t *testing.T) {
	token := GenerateToken(req)
	token = `Basic aGFuY2hhbjphZG1pbkFkI0AyMDIx`
	var url string
	param.BASE_URL = "https://bptest.dengbao.cloud"
	param.AdminUser = "admin"
	param.AdminPasswd = "admin"
	url = `/remote.php/dav/files/hanchan/test/`
	ls, err := ListFoldersWithPath(token, url)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ls.DirList)
	for _, v := range ls.List {
		t.Logf("%s，%d，%s,%s \n", v.URL, v.FileType, v.Name, v.UsedBytes)
	}
}

func TestDownLoadFileURLWithPath(t *testing.T) {
	var uRL string
	// uRL = `/remote.php/dav/files/admin/golang.png`
	s, err := DownLoadFileURLWithPath(uRL)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s)
}

func TestDeleteFileWithPath(t *testing.T) {
	token := GenerateToken(req)
	param.BASE_URL = "https://bptest.dengbao.cloud"

	var uRL string
	// uRL = `/remote.php/dav/files/admin/golang.png`
	uRL = `/remote.php/dav/files/admin/测试文件夹/`
	err := DeleteFileWithPath(token, uRL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUploadFileWithPath(t *testing.T) {
	token := GenerateToken(req)
	var uRL string
	// uRL = `/remote.php/dav/files/admin/新建文件夹/`

	by, err := os.ReadFile("golang.png")
	if err != nil {
		t.Fatal(err)
	}
	err = UploadFileWithPath(token, "golang.png", bytes.NewBuffer(by), uRL)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStoreNCToken(t *testing.T) {
	err := model.StoreNCToken("hanchan", "456789")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(param.AdminUser)
	t.Log(param.AdminPasswd)
	t.Log(param.BASE_URL)
}

func TestBindNCTokenAndUID(t *testing.T) {
	err := model.BindNCTokenAndUID("hanchan", 123)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQueryTokenByUID(t *testing.T) {
	token, err := model.QueryTokenByUID(123)
	if err != nil {
		t.Fatal(err)
	}

	if token != "456789" {
		t.Fail()
	}
}

func TestCheckConf(t *testing.T) {
	err := CheckConf("test_hanchan", "21ops.com123", `https://bptest.dengbao.cloud`)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDirectDownloadURL(t *testing.T) {
	token := GenerateToken(req)
	param.BASE_URL = "http://localhost:8080"

	dURL, err := GetDirectDownloadURL(423, token)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dURL)
}

func TestCreateUserV2(t *testing.T) {
	token := GenerateToken(req)
	param.BASE_URL = "https://bptest.dengbao.cloud"

	err := CreateUserV2(token, "test_hanchan", "21ops.com@")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteNCUser(t *testing.T) {
	param.BASE_URL = "https://bptest.dengbao.cloud"
	param.AdminUser = "admin"
	param.AdminPasswd = "21ops.com"

	err := DeleteNCUser("test_hanchan")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetNCUserInfo(t *testing.T) {
	token := GenerateToken(req)
	param.BASE_URL = "http://localhost:8088"

	quota, used, percent := GetNCUserInfo(token, "hanchan")
	t.Log(quota)
	t.Log(used)
	t.Log(percent)
}

func TestCreateFoler(t *testing.T) {
	token := GenerateToken(req)
	param.BASE_URL = "https://bptest.dengbao.cloud"

	err := CreateFoler(token, "", "测试文件夹")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUserPassword(t *testing.T) {
	token := `Basic dGVzdF9oYW5jaGFuOjIxcG9zLmNvbUA=`
	param.BASE_URL = "https://bptest.dengbao.cloud"

	err := UpdateUserPassword("21pos.com.", token)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Change(t *testing.T) {
	//Basic dGVzdF9oYW5jaGFuOjIxb3BzLmNvbTEyMw==
	//go  get  github.com/go-resty/resty/v2
	var Client = resty.New().SetDebug(true).SetTimeout(time.Second * 60)
	Client = Client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	url := "https://bptest.dengbao.cloud/settings/personal/changepassword"
	resp, err := Client.R().
		//SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		//SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36").
		//SetHeader("method", "POST").
		//SetHeader("authority", "bptest.dengbao.cloud").
		//SetHeader("path", "/settings/personal/changepassword").
		//SetHeader("scheme", "https").
		//SetHeader("accept", "*/*").
		//SetHeader("accept-encoding", "gzip, deflate, br").
		//SetHeader("accept-language", "zh-CN,zh;q=0.9").
		//SetHeader("content-length", "72").
		//SetHeader("ocs-apirequest", "true").
		//SetHeader("origin", "https://bptest.dengbao.cloud").
		//SetHeader("sec-ch-ua", "\"Google Chrome\";v=\"93\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"93\"").
		//SetHeader("sec-ch-ua-mobile", "?0").
		//SetHeader("sec-ch-ua-platform", "\"macOS\"").
		//SetHeader("sec-fetch-dest", "empty").
		//SetHeader("sec-fetch-mode", "cors").
		//SetHeader("sec-fetch-site", "same-origin").
		SetHeader("x-requested-with", "XMLHttpRequest").
		SetHeader("Authorization", "Basic dGVzdF9oYW5jaGFuOjIxb3BzLmNvbTEyMw==").
		SetHeader("requesttoken", "qrjoZ9+Bpgp0l5PQz2UQ0T6e2LkLLNrwcO7YWEfwLhc=:7vW4EpvF4UUx4tyKgA1yl2bMsYl5fY+BOd+2FAGKG1U=").
		//SetHeader("cookie", "oc_sessionPassphrase=Xtey00gffPBNMWif%2Fh5uprJGuzFMHayQNNxhxSjkgjLs9tHj72hgdIlO1umypGNe8I9mgAtJ74%2BHFEgCj%2BRF1sh0QscCFBSruSZwyxOyyrDkcC7fFfPFxZOrjybjwDHl; __Host-nc_sameSiteCookielax=true; __Host-nc_sameSiteCookiestrict=true; ocdh9htx8nbo=5b8cc0c04e726b0640002d0eb0b55da7; nc_username=test_hanchan; nc_token=OGlKC394rb2BzhHugIYUl0OuDPhRDO25; nc_session_id=5b8cc0c04e726b0640002d0eb0b55da7").
		SetFormData(map[string]string{
			"oldpassword":       "21pos.com.",
			"newpassword":       "21ops.com123",
			"newpassword-clone": "21ops.com123",
		}).
		Post(url)
	if err != nil {
		fmt.Println("err1=", err)
	}
	if resp.StatusCode() == 200 {
		//获取cookie

		fmt.Println("resp=", string(resp.Body()))
	}
}

func TestDownLoadFileWithPath(t *testing.T) {
	token := `Basic aGFuY2hhbjphZG1pbkFkI0AyMDIx`
	param.BASE_URL = "https://bptest.dengbao.cloud"
	uRL := `/remote.php/dav/files/hanchan/456/下载.png`
	rsp, err := DownLoadFileWithPath(token, uRL)
	if err != nil {
		t.Fatal(err)
	}

	// bb, err := io.ReadAll(rsp.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	t.Log(rsp.Header.Get("Content-type"))
	rsp.Body.Close()
}

func TestHasSpecialChar(t *testing.T) {
	str := "12345`"
	if !hasSpecialChar(str) {
		t.Fail()
	}
}
