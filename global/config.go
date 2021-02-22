package global

import (
	"fmt"
	"os"

	"github.com/vcqr/goini"
)

//ListenPort is webserver port
var ListenPort int64

//DBConn is DB DSN connection string
var DBConn string

//DBDriver is name of DB Driver (current: mysql only)
var DBDriver string

//SysSalt is password hash salt
var SysSalt string

//UserSalt is password hash salt
var UserSalt string

//Loadconf : Load configuation file
func Loadconf() {
	config := goini.Load(ConfigINIPath)
	confver := config.GetString("Ver", "app")

	if len(confver) < 2 {

		os.Exit(1)
	}
	ListenPort = config.GetInt("Port", "server", 8808)

	//
	DBDriver = config.GetString("Driver", "database", "mysql")

	host := config.GetString("Host", "database", "localhost")
	user := config.GetString("User", "database", "root")
	pass := config.GetString("Password", "database", "")
	name := config.GetString("Name", "database", "erp")
	port := config.GetInt("Port", "database", 3306)
	if port < 1024 || port > 65530 {
		port = 3306
	}
	DBConn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", user, pass, host, port, name)

}

//LoadSalt : system salt
func LoadSalt() {

}
