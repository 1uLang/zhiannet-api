package logs

//攻击日志
import (
	"fmt"
	"github.com/1uLang/zhiannet-api/ddos/model/ddos_host_ip"
	"github.com/1uLang/zhiannet-api/ddos/request"
	"github.com/1uLang/zhiannet-api/ddos/request/logs"
	"github.com/1uLang/zhiannet-api/ddos/server"
	host_status_server "github.com/1uLang/zhiannet-api/ddos/server/host_status"
	"time"
)

type (
	AttackLogReq struct {
		NodeId     uint64    `json:"node_id"`
		Addr       string    `json:"addr"`
		StartTime  time.Time `json:"start_time"`
		EndTime    time.Time `json:"end_time"`
		AttackType string    `json:"attack_type"`
		Status     int       `json:"status"`
	}
	EditBWReq struct {
		NodeId uint64   `json:"node_id"`
		Addr   []string `json:"addr"`
		White  bool     `json:"white"`
	}
)

//攻击日志列表
func GetAttackLogList(req *AttackLogReq) (list *logs.LogsReportAttack, err error) {
	//获取节点信息
	var logReq *request.LoginReq
	logReq, err = server.GetLoginInfo(server.NodeReq{NodeId: req.NodeId})
	fmt.Println("logReq==", logReq)
	if err != nil {
		return
	}
	if req.AttackType == "" {
		req.AttackType = "0"
	}
	if req.StartTime.IsZero() {
		timeStr := time.Now().Add(-time.Hour * 48).Format("2006-01-02")
		req.StartTime, _ = time.Parse("2006-01-02", timeStr)
	}
	if req.EndTime.IsZero() {
		timeStr := time.Now().Add(time.Hour * 24).Format("2006-01-02")
		req.EndTime, _ = time.Parse("2006-01-02", timeStr)
		req.EndTime = req.EndTime.Add(-time.Second)
	}
	nodes, _, err := host_status_server.GetHostList(&ddos_host_ip.HostReq{NodeId: req.NodeId, PageNum: 1, PageSize: 9999})
	if err != nil {
		return nil, err
	}
	list = &logs.LogsReportAttack{}

	args := &logs.AttackLogReq{
		Addr:       req.Addr,
		AttackType: req.AttackType,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
	}
	nodeSet := make(map[string]bool, 0)

	for _, v := range nodes {
		nodeSet[v.Addr] = true
	}
	for addr := range nodeSet {
		args.Addr = addr
		ls, err := logs.AttackLogList(args, logReq, true)
		if err != nil {
			return nil, err
		}
		list.Report = append(list.Report, ls.Report...)
	}

	return

}
