// Sequential link checker emitting JSON.
package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type result struct {
	Link       string `json:"link,omitempty"`
	StatusCode int    `json:"code,omitempty"`
	Error      error  `json:"error,omitempty"`
}

func (r result) WriteTo(w io.Writer) (int64, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}
	b = append(b, '\n')
	n, err := w.Write(b)
	return int64(n), err
}

func main() {
	f, err := os.Open("small.txt")
	if err != nil {
		log.Fatal(err)
	}
	br := bufio.NewReader(f)

	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		link := strings.TrimSpace(line)
		resp, err := http.Get(link)
		if err != nil {
			r := result{Link: link, Error: err}
			if _, err := r.WriteTo(os.Stdout); err != nil {
				log.Fatal(err)
			}
			continue
		}
		defer resp.Body.Close()
		r := result{Link: link, StatusCode: resp.StatusCode}
		if _, err := r.WriteTo(os.Stdout); err != nil {
			log.Fatal(err)
		}
	}
}
