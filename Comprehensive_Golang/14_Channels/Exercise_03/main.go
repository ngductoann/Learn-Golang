package main

import "fmt"

// This code pull the values off the channel using A FOR RANGE LOOP

func main() {
	c := gen() // gen returns a receive-only channel

	// send function receives from the receive-only channel
	send(c)

	// Uncommenting the following line will cause a compile-time error
	// c <- 42 // This will not work because c is a receive-only channel

	fmt.Println("----------")
	fmt.Printf("c\t%T\n", c) // This will print the type of the channel
}

func send(c <-chan int) {
	for v := range c {
		fmt.Println(v) // Receiving from a receive-only channel
	}
}

func gen() <-chan int { // gen returns a receive-only channel
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}
		close(c)
	}()

	// limit access channel to receive-only
	// outer only can receive from this channel
	return c
}
