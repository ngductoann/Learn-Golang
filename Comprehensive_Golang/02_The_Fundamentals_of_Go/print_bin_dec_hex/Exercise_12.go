package main

import "fmt"

/*
NOTE:
print verbs to show the following numbers: 747, 911, 90210
as decimal, binary, hexadecimal
*/

func main() {
	x, y, z := 747, 911, 90210

	fmt.Printf("%d \t\t %b \t\t %#X\n", x, x, x)
	fmt.Printf("%d \t\t %b \t\t %#X\n", y, y, y)
	fmt.Printf("%d \t\t %b \t\t %#X\n", z, z, z)
}
