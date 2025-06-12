package main

import (
	"fmt"
	"runtime"
)

/*
Write a program that:
- launches 10 goroutines
	- each gorountine adds 10 numbers to a channel
- pull the numbers off the channel and print them
*/

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := range x * y {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) chan int {
	c := make(chan int)
	for range x {
		go func() {
			for j := range y {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}

	return c
}
