package main

import "fmt"

func main() {
	n := 10
	c := make(chan int, n)
	done := make(chan bool)

	go func() {
		for i := range n {
			c <- i
		}
		close(c) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	for range n {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true // Signal that this goroutine is done
			// This is important to prevent deadlocks when pulling from the channel
			// Each goroutine will signal when it has finished processing
		}()
	}

	for range n {
		<-done // Wait for all goroutines to finish
	}
}
