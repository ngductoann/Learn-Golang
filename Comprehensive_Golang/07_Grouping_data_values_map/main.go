package main

import "fmt"

func main() {
	// create a map and add an element to map
	am := map[string]int{
		"Todd":   42,
		"Henry":  16,
		"Padget": 14,
	}

	// accessing element un a map
	fmt.Println("The age of Henry was", am["Henry"])

	fmt.Println(am)
	fmt.Printf("%#v\n", am)

	// create a map
	an := make(map[string]int)

	// add element to map
	an["Lucas"] = 28
	an["Steph"] = 37
	an["George"] = 78

	fmt.Println(an)

	// print out the entire map
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))

	// for range over a map
	for k, v := range an {
		fmt.Println(k, v) // -> Lucas 28
	}

	for _, v := range an {
		fmt.Println(v) // -> 28
	}

	for k := range an {
		fmt.Println(k) // -> Lucas
	}

	// delete element
	delete(an, "George")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage")      // won't panic
	fmt.Println(an["Georgey"]) // won't panice
	fmt.Println("------------------------")

	// comma ok idiom
	// v, ok := an["Lucas"]
	// if ok {
	// 	fmt.Println("the value prints", v)
	// } else {
	// 	fmt.Println("Key didn't exist")
	// }

	if v, ok := an["Lucas"]; !ok {
		fmt.Println("Key didn't exist")
	} else {
		fmt.Println("the value prints", v)
	}
}
