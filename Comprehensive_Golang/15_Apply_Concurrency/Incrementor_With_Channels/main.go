package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("Final Counter:", <-c3+<-c4)
}

func incrementor(s string) chan int {
	out := make(chan int)

	go func() {
		for i := range 20 {
			out <- i
			fmt.Println(s, i)
		}
		close(out) // Close the channel to signal that no more values will be sent
		// This is important to prevent deadlocks when pulling from the channel
	}()

	return out
}

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		for i := range c {
			out <- i
		}
		close(out) // Close the output channel to signal that no more values will be sent
		// This is important to prevent deadlocks when pulling from the channel
	}()

	return out
}
