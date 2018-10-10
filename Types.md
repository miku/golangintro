# Types

Go is statically typed. It has a number of basic data types and a few container
types.

## Numeric types

> A numeric type represents sets of integer or floating-point values. The
> predeclared architecture-independent numeric types are:

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

* [ref/spec](https://golang.org/ref/spec#Numeric_types* [ref/spec](https://golang.org/ref/spec#Numeric_types)

## String types

> A possibly empty sequence of bytes.

> In Go, a string is in effect a read-only slice of bytes. It's important to
> state right up front that a string holds arbitrary bytes. It is not required
> to hold Unicode text, UTF-8 text, or any other predefined format. It's
> important to state right up front that a string holds arbitrary bytes. It is
> not required to hold Unicode text, UTF-8 text, or any other predefined
> format.

* https://play.golang.org/p/emodHMMaRzc

```go
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    x := "Äther"
    fmt.Println(x)
    fmt.Println(len(x))
    fmt.Println(utf8.RuneCountInString(x))
    fmt.Println(utf8.RuneLen('Ä'))
    fmt.Printf("%x %x\n", x[0], x[1])
    fmt.Printf("%+q\n", x[0])
    fmt.Printf("%+q\n", x[1])

    y := `Multiline
    strings
    use backticks.
    `

    fmt.Println(y)
}
```

##
