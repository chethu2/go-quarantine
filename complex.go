package main

import (
	"fmt"
)

func main() {

	var x = 5 + 7i     // without specifying type(complex128)
	var y, z complex64 // with type complex64
	y = 5 + 7i

	z = complex(4, 3) // another way to initialise the complex number
	// 4 being real number and 3 being imaginary number.
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
