package purchase

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Handler struct {
	database *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{database: db}
}

func (h *Handler) handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Purchase!")
	w.Write([]byte("Hello, Purchase!"))
}
