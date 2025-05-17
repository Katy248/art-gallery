package users

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

func IsAdmin(db *gorm.DB, userId int) (isAdmin bool) {
	result := db.Raw(rawIsAdminQuery, userId).First(&isAdmin)
	if result.Error != nil {
		log.Errorf("Failed to check if user is admin, returning default false value. Error: %s", result.Error)
		return false
	}
	return isAdmin
}

var rawIsAdminQuery = `
	SELECT is_admin FROM users WHERE id = ?
`
