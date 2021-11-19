package server

import (
	gateway_model "github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

type gateway struct{ req *request.Request }

func (this *gateway) List(args *gateway_model.ListReq) ([]gateway_model.ListRes, int64, error) {
	return gateway_model.List(this.req, args)
}
func (this *gateway) Info(id string) (*gateway_model.GatewayInfo, error) {
	return gateway_model.GetInfo(this.req, id)
}
func (this *gateway) Create(args *gateway_model.CreateReq) error {
	return gateway_model.Create(this.req, args)
}
func (this *gateway) Update(args *gateway_model.UpdateReq) error {
	return gateway_model.Update(this.req, args)
}
func (this *gateway) Delete(id string) error {
	return gateway_model.Delete(this.req, id)
}
func (this *gateway) Authorize(args *gateway_model.AuthorizeReq) error {
	return gateway_model.Authorize(this.req, args)
}
func (this *gateway) ReConnect(id string) error {
	return gateway_model.Reconnect(this.req, id)
}
func (this *gateway) AuthorizeUserList(id string) ([]uint64, error) {
	return gateway_model.AuthorizeUserList(this.req, id)
}
