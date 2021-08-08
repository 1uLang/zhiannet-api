package asset

import "fmt"

type (
	ListReq struct {
		PageIndex int    `json:"pageIndex"`
		PageSize  int    `json:"pageSize"`
		Tags      string `json:"tags"`

		UserId      int64 `json:"-"`
		AdminUserId int64 `json:"-"`
	}
	CreateReq struct {
		AccountType  string `json:"accountType"`
		Description  string `json:"description"`
		IP           string `json:"ip"`
		Name         string `json:"name"`
		Password     string `json:"password"`
		Port         int    `json:"port"`
		Protocol     string `json:"protocol"`
		SshMode      string `json:"ssh-mode"`
		Tags         string `json:"tags"`
		Username     string `json:"username"`
		CredentialId string `json:"credentialId,omitempty"`
		UserId       uint64 `json:"-"`
		AdminUserId  uint64 `json:"-"`
	}
	UpdateReq struct {
		Id string `json:"-"`
		CreateReq
	}
	DetailsReq struct {
		Id string
	}
	DeleteReq struct {
		Id string
	}
	AuthorizeReq struct {
		AssetId      string

		UserId       uint64
		AdminUserId  uint64
		UserIds      []uint64 `json:"-"`
		AdminUserIds []uint64 `json:"-"`
	}
	AuthorizeUserListReq struct {
		AssetId string
	}
	ConnectReq struct {
		Id string
	}
)

var protos = map[string]bool{
	"rdp":    true,
	"ssh":    true,
	"vnc":    true,
	"telnet": true,
}

func (this *ListReq) check() error {
	if this.UserId == 0 && this.AdminUserId == 0 {
		return fmt.Errorf("参数错误")
	}
	if this.Tags == "" {
		this.Tags = fmt.Sprintf("user_%v", this.UserId)
		if this.AdminUserId != 0 {
			this.Tags = fmt.Sprintf("admin_%v", this.AdminUserId)
		}
	}
	return nil
}
func (this *CreateReq) check() error {

	if this.Tags == "" {
		if this.UserId == 0 && this.AdminUserId == 0 {
			return fmt.Errorf("参数错误")
		}

		this.Tags = fmt.Sprintf("user_%v", this.UserId)
		if this.AdminUserId != 0 {
			this.Tags = fmt.Sprintf("admin_%v", this.AdminUserId)
		}
	}
	if this.AccountType == "custom"{
		this.CredentialId = ""
	}
	_, isExist := protos[this.Protocol]

	if !isExist {
		return fmt.Errorf("暂不支持该协议")
	}
	return nil
}
func (this *UpdateReq) check() error {
	if this.Id == ""{
		return fmt.Errorf("参数错误")
	}
	return this.CreateReq.check()
}
