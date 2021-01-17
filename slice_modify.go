package main

import "fmt"

func main() {
	mySlice := []string{"Hello", "World", "Good", "Morning"}
	fmt.Printf("My slice at index 3 before update: %v\n", mySlice[3])
	mySlice[3] = "Evening"
	fmt.Printf("My slice at index 3 after update: %v\n", mySlice[3])
}
