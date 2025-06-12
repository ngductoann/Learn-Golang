package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := range 10 {
			c <- i
		}
		close(c) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true // Signal that this goroutine is done
		// This is important to prevent deadlocks when pulling from the channel
	}()

	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true // Signal that this goroutine is done
		// This is important to prevent deadlocks when pulling from the channel
	}()

	<-done // Wait for the first goroutine to finish
	<-done // Wait for the second goroutine to finish
}
