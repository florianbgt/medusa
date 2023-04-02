package routing

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/handlers"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/web"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func serveApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		file_extention := filepath.Ext(c.Request.URL.Path)
		if file_extention == "" && c.Request.URL.Path != "/" {
			c.Request.URL.Path = c.Request.URL.Path + ".html"
		}

		embeded_app := web.BuildFS()
		directory := http.FS(embeded_app)
		file_server := http.FileServer(directory)

		file_server.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

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

	router.POST("api/token/refresh", func(c *gin.Context) {
		handlers.RefreshToken(
			c,
			configs,
		)
	})

	router.GET("api/private", isAuthenticated, handlers.Private)

	router.NoRoute(serveApp())

	return router
}
