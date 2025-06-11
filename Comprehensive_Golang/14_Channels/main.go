package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := range 10 {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	/*
		don't work
		time.Sleep(time.Second)

		ca := make(chan int)
		ca <- 42
		fmt.Println(<-ca)
	*/

	c2 := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c2)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific to general doesn't assign
	// c = cr
	// c = cs

	// specific to specific doesn't assign
	// cs = cr

	// general to specific assigns
	cr = c2
	cs = c2

	// general to specific converts
	fmt.Println("-----")
	fmt.Printf("c\t%T\n", (<-chan int)(c2))
	fmt.Printf("c\t%T\n", (chan<- int)(c2))

	// specific to general doesn't convert
	// fmt.Println("-----")
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan int)(cs))
}
