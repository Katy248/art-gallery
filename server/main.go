package main

import (
	"fmt"

	"github.com/charmbracelet/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	port             = 8080
	ConnectionString = "../art-gallery.db"
)

func main() {
	MigrateDb(ConnectionString)
	addr := fmt.Sprintf(":%d", port)
	server := CreateServer(addr)
	server.Run(addr)
}

// Migrates database or exits with error
func MigrateDb(conn string) {
	db := ConnectToDbOrExit(conn)
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %s", err)
	}
	log.Info("Database migrated")
}

func ConnectToDbOrExit(conn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	return db
}
func ConnectToDb(conn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(conn), &gorm.Config{})
	if err != nil {
		// log.Errorf("Failed to connect to database: %s", err)
		return nil, err
	}
	return db, nil
}
