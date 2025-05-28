package purchase

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, database *sql.DB) {
	h := NewHandler(database)

	r.HandleFunc("/hello", h.handleHello).Methods("GET")
}
