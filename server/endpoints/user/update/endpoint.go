package update

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/users"
	u "art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func Handlers() []gin.HandlerFunc {
	var request Request
	handlers := []gin.HandlerFunc{
		auth.Middleware(),
		u.BindRequest(&request),
		handler(&request),
	}
	return handlers
}
func handler(r *Request) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dbUser, err := users.GetUser(auth.UserId(ctx))

		if err != nil {
			log.Errorf("Failed get user (id = %d) from database: %s", auth.UserId(ctx), err)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		database.Conn.Model(&dbUser).Update("Name", r.Name)
		ctx.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"success":    true,
		})
	}
}
