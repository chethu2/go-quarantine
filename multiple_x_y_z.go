package main

import "fmt"

func main() {
	var x, y, z int

	fmt.Print("Enter the value of X")
	fmt.Scanf("%s", &x)
	fmt.Print("Enter the value of Y")
	fmt.Scanf("%s", &y)
	fmt.Print("Enter the Range Z")
	fmt.Scanf("%s", &z)

	for i := 1; i < z; i++ {
		if i%x == 0 && i%y == 0 {
			fmt.Println(i)
		}
	}

}
