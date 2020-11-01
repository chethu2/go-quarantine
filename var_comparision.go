package main

import (
	"fmt"
)

func main() {
	var a, b = 5, 10

	if a < b {
		fmt.Printf("%v is less than %v\n", a, b)
	}

	if b > a {
		fmt.Printf("%v is greater than %v\n", b, a)
	}

	if a != b {
		fmt.Printf("%v is not equal to %v\n", a, b)
	}

	if a <= b {
		fmt.Printf("%v is less than or equal to  %v\n", a, b)
	}

	if b >= a {
		fmt.Printf("%v is greater than or equal to  %v\n", b, a)
	}

}
