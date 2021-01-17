package main

import "fmt"

func main() {
	var mySlice []int
	for i := 0; i < 5; i++ {
		mySlice = append(mySlice, i)
		fmt.Printf("Slice length %v\n", len(mySlice))
		fmt.Printf("Slice capacity %v\n", cap(mySlice))
	}
	fmt.Printf("Slice %v\n", mySlice)
	fmt.Printf("Slice length %v\n", len(mySlice))
	fmt.Printf("Slice capacity %v\n", cap(mySlice))
}
