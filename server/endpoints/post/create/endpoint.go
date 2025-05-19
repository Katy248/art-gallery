package create

import (
	"art-gallery-server/database"
	auth "art-gallery-server/middleware/auth"
	"art-gallery-server/models/posts"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func CreatePostHandlers() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		auth.Middleware(),
		createPost(),
	}
}

func createPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		desc := ctx.Request.Form.Get("description")
		if desc != "" {
			log.Errorf("Failed bind form data: description field is empty")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "description field is empty", "success": false})
			return
		}

		file, err := ctx.FormFile("picture")
		if err != nil {
			log.Errorf("Failed get file: %s", err)
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		post := posts.Post{
			PublisherID: auth.UserId(ctx),
			Description: desc,
		}

		database.Conn.Model(&post).Create(&post)

		imageName := fmt.Sprintf("picture-%d", post.ID)
		post.ImageUrl = fmt.Sprintf("/api/images/%s", imageName)

		database.Conn.Save(&post)
		err = ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s", database.ImagesDir(), imageName))
		if err != nil {
			log.Errorf("Failed save file: %s", err)
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true, "id": post.ID})
	}
}
