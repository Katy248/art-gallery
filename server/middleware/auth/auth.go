package auth

import (
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

const (
	Bearer = "Bearer"

	UserIdKey      = "USER_ID_______KEY"
	UserEmailKey   = "USER_EMAIL____KEY"
	UserIsAdminKey = "USER_IS_ADMIN_KEY"
)

func UserId(ctx *gin.Context) int {
	id := ctx.MustGet(UserIdKey).(float64)
	return int(id)
}
func UserIsAdmin(ctx *gin.Context) bool {
	return ctx.MustGet(UserIsAdminKey).(bool)
}

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, email, isAdmin, err := GetAuthInfo(ctx.Request)
		if err != nil {
			log.Error("Authorization failed", "error", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"success":    false,
				"message":    "Authorization failed: " + err.Error(),
			})
			return
		}

		ctx.Set(UserIdKey, id)
		ctx.Set(UserEmailKey, email)
		ctx.Set(UserIsAdminKey, isAdmin)
	}
}

func GetAuthInfo(req *http.Request) (id float64, email string, isAdmin bool, err error) {
	header := req.Header.Get("Authorization")
	if header == "" {
		return id, email, isAdmin, fmt.Errorf("header Authorization is empty")
	}
	words := strings.Split(header, " ")
	if len(words) < 2 {
		return id, email, isAdmin, fmt.Errorf("bad auth header: words length less than 2")
	}
	if words[0] != Bearer {
		return id, email, isAdmin, fmt.Errorf("bad auth header: there is no %q specified", Bearer)
	}
	tokenString := words[1]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
	if err != nil {
		return id, email, isAdmin, fmt.Errorf("failed parse JWT: %s", err)
	}
	claims := token.Claims.(jwt.MapClaims)
	id, ok := claims["id"].(float64)
	if !ok {
		return id, email, isAdmin, fmt.Errorf("id key not specified in payload")
	}
	email, ok = claims["email"].(string)
	if !ok {
		return id, email, isAdmin, fmt.Errorf("email key not specified in payload")
	}

	isAdmin, ok = claims["isAdmin"].(bool)
	if !ok {
		return id, email, isAdmin, fmt.Errorf("isAdmin key not specified in payload")
	}
	return id, email, isAdmin, nil
}
