package main

import (
	"fmt"
	"sort"
)

/*
NOTE:
- Arrays have theirs place, but they're a bit inflexible, so you don't see them to often in Go code.
- Slices, though, are everywhere. They build on arrays to provides great power and convenience.
- Slice is a data structure with three elements:
	-- len
	-- cap
	-- pointer to an underlying array

	type slice struct {
		array unsafe.Pointer
		len		int
		cap		int
	}
*/

func main() {
	// array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8
	fmt.Println(c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", c)

	{
		// declare a variable of type [7]int
		var ni [7]int
		// assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		// array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		// array literal
		// have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		// use the builtin len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))
	}

	// Slice literal
	xs := []string{"hello", "world"}
	fmt.Println(xs)

	xs2 := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)"}
	fmt.Println(xs2)
	fmt.Printf("%T\n", xs2)

	// blank identifier to not use a returned value
	for _, v := range xs2 {
		fmt.Printf("%v\n", v)
	}

	fmt.Println("-------------------")
	fmt.Println(xs2[0])
	fmt.Println(xs2[1])
	fmt.Println(xs2[2])

	// doesn't work
	// fmt.Println(xs[3])

	fmt.Println(len(xs2))

	for i := 0; i < len(xs2); i++ {
		fmt.Println(xs2[i])
	}

	// Append to slice
	fmt.Println("---------------------")
	xi := []int{42, 43, 44}
	fmt.Println(xi)
	fmt.Println("---------------------")

	// variadic parameter
	xi = append(xi, 45, 46, 47, 99, 777)
	fmt.Println(xi)
	fmt.Println("---------------------")

	// Slicing a slice
	xi2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("xi2 - %#v\n", xi2)
	fmt.Println("-------------")

	// [inclusive:exclusive]
	fmt.Printf("xi2 - %#v\n", xi2[0:4])
	fmt.Println("-------------")

	// [:exclusive]
	fmt.Printf("xi2 - %#v\n", xi2[:7])
	fmt.Println("-------------")

	// [inclusive:]
	fmt.Printf("xi2 - %#v\n", xi2[4:])
	fmt.Println("-------------")

	// [:]
	fmt.Printf("xi2 - %#v\n", xi2[:])
	fmt.Println("-------------")

	// deleting from a slice
	xi2 = append(xi2[:4], xi2[5:]...)
	fmt.Printf("xi2 - %#v\n", xi2)
	fmt.Println("-------------")

	// make
	xi3 := make([]int, 0, 10)
	fmt.Println(xi3)
	fmt.Println(len(xi3))
	fmt.Println(cap(xi3))
	xi3 = append(xi3, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi3)
	fmt.Println("------------")
	fmt.Println(len(xi3))
	fmt.Println(cap(xi3))
	fmt.Println("------------")
	xi3 = append(xi3, 10, 11, 12, 13)
	fmt.Println(xi3)
	fmt.Println(len(xi3))
	fmt.Println(cap(xi3))

	// multidimensional slice
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guinness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xxs := [][]string{jb, jm}
	fmt.Println(xxs)

	// slice internal & underlying array 01
	a2 := []int{0, 1, 2, 3, 4, 5}
	b2 := a2

	fmt.Println("a2", a2)
	fmt.Println("b2", b2)
	fmt.Println("-------------")

	a2[0] = 7

	// have same value
	fmt.Println("a2", a2)
	fmt.Println("b2", b2)
	fmt.Println("-------------")

	// slice internal & underlying array 02
	a3 := []int{0, 1, 2, 3, 4, 5}
	b3 := make([]int, 6)
	copy(b3, a3) // copy value on a3 -> b3

	fmt.Println("a3", a3)
	fmt.Println("b3", b3)
	fmt.Println("-------------")

	a3[0] = 7

	// have not same value
	fmt.Println("a3", a3)
	fmt.Println("b3", b3)
	fmt.Println("-------------")

	n := []float64{3, 1, 4, 2}

	fmt.Println(medianOne(n))
	fmt.Println("after medianOne", n) // n slice was sorted

	y := []float64{3, 1, 4, 2}
	fmt.Println(medianTwo(y))
	fmt.Println("after medianTwo", y)
}

// Uses the same underlying array
// everything is pass by value in go
// the value is being passed into this func
// and then assigned to x
func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
		// https://go.dev/play/p/2r5WrelUEh7
	}
	return (n[i-1] + n[i]) / 2
}
