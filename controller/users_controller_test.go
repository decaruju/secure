package controller

import (
	"bytes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"net/http/httptest"
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
	db.AutoMigrate(&model.ApiKey{})
}

func TestDB(t *testing.T) {
	initDB()
	db, err := db()
	if err != nil {
		t.Errorf("Could not open database")
	}
	defer db.Close()
}

func TestCreate(t *testing.T) {
	rr := mockCall("/users", "POST", "{}", t)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("\nStatus code differs.\nExpected %d.\nGot %d", http.StatusBadRequest, status)
	}
}

func mockRouter() *mux.Router {
	r := mux.NewRouter()
	UsersRouter(r)
	return r
}

func mockCall(route string, method string, payload string, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, route, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockRouter().ServeHTTP(rr, req)

	return rr
}
