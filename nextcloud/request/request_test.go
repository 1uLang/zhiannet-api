package request

import (
	"embed"
	"io"
	"testing"

	"github.com/1uLang/zhiannet-api/nextcloud/model"
)

var (
	req = &model.LoginReq{
		User:     "admin",
		Password: "admin",
	}
	fileName = "Nextcloud.png"
	//go:embed Nextcloud.png
	nexcloud string
	//go:embed golang.png
	//go:embed test.jpg
	uf embed.FS
)

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

func TestUploadFile(t *testing.T) {
	token := GenerateToken(req)
	f, err := uf.Open("golang.png")
	if err != nil {
		t.Fatal(err)
	}

	err = UploadFile(token, "golang.png", f)
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
	userNamer := "hanchan"
	passwd := "123456"

	err := CreateUser(token, userNamer, passwd)
	if err != nil {
		t.Fatal(err)
	}
}
