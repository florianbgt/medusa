package main

import (
	configs "florianbgt/medusa/internal/configs"
	db "florianbgt/medusa/internal/db"
	"florianbgt/medusa/internal/routing"
)

func main() {
	configs := configs.SetupConfigs()

	db_path := configs.SQLITE_FILE_PATH
	db := db.SetupDb(db_path)

	router := routing.SetupRouter(db, configs)

	router.Run()
}
