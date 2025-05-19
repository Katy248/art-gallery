package auth

import (
	"art-gallery-server/database"
	"art-gallery-server/models/users"
	"art-gallery-server/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func init() {
	jwtKey = []byte(viper.GetString("server.jwt_key"))
}

var (
	jwtKey       []byte
	signinMethod = jwt.SigningMethodHS256
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
type User struct {
	ID    int
	Email string
}

func AuthHandlers() []gin.HandlerFunc {
	var request AuthRequest
	handlers := []gin.HandlerFunc{
		utils.BindRequest(&request),
		auth(&request),
	}
	return handlers
}

func auth(r *AuthRequest) gin.HandlerFunc {
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

func Authorization(u *User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(authHeader) < 2 || authHeader[0] != bearer {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			log.Warn("Authorization failed: bad auth header")
			return
		}
		tokenString := authHeader[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			log.Warnf("Authorization failed: %s", err)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		u.Email = claims["email"].(string)
		u.ID = int(claims["id"].(float64))

	}
}

func GetUser(ctx *gin.Context) (*User, error) {
	authHeader := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
	if len(authHeader) < 2 || authHeader[0] != bearer {
		log.Warn("Authorization failed: bad auth header")
		return nil, fmt.Errorf("bad auth header")
	}
	tokenString := authHeader[1]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		log.Warnf("Authorization failed: %s", err)
		return nil, fmt.Errorf("auth failed: %s", err)
	}
	claims := token.Claims.(jwt.MapClaims)
	user := User{
		Email: claims["email"].(string),
		ID:    int(claims["id"].(float64)),
	}
	return &user, nil
}

const (
	bearer = "Bearer"
)
