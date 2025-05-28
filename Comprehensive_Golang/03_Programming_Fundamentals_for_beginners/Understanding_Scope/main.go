package main

/*
NOTE:
The imported package "fmt" is in the "file block" scope of this file
*/
import "fmt"

// NOTE: x is in "package block" scope
var x = 42

func printMe() {
	// NOTE: x can be accessed here
	fmt.Println(x)
}

// NOTE: type person is in "package block" scope
type person struct {
	first string
}

/*
NOTE:
The method sayHello which is atteched to VALUE of TYPE persion
is in "package block" scope
*/
func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}

func main() {
	// NOTE: x can be accessed here
	fmt.Println(x)

	// NOTE: x can be accessed here
	printMe()

	// NOTE: y doesn't exist here so this won't work
	// fmt.Println(y)

	// NOTE: declare y is in "block scope"
	y := 43
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello()

	// NOTE: variable "shadowing"
	x := 32
	fmt.Println(x)

	// NOTE: package block x is still the same
	printMe()

	// also good to know

	if z := 82; z > 50 {
		fmt.Println(z)
	}
	// NOTE: you can't access z here
	// fmt.Println(z)
}
