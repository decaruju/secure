package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"secure/controller"
	"secure/model"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.ApiKey{})

	router := mux.NewRouter()
	router.HandleFunc("/users/login", controller.Login)
	router.HandleFunc("/users", controller.Create)
	http.Handle("/", router)
	fmt.Println(http.ListenAndServe(":8081", nil))
}
