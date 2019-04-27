package controller

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"secure/model"
)

type loginParams struct {
	Username string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		panic(err)
	}

	db := db()
	defer db.Close()

	var user model.User
	db.Where("username = ?", params.Username).First(&user)

	if &user == nil {
		panic("User does not exist")
	}

	err = user.CheckPassword(params.Password)
	if err != nil {
		w.Write([]byte("Wrong password"))
		return
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
	var params loginParams
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &params)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("Error"))
		return
	}

	db := db()
	defer db.Close()

	user := model.User{Username: params.Username}
	user.CreatePassword(params.Password)
	db.Create(&user)

	payload, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(payload)
}

func db() *gorm.DB {
	db, _ := gorm.Open("sqlite3", "test.db")
	return db
}
