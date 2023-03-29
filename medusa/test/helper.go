package test

import (
	routing "florianbgt/medusa-api/medusa/routing"
	settings "florianbgt/medusa-api/medusa/settings"

	gin "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupApi() *gin.Engine {
	settings := settings.SetupSettings()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := routing.SetupRouter(db, settings)

	return router
}
