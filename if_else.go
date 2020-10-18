package main

import "fmt"

func main() {

	var a int
	fmt.Println("Input a number and message")
	fmt.Scanf("%v", &a)
	if a%2 == 0 {
		fmt.Printf("The number:%d is divisible by 2\n", a)
	} else if a%3 == 0 {
		fmt.Printf("The number:%d is divisible by 2\n", a)
	} else {
		fmt.Printf("The number:%d is not divisble by 2 or 3\n", a)
	}

}
