# Tools

* gofmt
* goimports
* go tool pprof, [example](https://github.com/miku/span/blob/master/docs/span-tag.0.1.135.png)

```
    ...
	if *cpuProfile != "" {
		file, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
    }
    ...
```

Or via debug runtime profiling, https://golang.org/pkg/net/http/pprof/.

* go vet
* linting: https://github.com/golangci/awesome-go-linters
* struct generators
* ...