package subassemblynode

import "github.com/1uLang/zhiannet-api/common/model/subassemblynode"

//所有节点列表
func GetNodeList() (list []*subassemblynode.Subassemblynode, err error) {
	list, err = subassemblynode.GetList(&subassemblynode.NodeReq{State: "1"})
	return
}

//添加节点
func Add(req *subassemblynode.Subassemblynode) (id uint64, err error) {
	return subassemblynode.Add(req)
}
