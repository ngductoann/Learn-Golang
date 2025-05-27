package main

import (
	"fmt"
	"math"
)

func main() {
	// exported names
	/*
		NOTE:
		In Go, a name is exported if it begins with a capital lettel. For example, Pizza is an exported name, as Pi, which is exported from the math package
	*/
	fmt.Println(math.Pi)
}
