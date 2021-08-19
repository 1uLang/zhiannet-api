package model

// FileListRsp 文件列表响应体
type FileListRsp struct {
	Total int64      `json:"total"`
	List  []FileInfo `json:"list"`
}

// FileInfo 文件信息
type FileInfo struct {
	ID        int64  `json:"id"`         // 主键ID
	Name      string `json:"name"`       // 文件名
	Describe  string `json:"describe"`   // 文件描述信息
	Size      string `json:"size"`       // 文件大小
	Format    string `json:"format"`     // 文件格式
	CreatedAt string `json:"created_at"` // 文件上传时间
}
