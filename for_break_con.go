package main

import "fmt"

func main() {
	for i := 1; i < 20; i++ {
		if i >= 20 {
			break // control go out of the loop
		}
		if i%3 != 0 {
			continue // control will be passed back
		}
		fmt.Printf("%v is divisible by 3\n", i)

	}
}
