package main

import "fmt"

var c, python, java bool

func main() {
	// Variables
	/*
		NOTE:
		- The var statement declares a list of variables; as in function arguments list, the type list
		- A var statement can be at package or function level. We see both in this example
	*/
	var i int
	fmt.Println(i, c, python, java)
}
