package main

import (
	"fmt"
	"math"
	"time"
)

var x = func() {
	for i := range 10 {
		fmt.Println(i)
	}
}

func outer() func() int {
	return func() int {
		return 42
	}
}

func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

func square(n int) int {
	return n * n
}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}

func doWork() {
	for i := range 2_000 {
		fmt.Println(i)
	}
}

func timeFunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	// anonymous function
	func() {
		for i := range 10 {
			fmt.Println(i)
		}
	}()

	// func expression
	x()

	y := func() {
		for i := range 10 {
			fmt.Println(i)
		}
	}

	y()

	// func return
	f := outer()
	fmt.Println(f())

	// callback
	fmt.Println(printSquare(square, 3))

	// closure
	x2 := powinator(2)
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())
	fmt.Println(x2())

	// wrapper
	timeFunc(doWork)
}
