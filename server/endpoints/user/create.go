package user

import (
	"net/http"

	"art-gallery-server/database"
	"art-gallery-server/models/users"

	// "art-gallery-server/endpoints/user"
	"art-gallery-server/utils"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func CreateUserHandlers() []gin.HandlerFunc {
	request := &createUserRequest{}
	handlers := []gin.HandlerFunc{
		utils.BindRequest(request), createUser(request),
	}
	return handlers
}

func createUser(r *createUserRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := users.NewUser(r.Name, r.Email, r.Password)
		if err != nil {
			log.Errorf("Error while creating new user: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		result := database.Conn.Create(&user)
		if result.Error != nil {
			log.Errorf("Failed to save user: %s", result.Error)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}
