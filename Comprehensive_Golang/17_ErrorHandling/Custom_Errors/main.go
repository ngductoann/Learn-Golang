package main

import (
	"fmt"
	"log"
)

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("norgate math: square root of negative number")
		// or
		// return 0, ErrNorgateMath // Using custom error
		// or
		nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	// implementation
	return 42, nil // Placeholder return value (42 is not the square root of f, nil is the error)
}
