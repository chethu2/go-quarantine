package main

import "fmt"

func main() {
	n, sum := 0, 0
	for {
		if n >= 5 {
			fmt.Printf("Sum is %v\n", sum)
			return
		}
		sum = n + sum
		n++
	}
}
