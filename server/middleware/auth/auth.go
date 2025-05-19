package auth

import (
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

		authHeader := strings.Split(ctx.Request.Header.Get("Authorization"), " ")
		if len(authHeader) < 2 || authHeader[0] != Bearer {
			log.Warn("Authorization failed: bad auth header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"statusCode": http.StatusUnauthorized,
				"success":    false,
			})
			return
		}
		tokenString := authHeader[1]
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return jwtKey, nil })
		if err != nil {
			log.Warnf("Authorization failed: %s", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":    false,
				"statusCode": http.StatusUnauthorized,
			})
			return
		}
		claims := token.Claims.(jwt.MapClaims)

		ctx.Set(UserIdKey, claims["id"].(float64))
		ctx.Set(UserEmailKey, claims["email"])
		ctx.Set(UserIsAdminKey, claims["isAdmin"].(bool))
	}
}
