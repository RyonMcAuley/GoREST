package handler

import "net/http"

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
