package migrate

import (
	"fmt"
	"time"

	"v2ex/go-erp/dbutil"
	"v2ex/go-erp/global"
	"v2ex/go-erp/util"

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

//User 用户表
type User struct {
	ID        uint32 `gorm:"primary_key;type:INT UNSIGNED NOT NULL AUTO_INCREMENT"`
	Username  string `gorm:"type:varchar(30);not null;unique"`
	Password  string `gorm:"type:varchar(32);not null"`
	Name      string `gorm:"type:varchar(8);not null"`
	Status    uint8  `gorm:"type:tinyint unsigned not null;default:1"`
	DepartID  uint32 `gorm:"type:INT UNSIGNED NOT NULL;index;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

//DicDepart 用户部门词典表
type DicDepart struct {
	ID       uint32 `gorm:"primary_key;type:INT UNSIGNED NOT NULL AUTO_INCREMENT"`
	Name     string `gorm:"type:varchar(32);not null"`
	Total    uint16 `gorm:"type:MEDIUMINT UNSIGNED NOT NULL"`
	ParentID uint32 `gorm:"type:INT UNSIGNED NOT NULL;index"`
}

//Group 用户权限组表
type Group struct {
	ID        uint32 `gorm:"primary_key;type:INT UNSIGNED NOT NULL AUTO_INCREMENT"`
	GroupName string `gorm:"type:varchar(30);not null;unique"`
}

//CheckUser 确认系统用户相关表情况
func CheckUser() {
	log.Info("Checking User table")
	db := dbutil.GetDB()
	if !db.Migrator().HasTable(&User{}) {
		log.Info("NO User TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
		log.Info("User table inited")
	}
	if !db.Migrator().HasTable(&DicDepart{}) {
		log.Info("NO DepartDic TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DicDepart{})
		log.Info("DepartDic table inited")

		log.Info("Trying to generate DepartDic base data")
		var depObj DicDepart
		depObj = DicDepart{ID: 1, Name: "总部直属", Total: 0, ParentID: 1}
		_ = db.Create(&depObj)
		log.Info("Done: DepartDic base data")
	}
	if !db.Migrator().HasTable(&Group{}) {
		log.Info("NO Group TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Group{})
		log.Info("Group table inited")
	}
}
