package main

import "fmt"

func main() {
	n, sum := 0, 0
	for n < 5 {
		sum = n + sum
		n++
	}
	fmt.Printf("Sum is %v\n", sum)
}
