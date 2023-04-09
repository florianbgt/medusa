package test

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/routing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupConfigs() *configs.Configs {
	return &configs.Configs{
		PORT:             "8080",
		API_KEY:          "API_KEY",
		DEFAULT_PASSWORD: "Password/123",
		ENABLE_CAMERA:    true,
		CAMERA_NAME:      "/dev/video0",
	}
}

func SetupApi() *gin.Engine {
	configs := SetupConfigs()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := routing.SetupRouter(db, configs)

	return router
}

func Setupdb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
