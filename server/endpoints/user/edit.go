package user

import (
	"art-gallery-server/database"
	"art-gallery-server/middleware/auth"
	"art-gallery-server/models/users"
	u "art-gallery-server/utils"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type EditRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func EditUserHandlers() []gin.HandlerFunc {
	var user auth.User
	var request EditRequest
	handlers := []gin.HandlerFunc{
		auth.Authorization(&user),
		u.BindRequest(&request),
		editUser(&request, &user),
	}
	return handlers
}
func editUser(r *EditRequest, u *auth.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		dbUser, err := users.GetUser(u.ID)

		if err != nil {
			log.Errorf("Failed get user (id = %d) from database: %s", u.ID, err)
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
