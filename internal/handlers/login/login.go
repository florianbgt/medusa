package login

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/helpers"
	"florianbgt/medusa/internal/models/password_model"
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

	var passwordInstance password_model.Password
	currentPassword, err := passwordInstance.GetPassword(db)
	if err != nil {
		panic(err)
	}

	if !password_model.CheckPasswordHash(payload.Password, configs.API_KEY, currentPassword) {
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
