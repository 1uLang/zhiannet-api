package _const

var (
	// BASE_URL = `https://123.129.208.232`
	// BASE_URL = `http://localhost:8080`
	// BASE_URL = `http://182.150.0.80:18002/backup`
	// AdminUser = `admin`
	// AdminPasswd = `Dengbao123!@#`
	BASE_URL = ``
	AdminUser = ``
	AdminPasswd = ``
)

const (
	DB_CONFIG_PATH = `./build/configs/api_db.yaml`
	// DB_CONFIG_PATH = `../docs/test.yaml`
	// %s：用户名 method：PROPFIND
	LIST_FOLDERS = `remote.php/dav/files/%s`
	// %s: 用户名  %s：文件路径 method：GET
	DOWNLOAD_FILES = `remote.php/dav/files/%s/%s`
	// %s: 用户名 method: PUT
	UPLOAD_FILES = `remote.php/dav/files/%s/`
	// method: POST
	Direct_Download = `ocs/v2.php/apps/dav/api/v1/direct`

	// 创建用户 method: POST
	CREATE_USER = `ocs/v1.php/cloud/users`
	CREATE_USER_V2 = `ocs/v2.php/cloud/users`
	// 删除用户 method: DELETE
	DELETE_USER = `ocs/v2.php/cloud/users/%s`
	// 用户信息 method: GET
	USER_INFO = `ocs/v1.php/cloud/users/%s`
)


