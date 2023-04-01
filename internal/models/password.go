package models

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Password struct {
	gorm.Model
	Password string
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}
	if password == strings.ToUpper(password) {
		return false
	}
	if password == strings.ToLower(password) {
		return false
	}
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}
	if !strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;':,./<>?") {
		return false
	}
	return true
}

func (p *Password) SetPassword(password string) error {
	if !isPasswordValid(password) {
		return errors.New("Password is not valid, must contains at least 8 characters, 1 uppercase, 1 lowercase, 1 number and 1 special character")
	}
	p.Password = password
	return nil
}

func SetupPassword(db *gorm.DB, password string) {
	fmt.Println("Migrating Password model...")
	db.AutoMigrate(Password{})

	var passwordModel Password
	db.First(&passwordModel)

	if passwordModel.ID == 0 {
		err := passwordModel.SetPassword(password)
		if err != nil {
			panic(err)
		}
		db.Create(&passwordModel)
	}
}
