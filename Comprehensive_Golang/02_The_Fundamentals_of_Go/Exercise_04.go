package main

import "fmt"

var i, j int = 1, 2

func main() {
	// Variables with initializers
	/*
		NOTE:
		- A var declaration can include initializers, one per variable
		- If an initializer is present, the type can be omitted; the variable will take the type of the initializer
	*/

	// Short variable declarations
	/*
		NOTE:
		- Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type
		- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not variable
	*/
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)

	fmt.Printf("%T \t %T \t %T \t %T \t %T \n", i, j, c, python, java)
}
