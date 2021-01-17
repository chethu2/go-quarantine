package main

import "fmt"

func main() {
	var a *int
	var b int
	fmt.Printf("Value of the int pointer before initializtion: %v\n", a)
	b = 20
	a = &b
	fmt.Printf("Value of the int pointer after initializtion: %v\n", a)
}
