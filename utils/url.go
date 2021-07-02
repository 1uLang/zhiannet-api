package utils

import (
	"fmt"
	"strings"
)

//按照协议 取对应地址
func CheckHttpUrl(url string, isSsl bool) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	if isSsl {
		url = fmt.Sprintf("https://%v", url)
	} else {
		url = fmt.Sprintf("http://%v", url)
	}
	return url
}
