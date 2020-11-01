package main

import "fmt"

func main() {
	var isTrue, isFalse = true, false
	res := isTrue || isFalse
	fmt.Printf("Logical OR of %v and %v is %v\n", isTrue, isFalse, res)

	res = isTrue && isFalse
	fmt.Printf("Logical AND of %v and %v is %v\n", isTrue, isFalse, res)

	res = !isTrue
	fmt.Printf("Logical Negation of %v is %v\n", isTrue, res)
}
