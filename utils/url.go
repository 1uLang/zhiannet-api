package utils

import (
	"fmt"
	"strings"
)

//按照协议 取对应地址
func CheckHttpUrl(url string, isSsl bool) string {
	url = strings.TrimLeft(url, "https://")
	url = strings.TrimLeft(url, "http://")
	if isSsl {
		url = fmt.Sprintf("https://%v", url)
	} else {
		url = fmt.Sprintf("http://%v", url)
	}
	return url
}
