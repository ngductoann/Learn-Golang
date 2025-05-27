package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// Basic Type
	/*
	   NOTE:
	   Go basic types are:
	   	-  bool
	   	- int, int8, int16, int32, int64, unit, uint8, uint16, uint32, uint64
	   	- byte // alias for uint8
	   	- rune // alias for int32. Represents a Unicode code point
	   	- float32, float64
	   	- complex64, complex128
	   The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements
	   The int, unit, and uintptr types are usually 32 bits wide in 32-bit systems and 64-bit wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
	*/

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Zero values (in 01_Getting Document)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversions
	/*
		NOTE:
		The expression T(v) converts the value v to the type T.
		Some numeric conversions:
			var i int = 42
			var f float64 = float64(i)
			var u uint = uint(f)
		Or, put more simply:
			i := 42
			f := float64(i)
			u := uint(f)
		Unlike in C, in Go assignment between items of different type requires and explicit conversion. Try remove the float64 or uint conversions in the example and see what happens
	*/

	var x, y int = 3, 4
	var f2 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f2)
	fmt.Println(x, y, z)
}
