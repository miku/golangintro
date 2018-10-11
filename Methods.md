# Methods

* http://127.0.0.1:3999/methods/1, https://tour.golang.org/methods/1

26 sections.

> Go does not have classes. However, you can define methods on types. A method
> is a function with a special receiver argument.

## Pointer receivers (4/26)

> Methods with pointer receivers can modify the value to which the receiver
> points (as Scale does here).

## Pointers and functions (5/26)

Non-pointer, do not take address.

```
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(v, 10)
	fmt.Println(Abs(v))
}
```

## Interfaces (9/26)

> An interface type is defined as a set of method signatures.

In line 22, write `a = &v`.

## Interfaces are implemented implicitly (10/26)

> Implicit interfaces decouple the definition of an interface from its
> implementation, which could then appear in any package without prearrangement.

## The empty interface (14/26)

> Empty interfaces are used by code that handles values of unknown type. For
> example, fmt.Print takes any number of arguments of type interface{}. 

