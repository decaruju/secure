package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique_index"`
	HashedPassword string
}

func (user *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
}

func (user *User) CreatePassword(password string) error {
	newHash, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return err
	}

	user.HashedPassword = string(newHash)
	return nil
}

func (user *User) UpdatePassword(oldPassword string, newPassword string) error {
	err := user.CheckPassword(oldPassword)
	if err != nil {
		return err
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), 16)
	if err != nil {
		return err
	}

	user.HashedPassword = string(newHash)
	return nil
}
