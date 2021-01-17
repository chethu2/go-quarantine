package main

import (
	"fmt"
)

func main() {
	// declaring the varibles
	var name string
	var age int

	// reading the inputs
	fmt.Println("Enter name!")
	fmt.Scanf("%s", &name)
	fmt.Printf("Welcome %s!!\n", name)
	fmt.Println("Enter Age!")
	fmt.Scanf("%d", &age)

	// let's compare

	switch {
	case age <= 25:
		fmt.Printf("%s is you are young!\n", name)
	case age > 25 && age < 60:
		fmt.Printf("%s is you are old!\n", name)
	case age > 60:
		fmt.Printf("%s is you are very old!\n", name)
	default:
		fmt.Printf("%s something is wrong!\n", name)

	}
}
