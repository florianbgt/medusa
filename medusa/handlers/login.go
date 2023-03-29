package handlers

import (
	"florianbgt/medusa-api/medusa/helpers"
	settings "florianbgt/medusa-api/medusa/settings"
	fmt "fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"
	gorm "gorm.io/gorm"
)

type payload struct {
	Password string `json:"password" binding:"required"`
}

func Login(
	c *gin.Context,
	db *gorm.DB,
	settings *settings.Settings,
) {
	var payload payload
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	if payload.Password != settings.PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	tokenString, err := helpers.GenerateToken(settings.API_KEY, settings.TOKEN_EXPIRATION)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": tokenString,
	})
}
