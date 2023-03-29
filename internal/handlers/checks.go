package handlers

import (
	gin "github.com/gin-gonic/gin"
)

func Healthy(c *gin.Context) {
	c.String(200, "healthy")
}
