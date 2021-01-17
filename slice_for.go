package main

import "fmt"

func main() {
	myslice := []string{"Good", "Morning", "Good", "Afternoon"}
	for i := 0; i < len(myslice); i++ {
		fmt.Printf("Value at index:%v is %v\n", i, myslice[i])
	}

}
