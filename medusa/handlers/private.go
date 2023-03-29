package handlers

import gin "github.com/gin-gonic/gin"

func Private(c *gin.Context) {
	c.JSON(200, "private")
}
