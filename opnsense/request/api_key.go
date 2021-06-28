package request

type (
	ApiKey struct {
		Username   string
		Password   string
		Addr       string
		Port       string
		Cookie     string
		XCsrfToken string
		IsSsl      bool //是否使用ssl协议登陆
	}
)
