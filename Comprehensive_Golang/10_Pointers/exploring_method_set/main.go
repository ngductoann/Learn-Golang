package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I am walking")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I am running")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{first: "Henry"}
	d1.walk() // Calls the value receiver method
	d1.run()  // Calls the pointer receiver method
	// youngRun(d1) // doesn't run because d1 is a value type and run requires a pointer receiver

	d2 := &dog{first: "Padget"}
	d2.walk()    // Calls the value receiver method
	d2.run()     // Calls the pointer receiver method
	youngRun(d2) // Calls the pointer receiver method through the interface
}
