package main

import "fmt"

func main() {
	// Uncommenting the following lines will cause a compile-time error
	// cs := make(chan<- int) // Creating a send-only channel this will not allow receiving
	// cs := make(<-chan int) // Creating a receive-only channel this will not allow sending
	// fix
	cs := make(chan int) // this line will allow both sending and receiving

	go func() {
		cs <- 42 // Sending to a send-only channel
	}()

	fmt.Println(<-cs) // This will cause a compile-time error

	fmt.Println("----------")
	fmt.Printf("cs\t%T\n", cs)
}
