package db

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/models/password_model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb(path string, configs *configs.Configs) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var passwordInstance password_model.Password
	passwordInstance.Setup(db, configs.DEFAULT_PASSWORD, configs.API_KEY)

	return db
}
