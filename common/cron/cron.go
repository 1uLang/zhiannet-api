package cron

import (
	audit_request "github.com/1uLang/zhiannet-api/audit/request"
	awvs_request "github.com/1uLang/zhiannet-api/awvs/server"
	"github.com/1uLang/zhiannet-api/common/server/attack_check_server"
	ddos_request "github.com/1uLang/zhiannet-api/ddos/request"
	hids_request "github.com/1uLang/zhiannet-api/hids/server"
	monitor_cron "github.com/1uLang/zhiannet-api/monitor/cron"
	nessus_request "github.com/1uLang/zhiannet-api/nessus/server"
	term_request "github.com/1uLang/zhiannet-api/next-terminal/server"
	nextcloud_request "github.com/1uLang/zhiannet-api/nextcloud/request"
	opnsense_request "github.com/1uLang/zhiannet-api/opnsense/request"
	teaweb_request "github.com/1uLang/zhiannet-api/resmon/request"
	"github.com/robfig/cron/v3"
)

func InitCron() {
	parser := cron.NewParser(
		cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	c := cron.New(cron.WithParser(parser))

	//c.AddFunc("*/1 * * * * *", func() {
	//	fmt.Println("tick every 1 second")
	//	time.Sleep(time.Second * 5)
	//})
	//运行结构体中的Run 方法,每次任务之间不重复 每5分钟执行一次
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&monitor_cron.PortCheck{}))
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&monitor_cron.CodeCheck{}))
	//安全审计组件定时检测状态是否可用
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&audit_request.LoginReq{}))
	//ddos金盾 组件定时检测状态是否可用
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&ddos_request.LoginReq{}))
	//下一代防火墙 组件定时检测状态是否可用
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&opnsense_request.ApiKey{}))
	//数据备份系统 组件定时检测状态是否可用
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&nextcloud_request.CheckRequest{}))
	//web漏洞扫描
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&awvs_request.CheckRequest{}))
	//主机漏洞扫描
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&nessus_request.CheckRequest{}))
	//主机防护
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&hids_request.CheckRequest{}))
	//堡垒机
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&term_request.CheckRequest{}))
	//tea web 节点监控
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&teaweb_request.CheckRequest{}))

	//入侵事件检测
	c.AddJob("0 */10 * * * *", cron.NewChain(cron.DelayIfStillRunning(cron.DefaultLogger)).Then(&attack_check_server.AttackCheckRequest{}))
	c.Start()
}
