# Reader, writers and the io package

IO zoo: https://github.com/miku/exploreio

## Exercise: Implement a simple byte counter satisfying the io.Writer interface

```go
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

var c ByteCounter
c.Write([]byte("hello"))
fmt.Println(c)

// Reset.
c = 0
var name = "Dolly"
fmt.Fprintf(&c, "hello, %s", name)
fmt.Println(c) // "12"
```

* Example: `chainr.go`