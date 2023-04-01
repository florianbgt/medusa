package db

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb(path string, configs *configs.Configs) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	models.SetupPassword(db, configs.DEFAULT_PASSWORD)

	return db
}
