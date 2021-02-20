package migrate

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"example.com/go-v2ex/dbutil"
	"example.com/go-v2ex/global"
	"example.com/go-v2ex/util"
	"github.com/lifei6671/gorand"
)

//Salt : HASH SALT
type Salt struct {
	ID        uint8  `gorm:"type:TINYINT UNSIGNED NOT NULL AUTO_INCREMENT"`
	SysSalt   string `gorm:"type:char(32);not null"`
	UserSalt  string `gorm:"type:char(32);not null"`
	CreatedAt time.Time
}

//CheckSalt : 检查salt是否设置，未设置新设置
func CheckSalt() {
	db := dbutil.GetDB()
	if !db.Migrator().HasTable(&Salt{}) {

		log.Info("NO SALT TABLE. PREPARE INIT ONE.")
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Salt{})
		log.Info("SALT table inited")
	}

	var saltObj Salt
	if result := db.First(&saltObj); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Info("NO SALT RECORED, PREPARE INIT ONE")

			salt1 := gorand.RandomAlphabetic(128)
			salt2 := gorand.NewUUID4().String()

			salt1 = util.Md5s(salt1)
			salt2 = util.Md5s(salt2)

			saltObj := Salt{SysSalt: salt1, UserSalt: salt2}
			result := db.Create(&saltObj)
			if result.RowsAffected != 1 {
				log.Fatal("INIT SALT ERROR")
			} else {
				global.SysSalt = salt1
				global.UserSalt = salt2
				log.Info("SALT INITED AND STORAGED")
			}
		} else {
			log.Fatal("DB FATAL ERROR")
		}
	} else {
		global.SysSalt = saltObj.SysSalt
		global.UserSalt = saltObj.UserSalt
		log.Info("SALT READ FROM STORAGE ENGINE")
	}
}
