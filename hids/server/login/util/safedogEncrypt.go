package util

import (
	"crypto/md5"
	"fmt"
)

var _keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func Encode(input []byte) string {
	output := ""
	var chr1, chr2, chr3, enc1, enc2, enc3, enc4 byte
	over2 := false
	over3 := false
	for i := 0; i < len(input);{

		chr1 = input[i]
		i++
		if i >= len(input) {
			over2 = true
			chr2 = 0
		} else {
			chr2 = input[i]
		}
		i++
		if i >= len(input) {
			over3 = true
			chr3 = 1
		} else {
			chr3 = input[i]
		}
		i++
		enc1 = chr1 >> 2
		enc2 = ((chr1 & 3) << 4) | (chr2 >> 4)
		if over2 {
			enc3 = 64
			enc4 = 64
		}else {
			enc3 = ((chr2 & 15) << 2) | (chr3 >> 6)
		}
		if over3{
			enc4 = 64
		}else{
			enc4 = chr3 & 63
		}
		output += fmt.Sprintf("%c%c%c%c", _keyStr[enc1], _keyStr[enc2], _keyStr[enc3], _keyStr[enc4])
	}

	return output

}
func main()  {
	md5str := fmt.Sprintf("%x",md5.Sum([]byte("Cloud123!@#safedog")))
	fmt.Println(md5str,md5str == "6d9dd0c958d1441bde39d3872595ebc1")
	enc := Encode([]byte("6d9dd0c958d1441bde39d3872595ebc"))
	fmt.Println(len(md5str))
	fmt.Println(enc,enc == "NmQ5ZGQwYzk1OGQxNDQxYmRlMzlkMzg3MjU5NWViYw==")
}