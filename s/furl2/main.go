// A sequential link checker that is reading from a file.
package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

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
			log.Printf("failed %s: %s", link, err)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode < 400 {
			log.Printf("ok[%d]: %s", resp.StatusCode, link)
		} else {
			log.Printf("failed[%d]: %s", resp.StatusCode, link)
		}
	}
}
