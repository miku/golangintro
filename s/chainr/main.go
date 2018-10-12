package main

import (
	"compress/gzip"
	"encoding/base64"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	r, _ = gzip.NewReader(r)
	io.Copy(os.Stdout, r)
}

// echo "Go" | gzip -c | base64
const data = `H4sIANMwwFsAA3PP5wIAIdwkiAMAAAA=`
