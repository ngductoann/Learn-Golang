package main

import "fmt"

func main() {
	c := make(chan int) // Create a bidirectional channel

	go func() {
		c <- 42
	}()

	// comma-ok idiom to receive from the channel
	v, ok := <-c       // Receive from the channel
	fmt.Println(v, ok) // Print the value and the ok status

	close(c)

	v, ok = <-c
	fmt.Println(v, ok) // Print the value and the ok status after closing the channel
}
