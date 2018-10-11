package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	fib, fib2 := 0, 0
	return func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			fib = 1
			return 1
		}
		fib = fib + fib2
		fib2 = fib - fib2
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
