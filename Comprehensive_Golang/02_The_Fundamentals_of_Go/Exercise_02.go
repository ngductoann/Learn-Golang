package main

import "fmt"

func add(x, y int) int { // or x int, y int
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	// Function
	/*
		NOTE:
		- A function can take zero or more arguments.
		- In this example, add takes two parameters or type int.
		- Notice that the type comes after the variable name
	*/

	fmt.Println(add(42, 13))
	sayHello()

	// Multiple results
	/*
		NOTE:
		- A function can return any number of results
		- The swap function returns two string
	*/
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return value
	/*
		NOTE:
		- Go's return values may be named. If so, they are treated as variables defined at the top of the function
		- These names should be used to document the meaning of the return values.
		- A return statement without arguments returns the named return values. This is known as a "naked" return
		- Naked return statements should be used only in short function, as with the example shown here. They can harm readability in logger functions
	*/
	fmt.Println(split(17))
}

func sayHello() {
	fmt.Println("hello")
}
