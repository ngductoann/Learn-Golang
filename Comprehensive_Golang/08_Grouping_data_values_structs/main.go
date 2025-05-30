package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)

	fmt.Println(p1.first, p1.last, p1.age)

	sa1 := secretAgent{
		// person: person{
		// 	first: "James",
		// 	last:  "Bond",
		// 	age:   32,
		// },
		// or
		person: p1,
		first:  "ETHAN HAWK",
		ltk:    true,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
	fmt.Println(sa1.first, sa1.person.first)

	// Anonymous struct
	p3 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Printf("%T \n", p3)
}
