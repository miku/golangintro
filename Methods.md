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

```go
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

## Type assertions (15/26)

> If i does not hold a T, the statement will trigger a panic. 

## Exercise: Stringers (18/26)

One solution:

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## Exercise: Errors (20/26)

> Note: A call to fmt.Sprint(e) inside the Error method will send the program
> into an infinite loop. You can avoid this by converting e first:
> fmt.Sprint(float64(e)). Why?

```go
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
	// return fmt.Sprintf("cannot Sqrt negative number: %v", e) // fails
	// return fmt.Sprint(float64(e)) // works
	// return fmt.Sprint(e) // runtime: goroutine stack exceeds 1000000000-byte limit
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

## Exercise: Readers (22/26)

```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}
	return len(p), nil
}

func main() {
	reader.Validate(MyReader{})
}
```

## Exercise: rot13Reader (23/26)

One solution:

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}
	for i := 0; i < len(p); i++ {
		switch {
		case p[i] >= 'A' && p[i] < 'N' || p[i] >= 'a' && p[i] < 'n':
			p[i] = p[i] + 13
		case p[i] >= 'N' && p[i] <= 'Z' || p[i] >= 'n' && p[i] <= 'z':
			p[i] = p[i] - 13
		}
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

## Exercise: Images (25/26)

```go
package main

import "golang.org/x/tour/pic"
import "image/color"
import "image"

type Image struct {
	w int
	h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}
func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func Pic(dx, dy int) Image {
	return Image{dx, dy}
}

func main() {
	m := Image{w: 256, h: 256}
	pic.ShowImage(m)
}
```
