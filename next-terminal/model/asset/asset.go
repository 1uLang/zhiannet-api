package asset

import (
	"encoding/json"
	"fmt"
	"github.com/1uLang/zhiannet-api/common/util"
	"github.com/1uLang/zhiannet-api/next-terminal/model"
	"github.com/1uLang/zhiannet-api/next-terminal/model/access_gateway"
	"github.com/1uLang/zhiannet-api/next-terminal/request"
	"strings"
	"time"
)

var (
	asset_path = "/assets"
	//authorize_path = "resource-sharers/overwrite-sharers"
)

//资产列表
func List(req *request.Request, args *ListReq) ([]interface{}, uint64, error) {
	if err := args.check(); err != nil {
		return nil, 0, err
	}
	req.Method = "GET"
	req.Path = asset_path + "/paging"
	req.Params = model.ToMap(args)

	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, 0, err
	}
	if resp.Code != 1 {
		return nil, 0, fmt.Errorf("服务器异常：%v", resp.Message)
	}
	if resp.Data == nil {
		return nil, 0, nil
	}
	retData := resp.Data.(map[string]interface{})
	total, _ := util.Interface2Uint64(retData["total"])
	return retData["items"].([]interface{}), total, nil
}

//创建资产
func Create(req *request.Request, args *CreateReq) error {
	if err := args.check(); err != nil {
		return err
	}
	req.Path = asset_path
	req.Method = "POST"
	req.Params = model.ToMap(args)
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	asset := resp.Data.(map[string]interface{})

	//写入数据库
	err = addAsset(&nextTerminalAssets{
		AssetsId:    asset["id"].(string),
		Name:        asset["name"].(string),
		Proto:       asset["protocol"].(string),
		UserId:      args.UserId,
		AdminUserId: args.AdminUserId,
		IsDelete:    0,
		CreateTime:  time.Now().Unix(),
	})
	if err != nil {
		_ = Delete(req, &DeleteReq{Id: asset["id"].(string)})
		return err
	}
	go access_gateway.SyncAssetGateway(asset["id"].(string), args.AccessGatewayId)
	return nil
}

//修改资产 默认同步到数据库
func Update(req *request.Request, args *UpdateReq, syncDB ...bool) error {
	if err := args.check(); err != nil {
		return err
	}

	//获取资产信息
	if len(syncDB) == 0 || syncDB[0] {
		asset, err := Details(req, &DetailsReq{args.Id})
		if err != nil {
			return err
		}
		args.Tags = asset["tags"].(string)
	}
	req.Path = asset_path + "/" + args.Id
	req.Method = "PUT"
	req.Params = model.ToMap(args)

	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	if len(syncDB) == 0 || syncDB[0] {
		//更新到数据库
		return updateAsset(&nextTerminalAssets{
			AssetsId: args.Id,
			Name:     args.Name,
			Proto:    args.Protocol,
			UserId:   args.UserId,
		})
	} else {
		return nil
	}

}

//删除资产
func Delete(req *request.Request, args *DeleteReq) error {

	req.Path = asset_path + "/" + args.Id
	req.Method = "DELETE"
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return err
	}
	if resp.Code != 1 {
		return fmt.Errorf("服务器异常：%v", resp.Message)
	}
	//删除数据库相关信息
	err = deleteAsset(args.Id)
	if err != nil {
		return err
	}
	go access_gateway.DeleteAsset(args.Id)
	return nil
}

//资产详情
func Details(req *request.Request, args *DetailsReq) (map[string]interface{}, error) {
	req.Path = asset_path + "/" + args.Id
	req.Method = "GET"
	req.Params = nil
	resp, err := req.DoAndParseResp()
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		return nil, fmt.Errorf("服务器异常：%v", resp.Message)
	}
	return resp.Data.(map[string]interface{}), nil
}

//授权资产
func Authorize(req *request.Request, args *AuthorizeReq) error {
	if args.UserId == 0 && args.AdminUserId == 0 {
		return fmt.Errorf("参数错误")
	}
	//获取资产信息
	asset, err := Details(req, &DetailsReq{args.AssetId})
	if err != nil {
		return err
	}
	//修改资产标签
	arg := &UpdateReq{Id: args.AssetId}
	arg.AccountType = asset["accountType"].(string)
	arg.Description = asset["description"].(string)
	arg.IP = asset["ip"].(string)
	arg.Protocol = asset["protocol"].(string)
	arg.Name = asset["name"].(string)
	arg.Password = asset["password"].(string)
	arg.Port, _ = util.Interface2Int(asset["port"])
	arg.SshMode = asset["accountType"].(string)
	arg.CredentialId = asset["credentialId"].(string)
	//加上授权用户

	tags := fmt.Sprintf("user_%v", args.UserId)
	if args.AdminUserId != 0 {
		tags = fmt.Sprintf("admin_%v", args.AdminUserId)
	}
	_ = resetAuthorize(args.AssetId)
	if len(args.UserIds) != 0 {
		for _, userId := range args.UserIds {
			tags += fmt.Sprintf(",user_%v", userId)
			err = addAsset(&nextTerminalAssets{
				AssetsId:    args.AssetId,
				UserId:      userId,
				AdminUserId: 0,
				IsDelete:    0,
				Auth:        1,
				CreateTime:  time.Now().Unix(),
			})
			if err != nil {
				return err
			}
		}
	} else {
		for _, userId := range args.AdminUserIds {
			tags += fmt.Sprintf(",admin_%v", userId)
			err = addAsset(&nextTerminalAssets{
				AssetsId:    args.AssetId,
				UserId:      0,
				AdminUserId: userId,
				IsDelete:    0,
				Auth:        1,
				CreateTime:  time.Now().Unix(),
			})
			if err != nil {
				return err
			}
		}
	}
	arg.Tags = tags
	arg.Username = asset["username"].(string)

	return Update(req, arg, false)
}

//授权用户列表
func AuthorizeUserList(req *request.Request, args *AuthorizeUserListReq) ([]string, error) {
	//获取资产信息
	asset, err := Details(req, &DetailsReq{args.AssetId})
	if err != nil {
		return nil, err
	}

	tags := asset["tags"].(string)
	tags = strings.ReplaceAll(tags, "user_", "")
	tags = strings.ReplaceAll(tags, "admin_", "")

	return strings.Split(tags, ","), nil
}

//连接资产
func Connect(req *request.Request, args *ConnectReq) (string, error) {
	req.Path = asset_path + "/" + args.Id + "/tcping"
	req.Method = "POST"
	req.Params = nil
	resp, err := req.Do()
	if err != nil {
		return "", err
	}
	retInfo := map[string]interface{}{}
	err = json.Unmarshal(resp, &retInfo)
	if err != nil {
		return "", err
	}
	fmt.Println(retInfo)
	if retInfo["code"].(float64) != 1 || retInfo["message"].(string) != "success" {
		return "", fmt.Errorf("连接资产失败：%v", retInfo["message"])
	}
	//获取本地数据库保存资产的名称跟协议
	info, err := getInfo(args.Id)
	if err != nil {
		return "", err
	}
	return "/#/access?assetId=" + args.Id + "&assetName=" + info.Name + "&protocol=" + info.Proto, nil
}
