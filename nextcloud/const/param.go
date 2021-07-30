package _const

const (
	BASE_URL = `http://123.129.208.232`
	// %s：用户名 method：PROPFIND
	LIST_FOLDERS = `remote.php/dav/files/%s`
	// %s: 用户名  %s：文件路径 method：GET
	DOWNLOAD_FILES = `remote.php/dav/files/%s/%s`
	// %s: 用户名 method: PUT
	UPLOAD_FILES = `remote.php/dav/files/%s/`

	// 创建用户 method: POST
	CREATE_USER = `ocs/v1.php/cloud/users`
	// 删除用户 method: DELETE
	DELETE_USER = `ocs/v1.php/cloud/users/%s`
	// 用户信息 method: GET
	USER_INFO = `ocs/v1.php/cloud/users/%s`
)
