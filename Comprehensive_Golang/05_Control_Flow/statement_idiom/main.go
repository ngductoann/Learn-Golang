package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		NOTE: The comma ok idiom is also like this
	*/

	x := 40

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
