package main

import (
	"fmt"
)

func main() {
	var a, b = 5, 10

	res := a + b
	fmt.Printf("Sum of %v and %v is %v\n", a, b, res)
	res = b - a
	fmt.Printf("Difference of %v and %v is %v\n", a, b, res)
	res = b / a
	fmt.Printf("Division of %v and %v is %v\n", a, b, res)
	res = a * b
	fmt.Printf("Product of %v and %v is %v\n", a, b, res)
	res = b % a
	fmt.Printf("Reminder of %v and %v is %v\n", a, b, res)
}
