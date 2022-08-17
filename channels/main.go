package main

import "fmt"

func main() {
	c := make(chan int)
	c2 := make(chan int, 1)

	go func() {
		c <- 42
	}()
	c2 <- 32

	fmt.Println(<-c)
	fmt.Println(<-c2)

	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send
	fmt.Println(cr)

	go func() {
		cs <- 42
	}()

	cFunc := make(chan int)
	go foo(cFunc)
	bar(cFunc)

	for v := range cFunc {
		fmt.Println(v)
	}

}

//send
func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- 77
	}
	close(c)
}

//receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

// we also have context. We can cancel the go routines...
