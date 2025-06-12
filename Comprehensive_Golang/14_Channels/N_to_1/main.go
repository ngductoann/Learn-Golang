package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2) // Two goroutines will send data to the channel

	go func() {
		for i := range 10 {
			c <- i
		}
		wg.Done() // Signal that this goroutine is done
		// Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	go func() {
		for i := range 10 {
			c <- i
		}
		wg.Done() // Signal that this goroutine is done
		// Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	go func() {
		wg.Wait() // Wait for both goroutines to finish
		/// Close the channel after all data has been sent
		close(c) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	for n := range c {
		fmt.Println(n)
	}
}
