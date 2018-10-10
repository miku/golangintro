// Converts an latin-1 encoded input to utf-8 encoded output.
package main

import (
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	if _, err := io.Copy(out, r); err != nil {
		log.Fatal(err)
	}

	out.Close()
	f.Close()
}
