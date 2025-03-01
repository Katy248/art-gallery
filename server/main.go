package main

import (
	"fmt"

	db "art-gallery-server/database"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

var (
	port             = 8080
	connectionString = "../art-gallery.db"
	serverMode       = gin.ReleaseMode
)

func main() {
	db.SetupConnectionString(connectionString)
	db.MustMigrateDb()
	SetServerMode(serverMode)
	addr := fmt.Sprintf(":%d", port)
	server := CreateServer(addr)
	server.Run(addr)
}

func SetServerMode(mode string) {
	log.Infof("Gin server mode - %s", mode)
	gin.SetMode(mode)
}
