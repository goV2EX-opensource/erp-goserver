package dbutil

import (
	log "github.com/sirupsen/logrus"

	"time"

	"example.com/go-v2ex/global"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

//NewConnection : 连接池建立新连接
func NewConnection() *gorm.DB {
	conn, err := gorm.Open(mysql.Open(global.DBConn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Print(err.Error())
	}
	return conn
}

//Setup : 数据库连接池设置
func Setup() {
	db = NewConnection()
	sqlDB, err := db.DB()
	log.Debug(err)

	sqlDB.SetMaxIdleConns(3)                    //最大空闲连接数
	sqlDB.SetMaxOpenConns(55)                   //最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 600) //设置连接空闲超时
	//db.LogMode(true)
}

//GetDB : 获取db对象公共方法
func GetDB() *gorm.DB {
	sqlDB, err := db.DB()
	log.Debug(err)
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		db = NewConnection()
	}
	return db
}
