package main

import (
	"fmt"

	db "art-gallery-server/database"
	"art-gallery-server/database/migration"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	port             = 8080
	connectionString = "../art-gallery.db"
	serverMode       = gin.ReleaseMode
)

func init() {
	SetupConfiguration()
	SetupLogging()
}

func main() {
	db.SetConnectionStringFromConf()
	db.MustConnect()

	migration.MustMigrateDb()

	addr := fmt.Sprintf(":%d", port)
	server := CreateServer(addr)
	server.Run(addr)
}

func SetupLogging() {
	level := viper.GetString("log.level")
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}

	log.Infof("Log level - %s", level)
}
