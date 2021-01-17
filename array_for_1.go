package main

import "fmt"

func main() {
	numbers := [10]int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Value at index  %d is %d\n", i, numbers[i])
	}
}
