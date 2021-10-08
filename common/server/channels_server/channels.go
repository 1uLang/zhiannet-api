package channels_server

import (
	"github.com/1uLang/zhiannet-api/common/model/channels"
	"github.com/1uLang/zhiannet-api/common/model/edge_users"
)

//所有列表
func GetList(req *channels.ChannelReq) (list []*channels.Channels, total int64, err error) {
	list, total, err = channels.GetList(req)
	return
}

//添加
func Add(req *channels.Channels) (id uint64, err error) {
	if req.Id == 0 {
		return channels.Add(req)
	}
	return Edit(req)
}

//修改
func Edit(req *channels.Channels) (rows uint64, err error) {
	return channels.Edit(req, req.Id)
}

//删除节点
func Del(req uint64) (err error) {
	return channels.DeleteByIds([]uint64{req})
}

//获取节点详细信息
func GetInfo(req uint64) (info *channels.Channels, err error) {
	return channels.GetChannelById(req)
}

//
//通过用户 获取节点详细信息
func GetInfoByUid(uid uint64) (info *channels.Channels, err error) {
	user := &edge_users.EdgeUsers{}
	user, err = edge_users.GetInfoById(uid)
	if err != nil {
		return
	}
	return channels.GetChannelById(user.ChannelId)
}
