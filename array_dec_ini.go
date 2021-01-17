package main

import "fmt"

func main() {
	numbers := [10]int{10, 20, 30, 40, 50}
	fmt.Println("data", numbers)
	fmt.Println("len", len(numbers))
	fmt.Println("cap", cap(numbers))
}
