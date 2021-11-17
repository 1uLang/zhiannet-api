package _const

const (
	AUDIT_HOST         = "https://182.131.30.171:28443"
	AUDIT_LOGIN_URL    = "/sysLogin/login"                 //登陆
	AUDIT_RESOURCE_URL = "/system/monitor/server/resource" //全局统计

	//数据库
	AUDIT_DB_LIST         = "/system/audit/dbList"       //数据库列表
	AUDIT_ADD_DB          = "/system/audit/addDB"        //添加
	AUDIT_EDIT_DB         = "/system/audit/editNameDB"   //修改
	AUDIT_DEL_DB          = "/system/audit/deleteDB"     //链接列表
	AUDIT_AUTH_EMAIL      = "/system/audit/authUser"     //添加授权
	AUDIT_AUTH_EMAIL_LIST = "/system/audit/authUserList" //授权列表
	AUDIT_DB_LOG_LIST     = "/system/audit/logList"      //日志列表

	//主机
	AUDIT_HOST_LIST     = "/system/audit/hostList"     //主机列表
	AUDIT_ADD_HOST      = "/system/audit/addHost"      //添加
	AUDIT_EDIT_HOST     = "/system/audit/editNameHost" //修改
	AUDIT_DEL_HOST      = "/system/audit/deleteHost"   //链接列表
	AUDIT_HOST_LOG_LIST = "/system/audit/hostLogList"  //日志列表

	//应用
	AUDIT_APP_LIST     = "/system/audit/appList"     //主机列表
	AUDIT_ADD_APP      = "/system/audit/addApp"      //添加
	AUDIT_EDIT_APP     = "/system/audit/editNameApp" //修改
	AUDIT_DEL_APP      = "/system/audit/deleteApp"   //链接列表
	AUDIT_APP_LOG_LIST = "/system/audit/appLogList"  //日志列表

	//报表
	AUDIT_FROM_LIST = "/system/audit/fromList"    //主机列表
	AUDIT_ADD_FROM  = "/system/audit/addFrom"     //添加
	AUDIT_EDIT_FROM = "/system/audit/editFrom"    //修改
	AUDIT_DEL_FROM  = "/system/audit/deleteFrom"  //删除
	AUDIT_GET_FROM  = "/system/audit/getFromInfo" //获取详情

	//用户
	AUDIT_ADD_USER = "/system/auth/addUser" //添加用户

	AUDIT_EMAIL_INFO  = "/system/auth/getCustomerInfo"   //邮件配置详情信息
	AUDIT_EMAIL_EDIT  = "/system/auth/editCustomerSMTP"  //保存邮件配置
	AUDIT_EMAIL_CHECK = "/system/auth/checkCustomerSMTP" //检测邮件配置

)
