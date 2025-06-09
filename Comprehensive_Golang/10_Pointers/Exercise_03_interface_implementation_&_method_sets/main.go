package main

import "fmt"

type dog struct {
	first string
}

type youngin interface {
	walk()
	run()
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I am running")
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)
}
