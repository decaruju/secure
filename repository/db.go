package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func db() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	db.LogMode(true)
	return db, err
}
