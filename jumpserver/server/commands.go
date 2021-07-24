package server

import (
	commands_model "github.com/1uLang/zhiannet-api/jumpserver/model/commands"
	"github.com/1uLang/zhiannet-api/jumpserver/request"
)

type commands struct{ req *request.Request }

func (this *commands) List(args *commands_model.ListReq) ([]map[string]interface{}, error) {
	return commands_model.List(this.req, args)
}
