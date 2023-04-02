package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Private(c *gin.Context) {
	c.JSON(http.StatusOK, "private")
}
