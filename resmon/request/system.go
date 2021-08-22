package request

import (
	"encoding/json"
	"fmt"

	param "github.com/1uLang/zhiannet-api/resmon/const"
	"github.com/1uLang/zhiannet-api/resmon/model"
)

func getCPUUsgae(aid string) (string, error) {
	body, err := createReq(param.CPU_USAGE, aid)
	if err != nil {
		return "-", err
	}

	cpuUsage := model.CPUUsage{}
	err = json.Unmarshal(body, &cpuUsage)
	if err != nil {
		return "-", fmt.Errorf("json解析失败：%w", err)
	}

	return fmt.Sprintf("%.2f", cpuUsage.Usage.Avg) + "%", nil
}

func getMemUsage(aid string) (string, error) {
	body, err := createReq(param.MEM_USAGE, aid)
	if err != nil {
		return "-", err
	}

	memUsage := model.MemUsage{}
	err = json.Unmarshal(body, &memUsage)
	if err != nil {
		return "-", fmt.Errorf("json解析失败：%w", err)
	}

	return fmt.Sprintf("%.2fGB/%.2fGB", memUsage.Usage.VirtualUsed, memUsage.Usage.VirtualTotal), nil
}

func getDiskUsage(aid string) (string, error) {
	body, err := createReq(param.DISK_USAGE, aid)
	if err != nil {
		return "-", err
	}

	diskUsage := model.DiskUsage{}
	err = json.Unmarshal(body, &diskUsage)
	if err != nil {
		return "-", fmt.Errorf("json解析失败：%w", err)
	}

	var used, total int64
	for _, du := range diskUsage.Partitions {
		if du.Name == "/" {
			used = du.Used
			total = du.Total
			break
		}
	}

	return fmt.Sprintf("%s/%s", formatByte(used), formatByte(total)), nil
}

// getAgentState 获取代理主机的状态信息
func getAgentState(aid string) (*model.AgentState, error) {
	as := &model.AgentState{}
	body, err := createReq(param.AGENT_STATE, aid)
	if err != nil {
		return as, err
	}

	err = json.Unmarshal(body, as)
	if err != nil {
		return as, fmt.Errorf("json解析失败：%w", err)
	}

	return as, nil
}
