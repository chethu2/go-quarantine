package main

import "fmt"

func main() {
	myslice := []string{"Good", "Morning", "Good", "Afternoon"}
	for i, val := range myslice {
		fmt.Printf("Value at index:%v is %v\n", i, val)
	}

}
