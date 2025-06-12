package main

func main() {
	c := incrementor()
	cSum := puller(c)

	for n := range cSum {
		println(n)
	}
}

func incrementor() chan int {
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

func puller(c chan int) chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		close(out) // Close the channel to signal completion
		// This is important to prevent deadlocks when pulling from the channel
	}()

	return out
}
