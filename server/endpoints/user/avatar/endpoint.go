package avatar

import (
	"art-gallery-server/models"
	"art-gallery-server/utils"
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
		var user models.User
		db := utils.ConnectToDbOrAbort(c)
		db.Where("id = ?", id).Select("email").First(&user)

		url := gravatar.NewAvatarUrl(user.Email, gravatar.DefaultImage(gravatar.DefaultRetro), gravatar.Size(512))
		c.JSON(200, gin.H{
			"url": url,
		})
	}
}

func Setup(router *gin.RouterGroup) {}
