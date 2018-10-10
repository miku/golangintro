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

Over time, your GOPATH gets populated with code that you write and libraries
that you use. We will talk about versioning and vendoring later.

More examples. Here is an example top level `$HOME/go/src` directory:

```
$ find ~/go/src -maxdepth 1
/home/tir/go/src
/home/tir/go/src/gonum.org
/home/tir/go/src/gopkg.in
/home/tir/go/src/github.com
/home/tir/go/src/perkeep.org
/home/tir/go/src/honnef.co
/home/tir/go/src/golang.org
/home/tir/go/src/mvdan.cc
/home/tir/go/src/code.gitea.io
/home/tir/go/src/launchpad.net
```

The import name can be used to fetch packages, you might have seen this notion
on GitHub or elsewhere.

```
$ go get github.com/miku/solrbulk/cmd/...
```

## Running code

The `go` program combines a couple of functions in a single tool. Among other
things, it can run go code:

```shell
$ go run hello/hello.go
Hello World
```

> Run compiles and runs the main package comprising the named Go source files.
A Go source file is defined to be a file ending in a literal ".go" suffix.

## Compiling a binary

The go tool can build a binary:

```shell
$ go build -o myprog hello/hello.go
$ ./myprog
Hello World
```

The binary is statically linked and contains the go runtime and is about 2MB in
size (about 240x larger than hello.c).

```shell
$ file myprog
myprog: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, with debug_info, not stripped
```





