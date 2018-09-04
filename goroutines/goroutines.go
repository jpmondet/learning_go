package main

import "fmt"

func parallhello(i int, c chan<- int) {
	fmt.Printf("I am goroutine number %d\n", i)
	c <- i
}

func main() {
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go parallhello(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("I'm receiving ", <-c)
	}
}
