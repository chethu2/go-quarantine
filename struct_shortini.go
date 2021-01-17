package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Phone string
}

func main() {
	// initialising the struct
	user := Person{"Chethan", 25, "0000000000"}
	fmt.Printf("Name: %v \n", user.Name)
	fmt.Printf("Age: %v \n", user.Age)
	fmt.Printf("Phone: %v \n", user.Phone)
}
