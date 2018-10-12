// A hello world server.
package main

import (
	"io"
	"log"
	"net/http"
)

// Handler is a basic handler.
type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	http.Handle("/", Handler{})
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
