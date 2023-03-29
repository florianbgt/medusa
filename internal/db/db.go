package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDb(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to database")

	// fmt.Println("Migrating User model...")
	// db.AutoMigrate(models.User{})

	return db
}
