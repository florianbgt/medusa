package printer

import (
	"florianbgt/medusa/internal/models/printer_model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPrinter(c *gin.Context, db *gorm.DB) {
	var p printer_model.Printer

	printer := p.GetPrinter(db)

	if printer == nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(200, printer)
}

func SetPrinter(c *gin.Context, db *gorm.DB) {
	var payload struct {
		X float32 `json:"x" binding:"required"`
		Y float32 `json:"y" binding:"required"`
		Z float32 `json:"z" binding:"required"`
	}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	var printer printer_model.Printer

	printer.X = payload.X
	printer.Y = payload.Y
	printer.Z = payload.Z

	printer.SetPrinter(db)

	c.JSON(200, printer)
}
