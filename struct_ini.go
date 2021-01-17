package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Phone string
}

func main() {
	var user Person
	// by default initialised to zero values
	fmt.Println("Before Initializtion")
	fmt.Printf("Name: %v \n", user.Name)
	fmt.Printf("Age: %v \n", user.Age)
	fmt.Printf("Phone: %v \n", user.Phone)

	// initialising the struct
	user = Person{"Chethan", 25, "0000000000"}
	fmt.Println("After Initializtion")
	fmt.Printf("Name: %v \n", user.Name)
	fmt.Printf("Age: %v \n", user.Age)
	fmt.Printf("Phone: %v \n", user.Phone)
}
