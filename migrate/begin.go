package migrate

//Begin : 初始化数据库
func Begin() {
	CheckSalt()
	CheckSysAdmin()
	CheckUser()
	CheckATD()
}

//Upgrade : 升级PATCH数据库
func Upgrade() {

}
