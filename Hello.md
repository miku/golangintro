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

The go tool builds and installs binaries to the bin directory (there is also
`pkg`, contains Go package objects, might go away).

The default location (controlled by GOPATH environment variable) is `$HOME/go`
on Linux and `%USERPROFILE%\go`, usually `C:\Users\YourName\go` on Windows.


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

Since the directory layout is fixed, tools can assume a structure, to download
libraries, compile code, find dependencies. It can be confusing at first.

Dissecting the path:

```
/home/tir/go/src/github.com/miku/golangintro/hello/

--------  -- --- ----------------------------------
   |      |   |               |
   |      |   |           import path
 $HOME    |  src
          |
          |
        GOPATH

    $ go env GOPATH
```


