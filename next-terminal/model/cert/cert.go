package cert

import (
	"fmt"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
	"time"
)

var cert_path = "/credentials"

func List(req *request.Request, args *ListReq) ([]interface{}, int64, error) {

	req.Path = cert_path + "/paging"
	req.Method = "get"
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, 0, err
	}
	if resp.Code != 1 {
		return nil, 0, fmt.Errorf("服务器异常：%v", resp.Message)
	}
	//读数据库
	list, total, err := getList(&listReq{
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
	})
	if err != nil || total == 0 {
		return nil, 0, err
	}
	contain := map[string]bool{}
	for _, v := range list {
		contain[v.CertId] = v.Auth == 1
	}
	ret := make([]interface{}, 0)
	for _, v := range resp.Data.(map[string]interface{})["items"].([]interface{}) {
		item := v.(map[string]interface{})
		if auth, isExist := contain[item["id"].(string)]; isExist {
			//授权
			item["auth"] = auth
			item["count"], _ = countAuthCert(item["id"].(string))
			ret = append(ret, item)
		}
	}
	return ret, total, nil
}
func Create(req *request.Request, args *CreateReq) error {
	if err := args.check(); err != nil {
		return err
	}
	req.Path = cert_path
	req.Method = "post"
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	//写数据库
	return addCert(&nextTerminalCert{
		CertId:      resp.Data.(map[string]interface{})["id"].(string),
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		IsDelete:    0,
		Auth:        0,
		CreateTime:  time.Now().Unix(),
	})
}
func Update(req *request.Request, args *UpdateReq) error {

	if err := args.check(); err != nil {
		return err
	}
	req.Path = cert_path + "/" + args.ID
	req.Method = "put"
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
func Delete(req *request.Request, args *DeleteReq) error {

	if args.ID == "" {
		return fmt.Errorf("参数错误")
	}
	req.Path = cert_path + "/" + args.ID
	req.Method = "delete"
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	//删除数据库
	return deleteCert(args.ID)
}
func Details(req *request.Request, args *DetailsReq) (map[string]interface{}, error) {

	if args.ID == "" {
		return nil, fmt.Errorf("参数错误")
	}
	req.Path = cert_path + "/" + args.ID
	req.Method = "get"
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, fmt.Errorf("服务器异常：%v", resp.Message)
	}
	return resp.Data.(map[string]interface{}), nil
}
func Authorize(req *request.Request, args *AuthorizeReq) error {

	info, err := Details(req, &DetailsReq{
		ID: args.ID,
	})
	if err != nil {
		return err
	}
	//清空cert授权
	err = resetAuthorize(args.ID)
	if err != nil {
		return err
	}
	//写数据库
	for _, user := range args.UserIds {
		err = addCert(&nextTerminalCert{
			CertId:      info["id"].(string),
			UserId:      user,
			AdminUserId: 0,
			IsDelete:    0,
			Auth:        1,
			CreateTime:  time.Now().Unix(),
		})
		if err != nil {
			return err
		}
	}
	for _, user := range args.AdminUserIds {
		err = addCert(&nextTerminalCert{
			CertId:      info["id"].(string),
			UserId:      0,
			AdminUserId: user,
			IsDelete:    0,
			Auth:        1,
			CreateTime:  time.Now().Unix(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func AuthorizeUserList(args *AuthorizeUserListReq) (resp *AuthorizeUserListResp, err error) {

	list, err := listAuthorize(args.ID)
	if err != nil {
		return nil, err
	}
	resp = &AuthorizeUserListResp{}
	for _,v := range list {
		if v.UserId != 0 {
			resp.UserIds = append(resp.UserIds, v.UserId)
		}else {
			resp.AdminUserId = append(resp.AdminUserId, v.AdminUserId)
		}
	}
	return
}
