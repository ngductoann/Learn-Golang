package main

import "fmt"

// Using func literals, aka, anonymous self-executing func

func main() {
	c := make(chan int)

	// a buffered channel
	// c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
