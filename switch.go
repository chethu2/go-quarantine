package main

import (
	"fmt"
)

func main() {
	fmt.Println("ASCII check")
	code := 165
	switch code {
	case 65:
		fmt.Println("The Letter is A")
	case 66:
		fmt.Println("The Letter is B")
	case 67:
		fmt.Println("The Letter is C")
	case 68:
		fmt.Println("The Letter is D")
	default:
		fmt.Printf("Invalid code:%d, code must be between 65 to 68\n", code)
	}
}
