package main

import "fmt"

// This code pull the values off the channel using A SELECT STATEMENT
func main() {
	q := make(chan int) // q is a send-only channel
	c := gen(q)         // gen returns a receive-only channel

	receive(c, q) // receive function receives from the receive-only channel

	fmt.Println("about to exit")
}

// receive function receives from the receive-only channel
func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v) // Receiving from a receive-only channel
		case <-q:
			return // Exit the loop when a signal is received from q
		}
	}
}

// q is a send-only channel
// gen returns a send-only channel
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}
		q <- 1
		close(c) // Close the channel when done sending
	}()

	return c
}
