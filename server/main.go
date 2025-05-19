package main

import (
	"fmt"

	db "art-gallery-server/database"
	"art-gallery-server/database/migration"
)

var (
	port = 8080
)

func init() {
	SetupConfiguration()
}

func main() {
	db.SetConnectionStringFromConf()
	db.MustConnect()

	migration.MustMigrateDb()

	addr := fmt.Sprintf(":%d", port)
	server := CreateServer(addr)
	server.Run(addr)
}
