package main

import "fmt"

func main() {
	in := gen()

	f := factorial(in)

	for n := range f {
		fmt.Println(n)
	}
}

// gen generates a sequence of integers from 3 to 12, ten times.
// It uses a goroutine to send values concurrently.
// The input channel is read-only (receive-only).
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

// factorial receives a channel of integers and returns a channel of their factorials
// It uses a goroutine to compute the factorials concurrently.
// The input channel is read-only (receive-only), and the output channel is also read-only(receive-only).
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
	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
