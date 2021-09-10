package host

import (
	"encoding/json"
	"fmt"
	_const "github.com/1uLang/zhiannet-api/zstack/const"
	"github.com/1uLang/zhiannet-api/zstack/request"
)

//镜像信息

type (
	ImageListReq struct {
		UUid string `json:"uuid"`
	}

	ImageListResp struct {
		Inventories []ImageInventories `json:"inventories"`
	}
	ImageInventories struct {
		ActualSize        int64               `json:"actualSize"`
		BackupStorageRefs []BackupStorageRefs `json:"backupStorageRefs"`
		CreateDate        string              `json:"createDate"`
		Description       string              `json:"description"`
		Format            string              `json:"format"`
		LastOpDate        string              `json:"lastOpDate"`
		Md5Sum            string              `json:"md5Sum"`
		MediaType         string              `json:"mediaType"`
		Name              string              `json:"name"`
		Platform          string              `json:"platform"`
		Size              int64               `json:"size"`
		State             string              `json:"state"`
		Status            string              `json:"status"`
		System            bool                `json:"system"`
		Type              string              `json:"type"`
		URL               string              `json:"url"`
		UUID              string              `json:"uuid"`
	}

	BackupStorageRefs struct {
		BackupStorageUUID string `json:"backupStorageUuid"`
		CreateDate        string `json:"createDate"`
		ID                int64  `json:"id"`
		ImageUUID         string `json:"imageUuid"`
		InstallPath       string `json:"installPath"`
		LastOpDate        string `json:"lastOpDate"`
		Status            string `json:"status"`
	}
)

//获取云计算主机镜像
func ImageList(req *ImageListReq) (resp *ImageListResp, err error) {
	//获取数据
	logReq, err := request.GetLoginInfo(&request.UserReq{})
	if err != nil {
		return
	}
	logReq.Addr = fmt.Sprintf("%v%v", logReq.Addr, _const.ZSTACK_IMAGE)
	logReq.ReqType = "get"
	var res []byte
	res, err = request.Request(logReq, true)
	if err != nil {
		return
	}

	err = json.Unmarshal(res, &resp)
	return
}
