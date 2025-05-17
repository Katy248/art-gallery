package post

import (
	"art-gallery-server/endpoints/post/all"
	"art-gallery-server/endpoints/post/create"
	"art-gallery-server/endpoints/post/delete"
	"art-gallery-server/endpoints/post/get"
	"art-gallery-server/endpoints/post/get_saved"
	"art-gallery-server/endpoints/post/get_users"
	"art-gallery-server/endpoints/post/save"
	"art-gallery-server/endpoints/post/unsave"
	"art-gallery-server/endpoints/post/update"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	// get users posts
	router.POST("/get-users", get_users.GetUsersPostHandlers()...)
	// create new post
	router.POST("/create", create.CreatePostHandlers()...)
	// save post
	router.POST("/save", save.Handlers()...)
	// unsave post
	router.POST("/unsave", unsave.Handlers()...)
	// get all saved post for specified user
	router.POST("/get-saved", get_saved.Handlers()...)
	// get all posts
	router.POST("/all", all.Handlers()...)
	// delete post
	router.DELETE("/delete", delete.Handlers()...)
	// get post by id
	router.POST("/get", get.Handlers()...)
	// update post info
	router.POST("/update", update.Handlers()...)
	// feed
	// router.POST("/feed")
}
