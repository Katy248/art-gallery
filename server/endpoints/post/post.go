package post

import (
	"art-gallery-server/endpoints/post/create"
	"art-gallery-server/endpoints/post/get_users"
	"art-gallery-server/endpoints/post/save"
	"art-gallery-server/endpoints/post/unsave"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	// get post by id
	router.GET("/get")
	// get users posts
	router.POST("/get-users", get_users.GetUsersPostHandlers()...)
	// create new post
	router.POST("/create", create.CreatePostHandlers()...)
	// update post info
	router.POST("/update")
	// delete post
	router.DELETE("/delete")
	// save post
	router.POST("/save", save.Handlers()...)
	// unsave post
	router.POST("/unsave", unsave.Handlers()...)
	// feed
	router.POST("/feed")
}
