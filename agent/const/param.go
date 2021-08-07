package _const

var (
	// MAX_FILE_SIZE 允许上传的最大文件尺寸 默认100M
	MAX_FILE_SIZE = 100 * 1024 * 1024
	// Describe_SIZE 描述信息字数限制，默认30字
	Describe_SIZE = 30
	// FILE_STORE_PATH 文件默认保存的绝对路径
	FILE_STORE_PATH = `/var/dengbaoyun/agent/files`
)

const (
	// B byte
	B float64 = 1
	// KB k byte
	KB float64 = 1024
	// MB M byte
	MB float64 = 1024 * 1024
	// GB G byte
	GB float64 = 1024 * 1024 * 1024

	// CONFIG_PATH 配置文件路径
	// CONFIG_PATH = `./build/configs/api_db.yaml`
)
