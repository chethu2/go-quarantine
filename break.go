// the sample helps usnderstand the working of the golang's break keyword.
package main

import "fmt"

func main() {

	var r int = 0
	for r < 10 {
		if r == 5 {
			fmt.Printf("%d is 5\n", r)
			break // control now will be passed out of the loop
		}
		fmt.Printf("%d is not 5\n", r)
		r++
	}
	fmt.Println("Fin")
}
