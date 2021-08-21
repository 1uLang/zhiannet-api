package _const

var (
	BASE_URL = ""
	TEA_KEY  = ""
)

const (
	// Method: GET
	AGENT_LIST  = "api/v1/agents"
	AGENT_STATE = "api/v1/agent/info/%s"
	DELETE_AGENT = "api/v1/agent/%s/delete"
	CPU_USAGE   = "api/v1/agent/%s/app/system/item/cpu.usage/latest"
	MEM_USAGE   = "api/v1/agent/%s/app/system/item/memory.usage/latest"
	DISK_USAGE  = "api/v1/agent/%s/app/system/item/disk.usage/latest"

	// Method: Post
	ADD_AGENT = "api/v1/agent/add"
	UPDATE_AGENT = "api/v1/agent/update"
)
