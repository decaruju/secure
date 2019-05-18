package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"secure/model"
	"testing"
)

func initDB() {
	os.Remove("test.db")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Apikey{})
}

func TestDB(t *testing.T) {
	initDB()
	db, err := db()
	if err != nil {
		t.Errorf("Could not open database")
	}
	defer db.Close()
}
