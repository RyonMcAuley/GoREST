package main

import (
	"fmt"
	"net/http"

	"github.com/RyonMcAuley/goREST/internal/handler"
)

func main() {
	r := handler.NewRouter()

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
