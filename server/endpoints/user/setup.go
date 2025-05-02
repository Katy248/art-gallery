package user

import "github.com/gin-gonic/gin"

func Setup(router *gin.RouterGroup) {
	// get user by id
	router.POST("/get", GetUserHandlers()...)
	// create new user, aka register
	router.POST("/create", CreateUserHandlers()...)
	// update user info
	router.POST("/edit", EditUserHandlers()...)
}
