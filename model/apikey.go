package model

import (
	"github.com/jinzhu/gorm"
)

type ApiKey struct {
	gorm.Model
	Key    string
	UserID uint `gorm:"index"`
	User   User `gorm:"foreignkey:UserId"`
}
