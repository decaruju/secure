package controller

import (
	"../model"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"net/http"
)

type loginParams struct {
	Username string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var user model.User
	db.Where("username = ?", params.Username).First(&user)

	if &user == nil {
		panic("User does not exist")
	}

	err = user.CheckPassword(params.Password)
	if err != nil {
		panic(err)
	}

	db.Where("user_id = ?", user.ID).Delete(model.ApiKey{})

	key, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}

	apiKey := model.ApiKey{
		Key:    key.String(),
		UserID: user.ID,
	}

	db.Create(&apiKey)
	payload := make(map[string]string)
	payload["key"] = apiKey.Key
	payload["message"] = "LoginSuccessful"

	data, err := json.Marshal(payload)
	w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	user := model.User{Username: params.Username}
	user.CreatePassword(params.Password)
	db.Create(&user)

	payload := make(map[string]string)
	payload["username"] = user.Username
	payload["message"] = "User created"

	data, err := json.Marshal(payload)
	w.Write(data)
}
