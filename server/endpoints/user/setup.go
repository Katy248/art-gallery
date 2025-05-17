package user

import (
	"art-gallery-server/endpoints/user/avatar"
	"art-gallery-server/endpoints/user/get_all"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	// get user by id
	router.POST("/get", GetUserHandlers()...)
	// create new user, aka register
	router.POST("/create", CreateUserHandlers()...)
	// update user info
	router.POST("/edit", EditUserHandlers()...)
	// get user avatar
	router.GET("/avatar/:id", avatar.Handlers()...)
	// get all users
	router.POST("/get-all", get_all.Handlers()...)
}
