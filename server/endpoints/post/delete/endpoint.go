package delete

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/utils"

	"github.com/gin-gonic/gin"
)

type Request struct {
	PostId int `json:"postId" binding:""`
}

func Handlers() []gin.HandlerFunc {
	var user auth.User
	var r Request
	return []gin.HandlerFunc{
		auth.Authorization(&user),
		utils.BindRequest(&r),
		handler(&r, &user),
	}
}

func handler(request *Request, user *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
