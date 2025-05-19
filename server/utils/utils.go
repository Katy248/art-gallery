package utils

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Validatable interface {
	Validate() error
}

// Binds JSON to request pointer
func BindRequest(r interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.BindJSON(&r); err != nil {
			log.Errorf("Failed bind JSON data: %s", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
}

func ValidateUri(r Validatable) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.ShouldBindUri(&r); err != nil {
			log.Errorf("Failed bind uri parameters: %s", err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if err := r.Validate(); err != nil {
			log.Errorf("Validation failed: %s", err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}
}

// Returns true if no errors occurred, otherwise false
func WrapDbResult(ctx *gin.Context, result *gorm.DB) bool {

	if result.Error != nil {
		log.Errorf("Failed perform database action: %s", result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"success": false,
		})
		return false
	}

	return true
}
