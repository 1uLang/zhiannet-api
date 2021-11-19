package access_gateway

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
)

var (
	access_gateway_path = "/access-gateways"
)

//List 用户网关列表
func List(req *request.Request, args *ListReq) ([]ListRes, int64, error) {
	var resp []ListRes
	list, total, err := getList(args)
	for _, v := range list {
		info, err := GetInfo(req, v.GatewayId)
		if err != nil {
			total--
			continue
		}
		info.AuthUser, _ = getUserNum(v.GatewayId)
		info.Auth = v.Auth == 0
		resp = append(resp, ListRes{*info})
	}
	return resp, total, err
}

func GetAll(req *request.Request, args *GetAllReq) ([]GetAllRes, error) {
	var resp []GetAllRes
	list, err := getAll(args)
	for _, v := range list {
		info, err := GetInfo(req, v.GatewayId)
		if err != nil || info == nil {
			continue
		}
		resp = append(resp, GetAllRes{info.ID, info.Name})
	}
	return resp, err
}

// Create 创建网关
func Create(req *request.Request, args *CreateReq) error {
	if err := args.check(); err != nil {
		return err
	}
	req.Method = "POST"
	req.Path = access_gateway_path
	req.Params = model.ToMap(args)

	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	if resp.Data == nil {
		return nil
	}
	err = addAccessGateway(&nextTerminalAccessGateway{GatewayId: resp.Data.(string),
		Auth:        0,
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
	})
	if err != nil {
		//todo : 删除当前创建的 网关
		_ = Delete(req, resp.Data.(string))
	}
	return err
}

// 获取网关信息
func GetInfo(req *request.Request, id string) (*GatewayInfo, error) {
	info := &GatewayInfo{}
	req.Method = "GET"
	req.Path = access_gateway_path + "/" + id
	req.Params = nil

	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, fmt.Errorf("服务器异常：%v", resp.Message)
	}
	if resp.Data == nil {
		return nil, fmt.Errorf("该网关不存在")
	}
	bytes, _ := json.Marshal(resp.Data)

	err = json.Unmarshal(bytes, &info)
	return info, err
}

// Delete 删除网关
func Delete(req *request.Request, id string) error {

	if id == "" {
		return fmt.Errorf("网关id不能为空")
	}

	ok, err := checkAssetGateway(id)
	if err != nil {
		return err
	}
	if ok {
		return fmt.Errorf("当前网关存在关联的资产正在使用，请先修改相关资产的关联关系。")
	}
	req.Method = "DELETE"
	req.Path = access_gateway_path + "/" + id
	req.Params = nil

	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	return deleteAccessGateway(id)
}

// Update 更新网关信息
func Update(req *request.Request, args *UpdateReq) error {
	if err := args.check(); err != nil {
		return err
	}
	req.Method = "PUT"
	req.Path = access_gateway_path + "/" + args.Id
	req.Params = model.ToMap(args)

	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	return nil
}

// Reconnect 重连网关
func Reconnect(req *request.Request, id string) error {

	if id == "" {
		return fmt.Errorf("网关id不能为空")
	}
	req.Method = "POST"
	req.Path = access_gateway_path + "/" + id + "/reconnect"
	req.Params = nil

	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	if resp.Data == nil {
		return fmt.Errorf("该网关不存在")
	}
	return nil
}

// Authorize 授权资产
func Authorize(req *request.Request, args *AuthorizeReq) error {
	err := args.check()
	if err != nil {
		return err
	}
	return authAccessGateway(args)
}

// AuthorizeUserList 授权用户列表
func AuthorizeUserList(req *request.Request, id string) ([]uint64, error) {
	if id == "" {
		return nil, fmt.Errorf("网关id不能为空")
	}
	return authUserList(id)
}
