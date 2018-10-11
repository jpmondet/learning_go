package main

import "fmt"

func fib(n int) (fn int) {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

func main() {

	res := 0
	for i := 0; i < 10; i++ {
		res = fib(i)
		fmt.Println(res)
	}

}
