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
		users := api.Group("/user")
		{
			// get user by id
			users.POST("/get", e.GetUserHandlers()...)
			// create new user, aka register
			users.POST("/create", e.CreateUserHandlers()...)
			// update user info
			users.POST("/edit", e.EditUserHandlers()...)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/", e.AuthHandlers()...)
		}
		post := api.Group("/post")
		{
			// get post by id
			post.GET("/get")
			// get users posts
			post.GET("/get-users")
			// create new post
			post.POST("/create")
			// update post info
			post.POST("/update")
			// delete post
			post.DELETE("/delete")
			// save post
			post.POST("/save")
			// unsave post
			post.POST("/unsave")
			// feed
			post.POST("/feed")
		}
		gallery := api.Group("/gallery")
		{
			gallery.GET("/:id")
			gallery.POST("/")
			gallery.POST("/update")
			gallery.DELETE("/")
		}
	}
	return server
}
