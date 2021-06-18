package request

type (
	ApiKey struct {
		Username   string
		Password   string
		Addr       string
		Port       string
		Cookie     string
		XCsrfToken string
	}
)
