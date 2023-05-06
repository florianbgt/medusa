package printer_model

import (
	"gorm.io/gorm"
)

type Printer struct {
	gorm.Model
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (p *Printer) Setup(db *gorm.DB) {
	db.AutoMigrate(Printer{})
}

func (p *Printer) SetPrinter(db *gorm.DB) {
	db.First(p)

	if p.ID == 0 {
		db.Create(p)
	} else {
		db.Save(p)
	}
}

func (p *Printer) GetPrinter(db *gorm.DB) *Printer {
	db.First(p)

	if p.ID == 0 {
		return nil
	} else {
		return p
	}
}
