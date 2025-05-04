package user

import (
	"art-gallery-server/endpoints/user/avatar"

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
}
