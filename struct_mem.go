package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

func main() {
	var user Person
	fmt.Printf("Size of Name: %v\n", unsafe.Sizeof(user.Name))
	fmt.Printf("Size of Age: %v\n", unsafe.Sizeof(user.Age))
	fmt.Printf("Size of Phone : %v\n", unsafe.Sizeof(user.Phone))

	fmt.Printf("Size of user: %v\n", unsafe.Sizeof(user))
}
