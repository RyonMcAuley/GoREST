package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()
	v1.HandleFunc("/v1/hello", handleHello).Methods("GET")
	r.HandleFunc("/hello", handleHello).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	return r
}
