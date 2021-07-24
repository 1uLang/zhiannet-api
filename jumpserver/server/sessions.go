package server

import (
	sessions_model "github.com/1uLang/zhiannet-api/jumpserver/model/sessions"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type sessions struct{ req *request.Request }

func (this *sessions) List(args *sessions_model.ListReq) ([]map[string]interface{}, error) {
	return sessions_model.List(this.req, args)
}
func (this *sessions) Monitor(id string) error {
	return sessions_model.Monitor(this.req, id)
}
func (this *sessions) Replay(id string) error {
	return sessions_model.Replay(this.req, id)
}
func (this *sessions) Download(id string) error {
	return sessions_model.Download(this.req, id)
}
