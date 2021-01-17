package main

import "fmt"

func main() {
	numbers := [10]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Value at index  %d is %d\n", index, value)
	}
}
