package controller

import "net/http"

func Cors(wrapped http.HandlerFunc) http.HandlerFunc {
	inner := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		wrapped(w, r)
	}

	return inner
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
