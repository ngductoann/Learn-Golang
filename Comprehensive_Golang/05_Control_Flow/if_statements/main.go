package main

import "fmt"

func main() {
	// SEQUENCE
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")

	x := 40 // this is the third statement to run
	y := 5  // this is the four statement to run
	fmt.Printf("x=%v \ny=%v\n", x, y)

	// CONDITIONAL
	// if statement
	// switch statement

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than meaning of life")
	}

	/*
		NOTE:
		"if" statements specify the conditional execution of two branches according to the value of a boolean expression.
		If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/

	if x < 42 && y < 42 {
		// && requires both to be true to run
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || y < 42 {
		// || requires one of them to be true to run
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		// && requires both to be true to run
		fmt.Println("x is not 42 & y is not 10")
	}
}
