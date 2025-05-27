package main

import "fmt"

type ByteSize int

const (
	_           = iota // ignore the first value by assigning to blank indentifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	// NOTE: Meansuring bits with bitwise operations

	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t\t %b\n", TB, TB)
	fmt.Printf("%d \t\t\t %b\n", PB, PB)
	fmt.Printf("%d \t\t\t %b\n", EB, EB)
}
