package main

import "fmt"

/*
Take the code from the previous exercise, then store the VALUES of type person in a map with the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
*/

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.favIC {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
