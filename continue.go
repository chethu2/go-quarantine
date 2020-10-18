// this sample code helpsus understand the working of continue keyword
// this code snippet prints the numbers between 1-10 ignoring multiples of 5
package main

import "fmt"

func main() {
	n := 1
	for n <= 10 {
		if n%5 == 0 {
			n++
			continue // contro is passed to the next iteration.
		}
		fmt.Printf("Value is: %d\n", n)
		n++
	}
	fmt.Println("Fin")
}
