package migration

import (
	"art-gallery-server/database"
	"art-gallery-server/models/posts"
	"art-gallery-server/models/users"

	"github.com/charmbracelet/log"
)

// Migrates database or exits with error
func MustMigrateDb() {
	database.SetupImages()
	err := database.Conn.AutoMigrate(
		&users.User{},
		&posts.Post{},
		&posts.PostSave{},
		&posts.Comment{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
	log.Info("Database migrated")
}
