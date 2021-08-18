package request

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"testing"

	param "github.com/1uLang/zhiannet-api/nextcloud/const"
	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

var (
	req = &model.LoginReq{
		User: "admin",
		// User: "admin_zhoumj",
		// Password: "Dengbao123!@#",
		Password: "admin",
		// Password: "21ops.com",
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
	user, err := ParseToken(token)
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
	// param.BASE_URL = "http://localhost:8080"
	// param.AdminUser = "admin"
	// param.AdminPasswd = "admin"
	lf, err := ListFolders(token)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range lf.List {
		t.Logf("%d %s\n", v.FileID, v.URL)
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
	var url string
	// param.BASE_URL = "http://localhost:8080"
	// param.AdminUser = "admin"
	// param.AdminPasswd = "admin"
	// url = `/remote.php/dav/files/admin/新建文件夹/`
	ls, err := ListFoldersWithPath(token, url)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range ls.List {
		t.Logf("%s\n", v.URL)
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
	var uRL string
	// uRL = `/remote.php/dav/files/admin/golang.png`
	uRL = `/remote.php/dav/files/admin/新建文件夹/`
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
	err := CheckConf("admin", "21ops.com", `https://bptest.dengbao.cloud`)
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
	param.BASE_URL = "http://localhost:8088"

	err := CreateUserV2(token, "test123456", "Test@test12345")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteNCUser(t *testing.T) {
	param.BASE_URL = "http://localhost:8088"
	param.AdminUser = "admin"
	param.AdminPasswd = "admin"

	err := DeleteNCUser("hanchan")
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
