package main

import (
	"net/http"

	"art-gallery-server/database"
	e "art-gallery-server/endpoints"
	m "art-gallery-server/models"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func CreateServer(addr string) *gin.Engine {

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
			users.POST("/", e.GetUserHandlers()...)
			// create new user, aka register
			users.POST("/create", e.CreateUserHandlers()...)
			// update user info
			users.POST("/update")
		}
		post := api.Group("/post")
		{
			// get post by id
			post.GET("/:id")
			// get users posts
			post.GET("/users/:id")
			// create new post
			post.POST("/")
			// update post info
			post.POST("/update")
			// delete post
			post.DELETE("/:id")
			// save post
			post.GET("/save/:id")
			// unsave post
			post.GET("/unsave/:id")
			// feed
			post.GET("/feed/:lastId")
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
