package main

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for range n {
		go func() {
			for j := range 10 {
				c <- j
			}
			done <- true
		}()
	}

	go func() {
		for range n {
			<-done
		}
	}()

	for n := range c {
		println(n)
	}
}
