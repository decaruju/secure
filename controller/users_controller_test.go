package controller

import (
	"bytes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	rr := mockCall("/users", "POST", "{'username': 'juju', 'password': '12345'}", t)

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
