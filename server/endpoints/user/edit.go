package user

import (
	"art-gallery-server/middleware/auth"
	m "art-gallery-server/models"
	u "art-gallery-server/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type editRequest struct {
	Name string `json:"name"`
}

func (r *editRequest) Validate() error {
	return errors.Join(u.ValidateNotEmptyNamed(r.Name, "name"))
}

func EditUserHandlers() []gin.HandlerFunc {
	var user auth.AuthUser
	var request editRequest
	handlers := []gin.HandlerFunc{
		auth.Authorization(&user),
		u.ValidateRequest(&request),
		editUser(&request, &user),
	}
	return handlers
}
func editUser(r *editRequest, user *auth.AuthUser) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := u.ConnectToDbOrAbort(ctx)

		var dbUser m.User
		query := fmt.Sprintf("id = %d", user.ID)
		result := db.Model(&m.User{}).First(&dbUser, query)

		if result.Error != nil {
			log.Errorf("Failed get user (id = %d) from database: %s", user.ID, result.Error)
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		db.Model(&dbUser).Update("Name", r.Name)
		ctx.AbortWithStatus(http.StatusOK)
	}
}
