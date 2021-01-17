package main

import "fmt"

func main() {
	b := 20
	a := &b
	fmt.Printf("Value of the int pointer after initializtion: %v\n", a)
}
