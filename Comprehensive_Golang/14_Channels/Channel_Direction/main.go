package main

import "fmt"

func main() {
	c := increment() // Call increment to get a channel that will send integers

	for n := range puller(c) { // Use range to read from the channel until it is closed
		fmt.Println(n)
	}
}

func increment() <-chan int { // Use <-chan to indicate that this channel is read-only
	out := make(chan int)

	go func() {
		for i := range 10 {
			out <- i
		}
		close(out) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	return out
}

func puller(c <-chan int) <-chan int { // Use <-chan to indicate that this channel is read-only
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	return out
}
