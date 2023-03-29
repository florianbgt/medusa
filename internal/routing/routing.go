package routing

import (
	"florianbgt/medusa/internal/configs"
	handlers "florianbgt/medusa/internal/handlers"
	"florianbgt/medusa/internal/helpers"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, configs *configs.Configs) *gin.Engine {
	router := gin.Default()

	isAuthenticated := func(c *gin.Context) {
		helpers.IsAuthCheck(c, configs.API_KEY)
	}

	router.GET("api/healthy", handlers.Healthy)

	router.POST("api/login", func(c *gin.Context) {
		handlers.Login(
			c,
			db,
			configs,
		)
	})

	router.Use(isAuthenticated).GET("api/private", handlers.Private)

	return router
}
