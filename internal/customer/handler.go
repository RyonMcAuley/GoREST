package customer

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

type Handler struct {
	database *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{database: db}
}

func (h *Handler) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Customer!")
	w.Write([]byte("Hello, Customer!"))
}

func (h *Handler) handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	// Example implementation for getting a customer
	customerID := r.URL.Query().Get("id")
	if customerID == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// Here you would typically query the database for the customer
	fmt.Fprintf(w, "Getting customer with ID: %s", customerID)
	idNum, _ := strconv.Atoi(customerID)
	cust, err := GetCustomerByID(h.database, idNum)
	if err != nil {
		fmt.Fprintf(w, "Error retrieving customer: %v", err)
	}

	fmt.Fprintf(w, "Customer: %+v", cust)

}
