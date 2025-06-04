package main

import "fmt"

func paradise(loc string) string {
	return fmt.Sprint("My idea of paradise is ", loc)
}

func main() {
	fmt.Println(paradise("Hawaii"))
}
