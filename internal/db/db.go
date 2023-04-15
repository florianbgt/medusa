package db

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/models/password_model"
	"florianbgt/medusa/internal/models/printer_model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb(path string, configs *configs.Configs) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var Password password_model.Password
	Password.Setup(db, configs.DEFAULT_PASSWORD, configs.API_KEY)

	var Printer printer_model.Printer
	Printer.Setup(db)

	return db
}
