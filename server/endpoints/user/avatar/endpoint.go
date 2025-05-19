package avatar

import (
	"art-gallery-server/models/users"
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	gravatar "github.com/katy248/gravatar/pkg/url"
)

func Handlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{handler()}
}

func handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			log.Warnf("Inavalid id %s: %s", idParam, err.Error())
			c.JSON(400, gin.H{
				"message": "invalid id",
			})
			return
		}

		user, err := users.GetUser(id)
		if err != nil {
			log.Errorf("Failed get user for avatar image")
			user.Email = "empty email" //default email string
		}

		url := gravatar.NewAvatarUrl(user.Email, gravatar.DefaultImage(gravatar.DefaultRetro), gravatar.Size(512))
		c.JSON(http.StatusOK, gin.H{
			"url": url,
		})
	}
}

func Setup(router *gin.RouterGroup) {}
