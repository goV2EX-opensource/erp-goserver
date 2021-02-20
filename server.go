package main

import (
	"fmt"

	"example.com/go-v2ex/dbutil"
	"example.com/go-v2ex/global"
	"example.com/go-v2ex/migrate"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	global.Init()
	log.Info("****** APP STARTED ******")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	global.Loadconf()
	dbutil.Setup()
	migrate.CheckSalt()
	migrate.CheckSysAdmin()

	port := fmt.Sprintf(":%d", global.ListenPort)
	log.Info(fmt.Sprintf("***WEBSERVER WILL LISTEN ON PORT :%d ***", global.ListenPort))

	r.Run(port)
}
