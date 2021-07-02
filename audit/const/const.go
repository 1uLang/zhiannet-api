package _const

const (
	USERNAME     = "cdadmin"
	PASSWORD     = "A16pBIzVJOwHSC#Q"
	FAILURE_INFO = "Invalid user or password mismatch, or automatically logged out"

	AUDIT_HOST                  = "https://182.131.30.171:28443"
	AUDIT_LOGIN_URL             = "/sysLogin/login"                 //登陆
	DDOS_STATUS_GLOBAL_URL      = "/cgi-bin/status_global.cgi"      //全局统计
	DDOS_STATUS_HEALTH_URL      = "/cgi-bin/status_health.cgi"      //负载
	DDOS_HOST_STATUS_URL        = "/cgi-bin/status_host.cgi"        //主机状体 主机列表
	DDOS_STATUS_FBLINK_URL      = "/cgi-bin/status_fblink.cgi"      //屏蔽列表
	DDOS_STATUS_LINK_URL        = "/cgi-bin/status_link.cgi"        //链接列表
	DDOS_STATUS_HOSTSET_URL     = "/cgi-bin/status_hostset.cgi"     //主机设置
	DDOS_STATUS_BWLIST_URL      = "/cgi-bin/status_bwlist.cgi"      //黑白名单
	DDOS_LOGS_REPORT_ATTACK_URL = "/cgi-bin/logs_report_attack.cgi" //攻击日志列表
	DDOS_LOGS_REPORT_FLOW_URL   = "/cgi-bin/logs_report_flow.cgi"   //流量日志列表
	DDOS_LOGS_REPORT_LINK_URL   = "/cgi-bin/logs_report_link.cgi"   //连接日志列表
)
