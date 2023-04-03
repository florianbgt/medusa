package login

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(
	c *gin.Context,
	db *gorm.DB,
	configs *configs.Configs,
) {
	var payload struct {
		Password string `json:"password" binding:"required"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	if payload.Password != configs.DEFAULT_PASSWORD {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password_incorrect",
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
