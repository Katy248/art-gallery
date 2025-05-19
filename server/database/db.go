package database

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Conn             *gorm.DB
	connectionString = "../art-gallery.db"
)

func SetConnectionStringFromConf() {
	conn := viper.GetString("database.connection_string")
	SetConnectionString(conn)
}
func SetConnectionString(conn string) {
	connectionString = conn
	log.Debugf("Connection string - '%s'", conn)
}

func MustConnect() {
	var err error
	Conn, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connect to database '%s': %s", connectionString, err)
	}
}
