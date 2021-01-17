package main

import "fmt"

func main() {
	var a *int
	b := 20
	a = &b
	fmt.Printf("the Pointer value(*a) is %v\n", *a)
	*a = 40
	fmt.Printf("the Pointer value(*a) is %v\n", *a)
}
