package main

import (
	"net/http"

	"art-gallery-server/database"
	e "art-gallery-server/endpoints"
	m "art-gallery-server/models"

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

	if gin.Mode() == gin.DebugMode {
		server.GET("/test", func(ctx *gin.Context) {
			log.Info("Test endpoint")
			db := database.MustConnectToDb()

			var users []m.User
			db.Model(&m.User{}).Find(&users)
			ctx.JSON(http.StatusOK, users)
		})
	}
	api := server.Group("/api")
	{
		e.Setup(api)
		// gallery := api.Group("/gallery")
		// {
		// 	gallery.GET("/:id")
		// 	gallery.POST("/")
		// 	gallery.POST("/update")
		// 	gallery.DELETE("/")
		// }
		api.Static("/images", database.ImagesDir())
	}
	return server
}
