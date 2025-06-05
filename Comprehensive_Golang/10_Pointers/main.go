package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // Print the memory address of x

	fmt.Printf("%v\t%T\n", &x, &x) // address variable and ype

	s := "James"
	fmt.Printf("%v\t%T\n", &s, &s) // address variable and ype

	y := &x
	fmt.Printf("%v\t%T\n", y, y) // address variable and ype
	fmt.Println(*y)              // value in address
	fmt.Println(*&x)             // value in address

	*y = 43
	fmt.Println(x)  // value
	fmt.Println(*y) // value
}
