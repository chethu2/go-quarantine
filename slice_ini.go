package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("My slice:%v\n", mySlice)
	fmt.Printf("Slice Length :%v\n", len(mySlice))
	fmt.Printf("Slice Capacity :%v\n", cap(mySlice))
}
