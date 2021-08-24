package request

import (
	"encoding/json"
	"errors"
	"fmt"

	param "github.com/1uLang/zhiannet-api/resmon/const"
	"github.com/1uLang/zhiannet-api/resmon/model"
	"github.com/1uLang/zhiannet-api/resmon/server"
)

// AgentList agent列表
func AgentList() (*model.AgentList, error) {
	// 获取teaweb节点信息
	server.GetNodeInfo()
	if param.BASE_URL == "" {
		return nil, errors.New("该节点暂未添加，请添加后重试")
	}

	al := &model.AgentList{}
	body, err := createReq(param.AGENT_LIST, "")
	if err != nil {
		return nil, err
	}

	agents := model.Agents{}
	err = json.Unmarshal(body, &agents)
	if err != nil {
		return nil, fmt.Errorf("json解析失败：%w", err)
	}

	al.Total = len(agents)
	infos := make([]model.AgentInfo, al.Total)
	for i, v := range agents {
		info := model.AgentInfo{}
		info.Id = v.Config.ID
		info.Key = v.Config.Key
		info.Name = v.Config.Name
		info.Host = v.Config.Host
		info.On = v.Config.On

		as, _ := getAgentState(info.Id)
		info.OS = as.OSName
		info.Status = as.IsActive
		info.Cpu, _ = getCPUUsgae(info.Id)
		info.Mem, _ = getMemUsage(info.Id)
		info.Disk, _ = getDiskUsage(info.Id)
		info.OsType = server.GetResmon(info.Id)

		if !info.Status {
			info.OS = "-"
			info.Cpu = "-"
			info.Mem = "-"
			info.Disk = "-"
		}

		infos[i] = info
	}
	al.List = infos

	return al, nil
}
