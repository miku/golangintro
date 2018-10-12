// A hello world server.
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	})
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
