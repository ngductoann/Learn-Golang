package main

import (
	"fmt"
	"slices"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	// x = append(x[:3], x[6:]...)
	// or
	x = slices.Delete(x, 3, 6)
	fmt.Println(x)
}
