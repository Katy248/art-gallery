package utils

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Binds JSON to request pointer
func BindRequest(r interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.BindJSON(&r); err != nil {
			log.Errorf("Failed bind JSON data: %s", err)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
}

// Returns true if no errors occurred, otherwise false
func WrapDbResult(ctx *gin.Context, result *gorm.DB) bool {

	if result.Error != nil {
		log.Errorf("Failed perform database action: %s", result.Error)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
		})
		return false
	}

	return true
}
