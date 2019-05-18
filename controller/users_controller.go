package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"secure/repository"
)

type loginParams struct {
	Username string
	Password string
}

func UsersRouter(router *mux.Router) {
	s := router.PathPrefix("/users").Subrouter()
	s.HandleFunc("/login", Cors(Login)).Methods("POST")
	s.HandleFunc("/logout", Cors(Logout)).Methods("POST")
	s.HandleFunc("", Cors(CreateUser)).Methods("POST")
}

func Login(w http.ResponseWriter, r *http.Request) {
	params, err := parseRequest(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := repository.FindUserByUsername(params.Username)
	if err != nil || user.CheckPassword(params.Password) != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	apikey, err := repository.CreateApikey(user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	payload, err := json.Marshal(apikey)
	w.Write(payload)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	key := r.Header.Get("Apikey")
	fmt.Println(key)

	user, err := repository.FindUserByKey(key)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = repository.DeleteAllApikeys(user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	params, err := parseRequest(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user, err := repository.CreateUser(params.Username, params.Password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	payload, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(payload)
}

func parseRequest(body io.ReadCloser) (loginParams, error) {
	params := loginParams{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return params, err
	}

	err = json.Unmarshal(data, &params)
	if err != nil || params.Username == "" || params.Password == "" {
		return params, ErrInvalidParams
	}
	return params, nil
}
