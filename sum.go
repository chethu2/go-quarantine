package main

import "fmt"

func main() {

	var number, sum int
	// reading the input from STDIN
	fmt.Scanf("%v", &number)
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Print(sum)
}
