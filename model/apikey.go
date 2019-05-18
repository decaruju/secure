package model

type Apikey struct {
	ID     uint `gorm:"primary_key"`
	Key    string
	UserID uint `gorm:"index"`
	User   User `gorm:"foreignkey:UserId"`
}
