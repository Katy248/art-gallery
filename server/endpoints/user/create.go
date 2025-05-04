package user

import (
	"errors"
	"net/http"

	m "art-gallery-server/models"
	u "art-gallery-server/utils"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *createUserRequest) Validate() error {
	return errors.Join(
		u.ValidateNotEmptyNamed(r.Name, "name"),
		u.ValidateNotEmptyNamed(r.Email, "email"),
		u.ValidateNotEmptyNamed(r.Password, "password"),
	)
}

func CreateUserHandlers() []gin.HandlerFunc {
	request := &createUserRequest{}
	handlers := []gin.HandlerFunc{
		u.ValidateRequest(request), createUser(request),
	}
	return handlers
}

func createUser(r *createUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := m.NewUser(r.Name, r.Email, r.Password)
		if err != nil {
			log.Errorf("Error while creating new user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		db := u.ConnectToDbOrAbort(ctx)
		if err := user.Save(db); err != nil {
			log.Errorf("Failed to save user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}
