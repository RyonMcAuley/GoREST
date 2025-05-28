package customer

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, database *sql.DB) {
	h := NewHandler(database)

	r.HandleFunc("/hello", h.handleHello).Methods("GET")
	r.HandleFunc("/customer", h.handleGetCustomer).Methods("GET")
	r.HandleFunc("/customer", h.handleCreateCustomer).Methods("POST")
}
