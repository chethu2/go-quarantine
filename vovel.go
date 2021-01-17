package main

import (
	"fmt"
)

func main() {
	var alphabet string

	// reading the input from STDIN
	fmt.Println("Input an alphabet")
	fmt.Scanf("%v", &alphabet)

	if alphabet == "a" || alphabet == "e" || alphabet == "i" || alphabet == "o" || alphabet == "u" {
		fmt.Printf("The Alphabet %s is a vovel\n", alphabet)
		return
	}
	fmt.Printf("The Alphabet %s is not a vovel\n", alphabet)
}
