package main

import "fmt"

func main() {
	var numbers [10]int
	fmt.Println("data", numbers)
	fmt.Println("len", len(numbers))
	fmt.Println("cap", cap(numbers))
}
