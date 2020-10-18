package main

import (
	"fmt"
)

// this sample program takes two input from STDIN
// and prints on the STOUT
func main() {
	// declaring the varibales
	var number int
	var message string

	// reading the input from STDIN
	fmt.Println("Input a numbe and message")
	fmt.Scanf("%v", &number)
	fmt.Scanf("%s", &message)

	// output
	fmt.Println(number)
	fmt.Println(message)

}
