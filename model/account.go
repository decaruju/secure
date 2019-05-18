package model

type Account struct {
	ID     uint `gorm:"primary_key"`
	UserID uint `gorm:"index"`
	User   User `gorm:"foreignkey:UserId"`
}
