package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(*tree.Tree, chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
	walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failure.")
	}

	if !Same(tree.New(1), tree.New(2)) {
		fmt.Println("Success.")
	} else {
		fmt.Println("Failure.")
	}
}
