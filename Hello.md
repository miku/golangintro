# Hello World

The first program is simple:

```
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

Observations:

* code lives in a package
* the entry point is a function `main` within a package `main`

There is one notion, that is (has been) unusual about Go. All code lives in a workspace.

## Workspace

What is a workspace?

> A fixed directory structure where Go code lives.

The documentation says:

> A workspace is a directory hierarchy with two directories at its root:

> **src** contains Go source files, and
> **bin** contains executable commands.

The go tool builds and installs binaries to the bin directory.

The default location is `$HOME/go` on Linux and `%USERPROFILE%\go`, usually
`C:\Users\YourName\go` on Windows.


```shell
$ tree /home/tir/go/src/github.com/miku/golangintro
/home/tir/go/src/github.com/miku/golangintro
├── Editors.md
├── hello
│   └── hello.go
├── Hello.md
├── img
│   ├── gopher.png
│   ├── screenshot-vim-go-complete.png
│   └── screenshot-vim-go-goimports.png
├── Installation.md
├── LICENSE
├── Origins.md
└── README.md

2 directories, 10 files
```


