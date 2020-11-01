package main

import "fmt"

func main() {
	var a, b uint = 4, 8
	// 4 = 0000 0100
	// 8 = 0000 1000

	res := a & b
	// 0 = 0000 0000
	fmt.Printf("Binary AND operation of %d and %d is %d\n", a, b, res)

	res = a | b
	// 12 = 000 1100
	fmt.Printf("Binary OR operation of %d and %d is %d\n", a, b, res)

	res = a ^ b
	// 12 = 000 1100
	fmt.Printf("Binary XOR operation of %d and %d is %d\n", a, b, res)

	res = a << 1
	// 8 = 0000 1000
	fmt.Printf("Binary left hand shift by 1 bit of  %d is %d\n", a, res)

	res = a >> 1
	// 2 = 0000 0010
	fmt.Printf("Binary left hand shift by 1 bit of  %d is %d\n", a, res)
}
