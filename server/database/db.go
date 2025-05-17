package database

import (
	m "art-gallery-server/models"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var connectionString = "../art-gallery.db"

func SetupConnectionStringFromConf() {
	conn := viper.GetString("database.connection_string")
	SetupConnectionString(conn)
}
func SetupConnectionString(conn string) {
	connectionString = conn
	log.Debugf("Connection string - '%s'", conn)
}

// Migrates database or exits with error
func MustMigrateDb() {
	setupImages()
	db := MustConnectToDb()
	err := db.AutoMigrate(&m.User{}, &m.Post{}, &m.PostSave{}, &m.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
	log.Info("Database migrated")
}

// Connects to db or exits
func MustConnectToDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	return db
}
func ConnectToDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		// log.Errorf("Failed to connect to database: %s", err)
		return nil, err
	}
	return db, nil
}
