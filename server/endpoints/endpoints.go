package endpoints

import (
	"art-gallery-server/endpoints/auth"
	"art-gallery-server/endpoints/post"
	"art-gallery-server/endpoints/user"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.RouterGroup) {
	post.Setup(router.Group("/post"))
	user.Setup(router.Group("/user"))
	auth.Setup(router)
}
