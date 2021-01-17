package main

import "fmt"

func main() {
	a := 20
	increment(a)
	fmt.Printf("address of a is %v\n", &a)
	fmt.Printf("value of a after increment: %v\n", a)
}

func increment(a int) {
	a = a + 1
	fmt.Printf("address of a is %v\n", &a)
	fmt.Printf("value of a inside function: %v\n", a)
}
