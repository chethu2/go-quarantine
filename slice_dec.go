package main

import "fmt"

func main() {
	var mySlice []int
	fmt.Printf("My slice:%v\n", mySlice)
	fmt.Printf("Slice Length :%v\n", len(mySlice))
	fmt.Printf("Slice Capacity :%v\n", cap(mySlice))
}
