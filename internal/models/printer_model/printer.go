package printer_model

import (
	"gorm.io/gorm"
)

type Printer struct {
	gorm.Model
	Name string
	X    int
	Y    int
	Z    int
}

func (p *Printer) Setup(db *gorm.DB) {
	db.AutoMigrate(Printer{})
}
