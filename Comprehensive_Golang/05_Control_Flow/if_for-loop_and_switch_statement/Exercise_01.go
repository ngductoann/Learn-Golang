package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("between 201 and 250")
	} else {
		fmt.Println("this was more than 250")
	}

	switch {
	case x <= 100:
		fmt.Println("less than 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("this was more than 250")
	}

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("iteration %v \t x and y are %v and %v\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("both are less than four")
		case x > 6 && y > 6:
			fmt.Println("both are greater than six")
		case x >= 4 && x <= 6:
			fmt.Println("x is from 4 to 6 inclusive or both numbers")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous were met")
		}
	}

	y := 0
	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop", y)
		y++
	}

	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Printf("index %v \t value %v\n", i, v)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for k, v := range m {
		fmt.Printf("key %v \t value %v\n", k, v)
	}
}
