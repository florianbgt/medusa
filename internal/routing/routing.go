package routing

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/handlers/checks"
	"florianbgt/medusa/internal/handlers/files"
	"florianbgt/medusa/internal/handlers/login"
	"florianbgt/medusa/internal/handlers/password_change"
	"florianbgt/medusa/internal/handlers/refresh"
	"florianbgt/medusa/internal/handlers/stream"
	"florianbgt/medusa/internal/handlers/system"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/web"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func serveApp() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Rediect to index.html (SPA)
		file_extention := filepath.Ext(c.Request.URL.Path)
		if file_extention == "" {
			c.Request.URL.Path = "/"
		}

		embeded_app := web.BuildFS()
		directory := http.FS(embeded_app)
		file_server := http.FileServer(directory)

		file_server.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func corsMiddleware(debug bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if debug {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SetupRouter(db *gorm.DB, configs *configs.Configs) *gin.Engine {
	if configs.DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.Use(corsMiddleware(configs.DEBUG))

	isAuthenticated := func(c *gin.Context) {
		helpers.IsAuthCheck(c, configs.API_KEY)
	}

	router.GET("api/healthy", checks.Healthy)
	router.GET("api/authenticated", isAuthenticated, checks.Healthy)

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
	router.POST("api/password/change", isAuthenticated, func(c *gin.Context) {
		password_change.ChangePassword(
			c,
			db,
			configs,
		)
	})

	router.GET("api/system", isAuthenticated, system.SystemInfo)
	router.GET("api/system/metrics", isAuthenticated, system.SystemMetrics)

	router.GET("api/stream", isAuthenticated, func(c *gin.Context) {
		if configs.ENABLE_CAMERA {
			stream.Stream(
				c,
				configs.API_KEY,
			)
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	})

	router.GET("api/files", isAuthenticated, files.ListFiles)
	router.POST("api/files", isAuthenticated, files.UploadFile)
	router.DELETE("api/files/:name", isAuthenticated, files.DeleteFile)
	router.GET("api/files/:name/gcode/info", isAuthenticated, files.GetGCodeInfo)
	router.GET("api/files/:name/gcode", isAuthenticated, files.GetGCode)

	router.NoRoute(serveApp())

	return router
}
