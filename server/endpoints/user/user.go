package user

import (
	"art-gallery-server/endpoints/user/avatar"
	"art-gallery-server/endpoints/user/create"
	"art-gallery-server/endpoints/user/delete"
	"art-gallery-server/endpoints/user/get"
	"art-gallery-server/endpoints/user/get_all"
	"art-gallery-server/endpoints/user/update"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	// get user by id
	router.POST("/get", get.Handlers()...)
	// create new user, aka register
	router.POST("/create", create.Handlers()...)
	// update user info
	router.POST("/update", update.Handlers()...)
	// get user avatar
	router.GET("/avatar/:id", avatar.Handlers()...)
	// get all users
	router.POST("/get-all", get_all.Handlers()...)
	// delete user
	router.DELETE("/delete", delete.Handlers()...)
}
