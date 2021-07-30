package request

import (
	"bytes"
	_ "embed"
	"io"
	"os"
	"testing"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

var (
	req = &model.LoginReq{
		User:     "admin",
		Password: "Dengbao123!@#",
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
	lf, err := ListFolders(token)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(lf)
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
