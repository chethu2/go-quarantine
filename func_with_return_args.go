package main

import "fmt"

func main() {
	n := 20
	isEven := checkEven(n)
	if isEven {
		fmt.Printf("%d is a even number\n.", n)
		return
	}
	fmt.Printf("%n is not a even number\n.", n)
}

func checkEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
