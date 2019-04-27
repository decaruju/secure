package controller

import (
	"gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func initDB() {
	db, err := gorm.Open("sqlite3", "test.db")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ApiKey{})
}

func destroyDB() {
	os.Remove("test.db")
}

func TestDB(t *testing.T) {
	db := db()
	if db == nil {
		t.Errorf("Could not open database")
	}
	defer db.Close()
}

func TestCreate(t *testing.T) {
	req, err := http.NewRequest("POST", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
}
