package handlers

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshToken(
	c *gin.Context,
	configs *configs.Configs,
) {
	var payload struct {
		Refresh string `json:"refresh" binding:"required"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	if !helpers.IsTokenValid(payload.Refresh, configs.API_KEY) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	token_pair, err := helpers.GenerateTokenPair(configs.API_KEY)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token_pair.Access,
		"refresh_token": token_pair.Refresh,
	})
}
