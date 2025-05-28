package main

import (
	"fmt"
	"net/http"

	"github.com/RyonMcAuley/goREST/internal/customer"
	"github.com/RyonMcAuley/goREST/internal/platform/db"
	"github.com/RyonMcAuley/goREST/internal/purchase"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	database := db.Init()
	defer database.Close()

	customerRouter := r.PathPrefix("/customer").Subrouter()
	purchaseRouter := r.PathPrefix("/purchase").Subrouter()

	customer.RegisterRoutes(customerRouter, database)
	purchase.RegisterRoutes(purchaseRouter, database)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
