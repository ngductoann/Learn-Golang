package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42

	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("%v\t%T\n", a, a)
	fmt.Printf("%v\t%T\n", b, b)
	fmt.Printf("%v\t%T\n", c, c)
	fmt.Printf("%v\t%T\n", d, d)

	fmt.Println("Value of a:", *a)
	fmt.Println("Value of b:", *b)
	fmt.Println("Value of c:", *c)
	fmt.Println("Value of d:", *d)
}
