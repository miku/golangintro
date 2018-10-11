# Concurrency

## Exercise: Equivalent Binary Trees (8/11)

```go
package main

import "golang.org/x/tour/tree"
import "fmt"

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    var walker func(t *tree.Tree)
    walker = func (t *tree.Tree) {
        if (t == nil) {
            return
        }
        walker(t.Left)
        ch <- t.Value
        walker(t.Right)
    }
    walker(t)
    close(ch)
}
// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
	same := Same(tree.New(1), tree.New(2))
	fmt.Println(same)
}
```

## Exercise: Web Crawler (10/11)

In this exercise you'll use Go's concurrency features to parallelize a web
crawler.

