package _const

const (
	OPNSENSE_GLOBAL_STATUS_URL = "/widgets/api/get.php?load=system%2Cinterfaces&_=" //全局状态

	OPNSENSE_LOGS_URL             = "/api/diagnostics/log/core/suricata"           //请求日志列表
	OPNSENSE_CLEAR_LOGS_URL       = "/api/diagnostics/log/core/suricata/clear"     //清空日志
	OPNSENSE_IPS_LIST_URL         = "/api/ids/settings/searchinstalledrules"       //ips规则
	OPNSENSE_IPS_EDIT_URL         = "/api/ids/settings/toggleRule"                 //ips规则 启动停止
	OPNSENSE_IPS_DEL_URL          = "/api/ids/settings/toggleRule/%v/drop"         //ips规则 删除
	OPNSENSE_IPS_APPLY_URL        = "/api/ids/service/reloadRules"                 //ips规则 应用
	OPNSENSE_IPS_ACTIOB_URL       = "/api/ids/settings/setRule"                    //ips规则 操作修改
	OPNSENSE_IPS_ALARM_LIST_URL   = "/api/ids/service/queryAlerts"                 //ips-报警
	OPNSENSE_IPS_ALARM_TIME_URL   = "/api/ids/service/getAlertLogs"                //ips-下拉时间
	OPNSENSE_IPS_ALARM_IFACE_URL  = "/api/diagnostics/interface/getInterfaceNames" //ips-接口名称
	OPNSENSE_DIAGNOSTICS_LIST_URL = "/api/diagnostics/firewall/query_pf_top"       //会话列表
	OPNSENSE_IPS_RULE             = "/api/ids/settings/listRulesets"               //规则列表
	OPNSENSE_FIRMWARE             = "/api/core/firmware/info"                      //软件包

	OPNSENSE_FILTER_SEARCH_URL = "/api/firewall/filter/searchRule" //过滤规则 搜索
	OPNSENSE_FILTER_ENABLE_URL = "/api/firewall/filter/toggleRule" //过滤规则 启用停用
	OPNSENSE_FILTER_DEL_URL    = "/api/firewall/filter/delRule"    //过滤规则 启用停用
	OPNSENSE_FILTER_INFO_URL   = "/api/firewall/filter/getRule"    //过滤规则 详情
	OPNSENSE_FILTER_SET_URL    = "/api/firewall/filter/setRule"    //过滤规则 设置
	OPNSENSE_FILTER_ADD_URL    = "/api/firewall/filter/addRule"    //过滤规则 添加

	OPNSENSE_NAT_1TO1_LIST_URL   = "/firewall_nat_1to1.php"      //nat 1:1 列表
	OPNSENSE_NAT_1TO1_INFO_URL   = "/firewall_nat_1to1_edit.php" //nat 1:1 详情或修改
	OPNSENSE_NAT_1TO1_STATUS_URL = "/firewall_nat_1to1.php"      //nat 1:1 启动停止,应用修改

	OPNSENSE_ACL_LIST_URL    = "/firewall_rules.php"                   //acl 规则列表
	OPNSENSE_ACL_INFO_URL    = "/firewall_rules_edit.php"              //acl 规则详情
	OPNSENSE_CLAMAV_INFO_URL = "/api/clamav/service/version"           //病毒库版本
	OPNSENSE_CLAMAV_LOG_URL  = "/api/diagnostics/log/clamav/freshclam" //病毒库更新列表

)
