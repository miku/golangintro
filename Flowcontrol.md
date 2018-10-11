# Flowcontrol

* https://tour.golang.org/flowcontrol/1 (14 sections)

## Exercise: Loops and functions

* https://tour.golang.org/flowcontrol/8

## Solution

```
package main

import (
	"fmt"
	"math/rand"
)

func Sqrt(x float64) float64 {
	z := float64(rand.Intn(10))
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
```