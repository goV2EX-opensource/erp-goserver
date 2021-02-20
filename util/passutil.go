package util

import (
	"fmt"

	"example.com/go-v2ex/global"
)

//SysPass : System administrator password hash
func SysPass(newpass string) string {
	var result string
	newPassStr := fmt.Sprintf("%s()%s", newpass, global.SysSalt)
	result = Md5s(newPassStr)
	return result

}

//UserPass : User password hash
func UserPass(newpass string) string {
	var result string
	newPassStr := fmt.Sprintf("%s()%s", newpass, global.UserSalt)
	result = Md5s(newPassStr)
	return result
}
