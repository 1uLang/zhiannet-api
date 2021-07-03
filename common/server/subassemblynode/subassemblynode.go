package subassemblynode

import "github.com/1uLang/zhiannet-api/common/model/subassemblynode"

type (
	NodeNumReq struct {
		Type  int    `json:"type"`
		State string `json:"state"`
	}
)

//所有节点列表
func GetNodeList() (list []*subassemblynode.Subassemblynode, total int64, err error) {
	list, total, err = subassemblynode.GetList(&subassemblynode.NodeReq{})
	return
}

//获取节点数量
func GetNodeNum(req *NodeNumReq) (total int64, err error) {
	total, err = subassemblynode.GetNum(&subassemblynode.NodeReq{
		Type:  req.Type,
		State: req.State,
	})
	return
}

//添加节点
func Add(req *subassemblynode.Subassemblynode) (id uint64, err error) {
	return subassemblynode.Add(req)
}

//修改节点
func Edit(req *subassemblynode.Subassemblynode) (rows int64, err error) {
	return subassemblynode.Edit(req, req.Id)
}

//删除节点
func Del(req uint64) (err error) {
	return subassemblynode.DeleteByIds([]uint64{req})
}

//获取节点详细信息
func GetNodeInfo(req uint64) (info *subassemblynode.Subassemblynode, err error) {
	return subassemblynode.GetNodeInfoById(req)
}
