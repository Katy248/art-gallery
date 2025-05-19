package main

import (
	"art-gallery-server/database"
	e "art-gallery-server/endpoints"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetServerMode() {
	mode := viper.GetString("server.mode")
	gin.SetMode(mode)
	log.Infof("Gin server mode - %s", gin.Mode())
}
func CreateServer(addr string) *gin.Engine {
	SetServerMode()

	server := gin.Default()

	api := server.Group("/api")
	{
		e.Setup(api)
		api.Static("/images", database.ImagesDir())
	}
	return server
}
