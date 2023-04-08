package checks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthy(c *gin.Context) {
	c.JSON(http.StatusOK, "healthy")
}
