package main

import "sync"

func main() {
	in := gen()

	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)
	c10 := factorial(in)

	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10) {
		println(n)
	}
}

// gen generates a sequence of integers from 3 to 12, repeated 10 times.
// It returns a channel receive only
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

// factorial computes the factorial of numbers received from the input channel.
// It returns a channel receive only
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

// fact computes the factorial of a given number n.
func fact(n int) int {
	total := 1

	for i := n; i > 1; i-- {
		total *= i
	}

	return total
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
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern to accomplish this.

CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
*/
