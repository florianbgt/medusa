package test

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/routing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupApi() *gin.Engine {
	configs := configs.SetupConfigs()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := routing.SetupRouter(db, configs)

	return router
}
