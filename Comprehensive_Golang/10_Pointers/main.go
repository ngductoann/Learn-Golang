package main

import "fmt"

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // Print the memory address of x

	fmt.Printf("%v\t%T\n", &x, &x) // address variable and ype

	s := "James"
	fmt.Printf("%v\t%T\n", &s, &s) // address variable and ype

	y := &x
	fmt.Printf("%v\t%T\n", y, y) // address variable and ype
	fmt.Println(*y)              // value in address
	fmt.Println(*&x)             // value in address

	*y = 43
	fmt.Println(x)  // value
	fmt.Println(*y) // value

	fmt.Println("---------------------")
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])

	fmt.Println("---------------------")
	// value semantics
	a2 := 1
	fmt.Println(a2)         // 1
	fmt.Println(addOne(a2)) // 2, value semantics
	fmt.Println(a2)         // 1, original value remains unchanged

	// pointer semantics
	b := 1
	fmt.Println(b) // 1
	addOneP(&b)
	fmt.Println(b) // 2, original value is changed.

	/*
		escapes to the heap
		variable shared between main() and Println()

		moved to heap
		variable moved to the heap
	*/
}
