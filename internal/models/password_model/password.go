package password_model

import (
	"errors"
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

func (p *Password) UpdatePassword(db *gorm.DB, password string) error {
	db.First(p)

	if p.ID == 0 {
		panic("no_password_in_db")
	}

	if !isPasswordValid(password) {
		return errors.New("invalid_password")
	}

	p.Password = password

	db.Save(p)
	return nil
}

func (p *Password) GetPassword(db *gorm.DB) (string, error) {
	db.First(p)

	if p.ID == 0 {
		return "", errors.New("no_password_in_db")
	}

	return p.Password, nil
}

func (p *Password) Setup(db *gorm.DB, password string) {
	db.AutoMigrate(Password{})

	db.First(p)

	if p.ID == 0 {
		if !isPasswordValid(password) {
			panic("Password is not valid, must contains at least 8 characters, 1 uppercase, 1 lowercase, 1 number and 1 special character")
		}
		p.Password = password
		db.Create(&p)
	}
}
