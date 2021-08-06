package model

// ListFoldersResp 列举用户文件的返回
type ListFoldersResp struct {
	Response []struct {
		Href     string `xml:"href"`
		Propstat struct {
			Prop struct {
				Getlastmodified  string `xml:"getlastmodified"`
				QuotaUsedBytes   string `xml:"quota-used-bytes"`
				Getcontentlength string `xml:"getcontentlength"`
				Getcontenttype   string `xml:"getcontenttype"`
				FileID           uint64 `xml:"fileid"`
			} `xml:"prop"`
		} `xml:"propstat"`
	} `xml:"response"`
}

// CreateUserResp 创建用户返回
type CreateUserResp struct {
	Meta struct {
		Status     string `xml:"status"` // ok failure
		Statuscode int    `xml:"statuscode"`
		Message    string `xml:"message"`
	} `xml:"meta"`
}

// DeleteFileError 删除文件错误
type DeleteFileError struct {
	Message string `xml:"message"`
}

// FolderBody 文件实体属性
type FolderBody struct {
	FileID       uint64 `json:"file_id"`
	URL          string `json:"url"`
	Name         string `json:"name"`
	LastModified string `json:"last_modified"`
	UsedBytes    string `json:"used_bytes"`
	ContentType  string `json:"content_type"`
}

// FolderList 文件列表
type FolderList struct {
	List []FolderBody `json:"list"`
}
