package utils

import (
	"art-gallery-server/database"
	"errors"
	"fmt"
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

// Bind JSON to request pointer, validate it, and if it not valid return BadRequest
func ValidateRequest(r Validatable) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := ctx.BindJSON(&r); err != nil {
			log.Errorf("Failed bind JSON data: %s", err)
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
func validateNotEmpty(value string, valueName string) error {
	if valueName == "" {
		valueName = "string field"
	}
	if value == "" {
		return errors.New(
			fmt.Sprintf("%s is empty", valueName))
	}
	return nil
}

func ConnectToDbOrAbort(ctx *gin.Context) *gorm.DB {
	db, err := database.ConnectToDb()
	if err != nil {
		log.Errorf("Failed to connect to database: %s", err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return nil
	}
	return db
}

func ValidateNotEmpty(value string) error {
	return validateNotEmpty(value, "")
}
func ValidateNotEmptyNamed(value string, valueName string) error {
	return validateNotEmpty(value, valueName)
}

func ValidateMoreThan(value int, secondValue int, valueName string) error {
	if value <= secondValue {
		return fmt.Errorf("%s is less than %d (%d)", valueName, secondValue, value)
	}
	return nil
}
