package main

import "fmt"

/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results
*/

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "James Bond"
}

func main() {
	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)
}
