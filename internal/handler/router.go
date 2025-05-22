package handler

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// api := r.PathPrefix("/api").Subrouter()
	r.HandleFunc("/hello", handleHello).Methods("GET")

	return r
}
