package main

import "fmt"

func main() {
	mySlice := make([]int, 5)
	fmt.Printf("mySlice :%v\n", mySlice)
	fmt.Printf("Capacity of mySlice :%v\n", cap(mySlice))
	fmt.Printf("length of mySlice :%v\n", len(mySlice))
}
