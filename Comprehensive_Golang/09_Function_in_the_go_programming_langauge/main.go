package main

import (
	"fmt"
)

/*
func syntax

no params, no returns
1 param, no returns
1 param, 1 return
2 params, 2 returns
*/

func foo() {
	fmt.Println("I am from foo")
}

func bar(s string) {
	fmt.Println("My name is", s)
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr. ", s)
}

func dogYears(name string, age int) (string, int) {
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r receiver) identifier(p parameters) (return(s)) {code}
func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)

	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("I am", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	foo()

	bar("Todd")

	s := aloha("Todd")
	fmt.Println(s)

	s1, n := dogYears("Todd", 40)
	fmt.Println(s1, n)

	x := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The sum is", x)

	// f, err := os.Open("./Document.md")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// defer f.Close()

	// continue to work with file

	p1 := person{
		first: "James",
	}

	p2 := person{
		first: "Jenny",
	}

	p1.speak()
	p2.speak()

	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	saySomething(sa1)
	saySomething(p2)
}
