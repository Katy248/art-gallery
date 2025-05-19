package auth

import (
	"art-gallery-server/database"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func AuthHandlers() []gin.HandlerFunc {
	var request AuthRequest
	handlers := []gin.HandlerFunc{
		utils.BindRequest(&request),
		authenticateUser(&request),
	}
	return handlers
}

func authenticateUser(r *AuthRequest) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		query := fmt.Sprintf("email = '%s'", r.Email)
		var u users.User
		result := database.Conn.Model(&users.User{}).First(&u, query)
		if result.Error != nil {
			log.Warnf("Error while getting user '%s': %s", r.Email, result.Error)
			ctx.JSON(http.StatusNotFound, AuthResponse{
				Success: false,
				Error:   "Wrong email or password",
			})
			// TODO: return response with error message for frontend
			return
		}
		if !u.CheckPassword(r.Password) {
			log.Warnf("Wrong password for user '%s'", r.Email)
			ctx.JSON(http.StatusNotFound, AuthResponse{
				Success: false,
				Error:   "Wrong email or password",
			})
			// TODO: return response with error message for frontend
			return
		}

		token := jwt.NewWithClaims(signinMethod, jwt.MapClaims{
			"id":      u.ID,
			"email":   u.Email,
			"name":    u.Name,
			"isAdmin": u.IsAdmin,
		})
		stringToken, err := token.SignedString(jwtKey)
		if err != nil {
			log.Errorf("Failed to sign jwt token: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, AuthResponse{Success: true, Token: stringToken})
	}
}
