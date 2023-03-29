package routing

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/handlers"
	"florianbgt/medusa/internal/helpers"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Renders the static website if the path is not an API route
func spaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Next()
			return
		}

		directory := http.FileSystem(http.Dir("./website"))
		file_server := http.FileServer(directory)

		file_extention := filepath.Ext(c.Request.URL.Path)
		if file_extention == "" && c.Request.URL.Path != "/" {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, c.Request.URL.Path, c.Request.URL.Path+".html", 1)
		}

		fmt.Println(filepath.Ext(c.Request.URL.Path))

		file_server.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func SetupRouter(db *gorm.DB, configs *configs.Configs) *gin.Engine {
	router := gin.Default()

	// Serve static website
	router.Use(spaMiddleware())

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
