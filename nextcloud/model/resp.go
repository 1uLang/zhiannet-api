package model

// ListFoldersResp 列举用户文件的返回
type ListFoldersResp struct {
	Response []struct {
		Href struct {
			Text string `xml:",chardata"`
		} `xml:"href"`
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

// FolderBody 文件实体属性
type FolderBody struct {
	URL  string
	Name string
}

// FolderList 文件列表
type FolderList struct {
	List []FolderBody
}
