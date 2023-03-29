package main

import (
	db "florianbgt/medusa-api/medusa/db"
	"florianbgt/medusa-api/medusa/routing"
	settings "florianbgt/medusa-api/medusa/settings"
)

func main() {
	settings := settings.SetupSettings()

	db_path := settings.SQLITE_FILE_PATH
	db := db.SetupDb(db_path)

	router := routing.SetupRouter(db, settings)

	router.Run()
}
