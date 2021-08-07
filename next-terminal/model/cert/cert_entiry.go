package cert

import "fmt"

type (
	ListReq struct {
		PageIndex   int    `json:"pageIndex"`
		PageSize    int    `json:"pageSize"`
		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
	CreateReq struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Type     string `json:"type"`
		Username string `json:"username"`

		UserId      uint64 `json:"-"`
		AdminUserId uint64 `json:"-"`
	}
	UpdateReq struct {
		ID string `json:"-"`
		CreateReq
	}
	DeleteReq struct {
		ID string
	}
	DetailsReq struct {
		ID string
	}
	AuthorizeReq struct {
		ID string

		UserIds      []uint64 `json:"-"`
		AdminUserIds []uint64 `json:"-"`
	}
	AuthorizeUserListReq struct {
		ID string
	}
	AuthorizeUserListResp struct {
		UserIds     []uint64
		AdminUserId []uint64
	}
)

func (this *CreateReq) check() error {

	if this.Username == "" || this.Password == "" || this.Name == "" {
		return fmt.Errorf("参数错误")
	}
	this.Type = "custom"
	return nil
}

func (this *UpdateReq) check() error {
	if this.ID == "" {
		return fmt.Errorf("参数错误")
	}
	return this.CreateReq.check()
}
