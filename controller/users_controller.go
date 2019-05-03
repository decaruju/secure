package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
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

func UsersRouter(router *mux.Router) {
	s := router.PathPrefix("/users").Subrouter()
	s.HandleFunc("/login", Cors(Login)).Methods("POST")
	s.HandleFunc("", Cors(CreateUser)).Methods("POST")
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var params loginParams
	err := decoder.Decode(&params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := db()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user model.User
	db.Where("username = ?", params.Username).First(&user)

	if &user == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = user.CheckPassword(params.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	db.Where("user_id = ?", user.ID).Delete(model.ApiKey{})

	key, err := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiKey := model.ApiKey{
		Key:    key.String(),
		UserID: user.ID,
	}

	db.Create(&apiKey)
	payload := make(map[string]string)
	payload["key"] = apiKey.Key

	data, err := json.Marshal(payload)
	w.Write(data)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var params loginParams
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &params)
	if err != nil || params.Username == "" || params.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, err := db()
	defer db.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := model.User{Username: params.Username}
	user.CreatePassword(params.Password)
	db.Create(&user)

	payload, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

func db() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	return db, err
}
