package delete

import (
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func Handlers() []gin.HandlerFunc {
	var r Request
	return []gin.HandlerFunc{
		auth.Middleware(),
		utils.BindRequest(&r),
		handler(&r),
	}
}

type Request struct {
	ID int `json:"id" binding:"required,gte=1"`
}

func handler(r *Request) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !users.IsAdmin(auth.UserId(c)) {
			log.Error("Can't delete user if requester is not an admin", "userId", auth.UserId(c))
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "deleting users is not allowed", "success": false})
			return
		}
		if err := users.Delete(r.ID); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "success": false})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
