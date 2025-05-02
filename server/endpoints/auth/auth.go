package auth

import (
	middleware "art-gallery-server/middleware/auth"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	router.POST("/auth", middleware.AuthHandlers()...)
}
