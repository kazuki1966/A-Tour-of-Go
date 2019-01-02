package main

import "fmt"

func calcFibonacci(a0, a1 *int) {
	*a0, *a1 = *a1, *a0+*a1
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a0, a1 := 0, 1

	return func() int {
		defer calcFibonacci(&a0, &a1)
		return a0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
