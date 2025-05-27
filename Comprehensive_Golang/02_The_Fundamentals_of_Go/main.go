package main

import "fmt"

func main() {
	// Declare a variable, value variable is zero value (document/zero values)
	var g int
	fmt.Println(g)

	// Assign value to variable
	g = 43
	fmt.Println(g)

	// declare a variable to hold a VALUE of a certain TYPE
	// initializing a variable
	var h int = 44
	fmt.Println(h)

	// Or short declare and assign value to variable
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	// this would not work (e variable is unused)
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d)
	*/
}
