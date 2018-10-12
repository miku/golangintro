// A hello world server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var count int64
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++ // or: atomic.AddInt64(&count, 1)
		fmt.Fprintf(w, "Hello World, #%d", count)
	})
	log.Println("http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
