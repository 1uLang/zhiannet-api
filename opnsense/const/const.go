package _const

const (
	OPNSENSE_LOGS_URL       = "/api/diagnostics/log/core/suricata"       //请求日志列表
	OPNSENSE_CLEAR_LOGS_URL = "/api/diagnostics/log/core/suricata/clear" //清空日志
	OPNSENSE_IPS_LIST_URL   = "/api/ids/settings/searchinstalledrules"   //ips规则
	OPNSENSE_IPS_EDIT_URL   = "/api/ids/settings/toggleRule"             //ips规则 启动停止
	OPNSENSE_IPS_DEL_URL    = "/api/ids/settings/toggleRule/%v/drop"     //ips规则 删除
)
