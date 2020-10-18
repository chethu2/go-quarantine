package main

import "fmt"

func main() {

	// declare variables
	var size, sum int

	// read the input
	fmt.Scanf("%v", size)

	elements := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%v", &elements[i])
	}

	// iterate over elements and find the answer
	for i := 0; i < len(elements)-1; i++ {
		sum = sum + elements[i]*elements[i+1]
	}

	// print output
	fmt.Print(sum)

}
