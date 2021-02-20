package migrate

import (
	"fmt"
	"time"

	"example.com/go-v2ex/dbutil"
	"example.com/go-v2ex/global"
	"example.com/go-v2ex/util"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//SysUser : 系统管理员数据库结构
type SysUser struct {
	ID        uint16 `gorm:"primary_key;type:MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT"`
	Uname     string `gorm:"type:varchar(30);not null;unique"`
	Upass     string `gorm:"type:varchar(32);not null"`
	Ustat     uint8  `gorm:"type:tinyint unsigned not null;default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//CheckSysAdmin 确认系统初始管理员表是否存在，如果不存在新建并添加一个管理员
func CheckSysAdmin() {
	db := dbutil.GetDB()
	if !db.Migrator().HasTable(&SysUser{}) {
		log.Info("NO SYSADMIN TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&SysUser{})
		log.Info("SYSADMIN table inited")

		newstr := fmt.Sprintf("%s-!-%s", "root", global.SysSalt)
		newpass := util.Md5s(newstr)

		sysUserObj := SysUser{Uname: "sysadmin", Upass: newpass, Ustat: 1}
		result := db.Create(&sysUserObj)
		if result.RowsAffected != 1 {
			log.Fatal("INIT ADMIN ERROR")
		} else {
			log.Info("First sys admin generated!")
		}
	}
}
