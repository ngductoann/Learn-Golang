package main

import "fmt"

func main() {
	/*
		////////////////////
		Printf for decimal & hexadecimal values
		///////////////////
	*/

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("%v \t %b \t %#X\n", a, a, a)
	fmt.Printf("%v \t %b \t %#X\n", b, b, b)
	fmt.Printf("%v \t %b \t %#X\n", c, c, c)
	fmt.Printf("%v \t %b \t %#X\n", d, d, d)
	fmt.Printf("%v \t %b \t %#X\n", e, e, e)
	fmt.Printf("%v \t %b \t %#X\n", f, f, f)
}
