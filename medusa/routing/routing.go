package routing

import (
	handlers "florianbgt/medusa-api/medusa/handlers"
	"florianbgt/medusa-api/medusa/helpers"
	settings "florianbgt/medusa-api/medusa/settings"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, settings *settings.Settings) *gin.Engine {
	router := gin.Default()

	isAuthenticated := func(c *gin.Context) {
		helpers.IsAuthCheck(c, settings.API_KEY)
	}

	router.GET("api/healthy", handlers.Healthy)

	router.POST("api/login", func(c *gin.Context) {
		handlers.Login(
			c,
			db,
			settings,
		)
	})

	router.Use(isAuthenticated).GET("api/private", handlers.Private)

	return router
}
