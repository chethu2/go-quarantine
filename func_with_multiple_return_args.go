package main

import "fmt"

func main() {
	n := 20
	isEven, message := checkEven(n)
	if isEven {
		fmt.Printf("%d is a even number, Status message: %s\n.", n, message)
		return
	}
	fmt.Printf("%n is not a even number, Status message: %s\n.", n, message)
}

func checkEven(n int) (bool, string) {
	if n < 0 {
		return false, "Failure: negative number"
	}
	if n%2 == 0 {
		return true, "sucess"
	}
	return false, "success"
}
