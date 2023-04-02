package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthy(c *gin.Context) {
	c.String(http.StatusOK, "healthy")
}
