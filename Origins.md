## Golang Origins

### A timeline of events

* [The Go programming language](https://www.youtube.com/watch?v=rKnDgT73v8s), November 2009
* [Go version 1 is released](https://blog.golang.org/go-version-1-is-released), 28 March 2012

The Go 1 release marked an important milestone, which is codified in [Go 1 and
the Future of Go Programs](https://golang.org/doc/go1compat):

> Go 1 defines two things: first, the specification of the language; and
> second, the specification of a set of core APIs, the "standard packages" of
> the Go library. The Go 1 release includes their implementation in the form
> of two compiler suites (gc and gccgo), and the core libraries themselves.

> It is intended that programs written to the Go 1 specification will continue
> to compile and run correctly, unchanged, over the lifetime of that
> specification. 

In Spring 2016, Brad Fitzgerald gave a talk: [Go 1.6: Asymptotically
approaching boring](https://www.youtube.com/watch?v=4Dr8FXs9aJM)
([slides](https://docs.google.com/presentation/d/1JsCKdK_AvDdn8EkummMNvpo7ntqteWQfynq9hFTCkhQ/view?slide=id.p)).

Go has releases approximately every six months (seems attractive to
[others](https://www.infoq.com/news/2017/09/Java6Month), too).

Rationale: Stable foundation - to build stuff on top.

* Language ([ref/spec](https://golang.org/ref/spec))
* Standard Library ([pkg](https://golang.org/pkg/))
* Runtime ([GC](https://www.youtube.com/watch?v=aiv1JOfMjm0), scheduler, and
  other pieces under active development)
* Tools (go, godoc, go vet, gofmt, goimports, ... under active development)
* Ecosystem (external packages, conferences, user groups, and much more)

The current release is [Go 1.11](https://blog.golang.org/go1.11) from 24 August
2018.

### Origins

A multitude of ideas and influences:

* C++ might be slow to compile and bloated
* all programming languages seem to add and add features
* we are entered the multicore era
* design for a networked world
* bridge the gap between dynamic and static programming languages (be safe, yet
  ease to write)
* focus on long-term maintenance
* designed with tools in mind (gofmt being the prototypical tool)
* a different approach to concurrency
* a stripped version of object orientation

The original designers are Robert Griesemer, Rob Pike, Ken Thomson.
