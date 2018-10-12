// A hello world server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler is a basic handler.
type Handler struct {
	Name string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World (from %s)", h.Name)
}

func main() {
	http.Handle("/", Handler{Name: "serv1"})
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
