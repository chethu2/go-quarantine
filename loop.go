package main

import (
	"fmt"
)

func main() {
	// declaring the varibles
	var number, sum int

	// reading the inputs
	fmt.Println("Enter Number N")
	fmt.Scanf("%d", &number)
	for i := 1; i < number; i++ {
		sum += i
	}
	fmt.Printf("Sum of Natural Number below %d is: %d\n", number, sum)
}
