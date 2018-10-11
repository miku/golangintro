# Working with files

Realized by the os package.

## Opening a file

```go
filename := "hello.txt"
f, err := os.Open(filename)
if err != nil {
    log.Fatal(err)
}
defer f.Close()

if _, err := io.Copy(os.Stdout, f); err != nil {
    log.Fatal(err)
}
```

## Writing to a file

A file is (implements) `io.Writer`, but there are convenience methods available
in the `fmt` or `io` package.

```go
_, _ f.Write([]byte("Hello World"))
```

Using the `fmt` package:

```go
fmt.Fprintf(w, "Using Fprintf")
```

```go
_, _ = io.WriteString(w, "Hello World")
```

## Buffered IO

The [`bufio`](https://golang.org/pkg/bufio/) package provides methods for
buffered reads and writes.

```go
br := bufio.NewReader(os.Stdin)
for {
    line, err := br.ReadString('\n')
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
}
```
