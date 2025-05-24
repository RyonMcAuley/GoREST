package purchase

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Purchase!")
	w.Write([]byte("Hello, Purchase!"))
}
