package main

import "fmt"

/*
NOTE:
- one variable to store a value of type int8. Assign to it the largest number possible, the print it
- one variable to store a value of type uint8. Assign to it the largest number possible, the print it
*/

func main() {
	var x uint8 = 255
	var y int8 = 127

	fmt.Printf("%v is of type %T", x, x)
	fmt.Printf("%v is of type %T", y, y)
}
