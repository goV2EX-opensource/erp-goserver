package util

import (
	"crypto/md5"
	"fmt"
)

//Md5s : MD5 字符串
func Md5s(str string) string {
	data := []byte(str)
	result := fmt.Sprintf("%x", md5.Sum(data))
	return result
}
