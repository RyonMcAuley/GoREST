package customer

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Customer!")
	w.Write([]byte("Hello, Customer!"))
}
