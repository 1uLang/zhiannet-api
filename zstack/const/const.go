package _const

var (
	ZSTACK_LOGIN               = "/zstack/v1/accounts/login"                         //登陆
	ZSTACK_HOST_LIST           = "/zstack/v1/vm-instances"                           //云主机列表
	ZSTACK_HOSTS               = "/zstack/v1/hosts"                                  //物理机列表
	ZSTACK_SUSPEND             = "/zstack/v1/vm-instances/%v/actions"                //操作云主机
	ZSTACK_GLOBAL              = "/zstack/v1/global-configurations/%v/%v/actions"    //修改全局参数
	ZSTACK_SPEC                = "/zstack/v1/instance-offerings"                     //计算规格
	ZSTACK_IMAGE               = "/zstack/v1/images"                                 //镜像列表
	ZSTACK_DISK                = "/zstack/v1/disk-offerings"                         //云盘列表
	ZSTACK_NETWORK             = "/zstack/v1/l3-networks"                            //3层网络列表
	ZSTACK_MIGRATION_CANDIDATE = "/zstack/v1/vm-instances/%v/migration-target-hosts" //迁移可选物理机
	ZSTACK_CANDIDATE_STORAGES  = "/zstack/v1/vm-instances/candidate-storages"        //可选择的主存储

)
