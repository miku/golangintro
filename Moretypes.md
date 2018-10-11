# More types

* https://tour.golang.org/moretypes/1

Contains 27 sections.

## Pointers (1/27)

* holds memory address
* no pointer arithmetic

## Struct literals (5/27)

* no mixed initialization allowed

## Arrays (6/27)

* an array defines an own type, e.g. `[10]int`
* https://play.golang.org/p/wViv2k97aty

## Slice length and capacity (11/27)

* Will an unreachable (via slicing) array part be garbage collected?

## Creating a slice with make (13/27)

* http://localhost:6060/ref/spec#Making_slices_maps_and_channels

## Exercise: Slices (18/27)

Example solution:

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var result [][]uint8
	for y := 0; y < dy; y++ {
		var row[] uint8
		for x := 0; x < dx; x++ {
			row = append(row, uint8((x + y) / 2))
		}
		result = append(result, row)
	}
	return result
}

func main() {
	pic.Show(Pic)
}
```

Another solution:

```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	yy := make([][]uint8, dx)
	for y := 0; y < dy; y++ {
		yy[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			yy[y][x] = uint8((x + y) / 2)
		}
	}
	return yy
}

func main() {
	pic.Show(Pic)
}
```

## Maps (19/27)

A missing string key is empty.

* https://play.golang.org/p/1AifJZYD6tS


## Exercise: Maps (23/27)

> Implement word count.

```go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	freq := make(map[string]int)
	for _, w := range strings.Fields(s) {
		freq[w] += 1
	}
	return freq
}

func main() {
	wc.Test(WordCount)
}
```

## Function values (24/27)

> Functions are values too. They can be passed around just like other values.

## Exercise: Fibonacci closure (26/27)

One solution:

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	pred := []int{0, 1}
	var value, i int
	
	return func() int {
		switch {
		case i < 2:
			value = i
		default:
			value = pred[1] + pred[0]
			pred[0] = pred[1]
			pred[1] = value
		}
		i++
		return value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

More elegant:

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

Even shorter:

```go
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x + y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
