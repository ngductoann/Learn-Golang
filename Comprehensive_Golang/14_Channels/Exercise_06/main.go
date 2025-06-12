package main

import "fmt"

/*
write a program that
1. puts 100 numbers to a channel
2. pull the numbers off the channel and print them
*/

func main() {
	c := make(chan int)

	go func() {
		for i := range 100 {
			c <- i
		}
		close(c) // Close the channel after sending all numbers
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
