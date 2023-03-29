package routing

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/handlers"
	"florianbgt/medusa/internal/helpers"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, configs *configs.Configs) *gin.Engine {
	router := gin.Default()

	// Serve static files from website folder (static website)
	router.Use(static.Serve("/", static.LocalFile("./website", false)))

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
