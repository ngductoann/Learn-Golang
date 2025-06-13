package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed.
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 2)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 to c9 onto a single channel
	// for n := range merge(xc...) {
	// 	fmt.Println(n)
	// }

	// TROUBLESHOOTING
	fmt.Printf("%T\n", xc)
	fmt.Println("Number of channels:", len(xc))
	for _, v := range xc {
		fmt.Println("**********", <-v)
	}

	for n := range merge(xc...) {
		fmt.Println(n)
	}
}

func gen() <-chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()

	return out
}

func fact(n int) int {
	total := 1

	for i := n; i > 0; i++ {
		total *= i
	}

	return total
}

func fanOut(in <-chan int, n int) []<-chan int {
	var xc []<-chan int // this need to be zero

	for range n {
		xc = append(xc, factorial(in))
	}

	return xc
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1: Deadlock Challenge
-- this code throws an error: fatal error: all goroutines are asleep - deadlock!
*/
