package password_change

import (
	"florianbgt/medusa/internal/configs"
	"florianbgt/medusa/internal/models/password_model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangePassword(
	c *gin.Context,
	db *gorm.DB,
	configs *configs.Configs,
) {
	var payload struct {
		OldPassword string `json:"old_password" binding:"required"`
		Password    string `json:"password" binding:"required"`
		Password2   string `json:"password2" binding:"required"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	var Password password_model.Password

	currentPassword, err := Password.GetPassword(db)
	if err != nil {
		panic(err)
	}

	if !password_model.CheckPasswordHash(payload.OldPassword, configs.API_KEY, currentPassword) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "old_password_incorrect",
		})
		return
	}

	if payload.Password != payload.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password2_does_not_match",
		})
		return
	}

	err = Password.UpdatePassword(db, payload.Password, configs.API_KEY)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
