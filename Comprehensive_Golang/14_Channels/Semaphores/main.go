package main

func main() {
	c := make(chan int)
	done := make(chan bool) // Channel to signal when goroutines are done

	go func() {
		for i := range 10 {
			c <- i
		}
		done <- true // Signal that this goroutine is done
	}()

	go func() {
		for i := range 10 {
			c <- i
		}
		done <- true // Signal that this goroutine is done
	}()

	go func() {
		<-done   // Wait for the first goroutine to finish
		<-done   // Wait for the second goroutine to finish
		close(c) // Close the channel to signal completion
	}()

	for n := range c {
		println(n)
	}
}
