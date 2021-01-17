package main

import "fmt"

func main() {
	var numbers [10]int
	numbers[0] = 25
	numbers[1] = 45
	numbers[2] = 50
	numbers[3] = 60
	numbers[4] = 70
	fmt.Println("data", numbers)
	fmt.Println("len", len(numbers))
	fmt.Println("cap", cap(numbers))
}
