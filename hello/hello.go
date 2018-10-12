package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("args: %s\n", os.Args)
	os.Exit(1)
}
