package main

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := range 10 {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := range 10 {
			c <- i
		}
		done <- true
	}()

	// We block here until done <- true
	<-done
	<-done
	close(c)

	// to unlock above
	// we need to take values off chan c here
	// but we never get here, because we're block above
	for n := range c {
		println(n)
	}
}
