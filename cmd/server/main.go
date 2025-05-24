package main

import (
	"fmt"
	"net/http"

	"github.com/RyonMcAuley/goREST/internal/customer"
	"github.com/RyonMcAuley/goREST/internal/purchase"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	customerRouter := r.PathPrefix("/customer").Subrouter()
	purchaseRouter := r.PathPrefix("/purchase").Subrouter()

	customer.RegisterRoutes(customerRouter)
	purchase.RegisterRoutes(purchaseRouter)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
