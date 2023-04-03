package routing

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/handlers/checks"
	"florianbgt/medusa/internal/handlers/login"
	"florianbgt/medusa/internal/handlers/password_change"
	"florianbgt/medusa/internal/handlers/private"
	"florianbgt/medusa/internal/handlers/refresh"
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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter(db *gorm.DB, configs *configs.Configs) *gin.Engine {
	router := gin.Default()

	router.Use(corsMiddleware())

	isAuthenticated := func(c *gin.Context) {
		helpers.IsAuthCheck(c, configs.API_KEY)
	}

	router.GET("api/healthy", checks.Healthy)

	router.POST("api/login", func(c *gin.Context) {
		login.Login(
			c,
			db,
			configs,
		)
	})

	router.POST("api/token/refresh", func(c *gin.Context) {
		refresh.RefreshToken(
			c,
			configs,
		)
	})

	router.POST("api/password/change", func(c *gin.Context) {
		password_change.ChangePassword(
			c,
			db,
			configs,
		)
	})

	router.GET("api/private", isAuthenticated, private.Private)

	router.NoRoute(serveApp())

	return router
}
