package main

import "fmt"

func foo() {
	fmt.Println("Foo ran")
}

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name", s)
	}

	x()
	y("Todd")
}
