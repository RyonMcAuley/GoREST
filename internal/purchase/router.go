package purchase

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello", handleHello).Methods("GET")
}
