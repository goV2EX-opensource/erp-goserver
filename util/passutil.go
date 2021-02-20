package util

import (
	"fmt"

	"v2ex/go-erp/global"
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
