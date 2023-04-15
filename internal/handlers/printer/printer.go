package printer

import (
	"florianbgt/medusa/internal/models/printer_model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPrinters(c *gin.Context, db *gorm.DB) {
	var Printer printer_model.Printer

	printers := db.Find(&Printer)

	c.JSON(200, printers)
}
