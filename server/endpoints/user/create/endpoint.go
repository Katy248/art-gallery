package create

import (
	"net/http"

	"art-gallery-server/database"
	"art-gallery-server/models/users"

	// "art-gallery-server/endpoints/user"
	"art-gallery-server/utils"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Handlers() []gin.HandlerFunc {
	request := &Request{}
	handlers := []gin.HandlerFunc{
		utils.BindRequest(request), handler(request),
	}
	return handlers
}

func handler(r *Request) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := users.NewUser(r.Name, r.Email, r.Password)
		if err != nil {
			log.Errorf("Error while creating new user: %s", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "invalid user data"})
			return
		}
		result := database.Conn.Create(&user)
		if result.Error != nil {
			log.Errorf("Failed to save user: %s", result.Error)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false, "error": "invalid user data, probably user with this email already exists"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true})
	}
}
