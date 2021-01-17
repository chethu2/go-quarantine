package main

import "fmt"

func main() {
	twoDArray := [3][5]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {1, 2, 3, 4, 5}}
	for row := 0; row < 3; row++ {
		for col := 0; col < 5; col++ {
			fmt.Printf("Array value at ar[%v][%v] is %v\n", row, col, twoDArray[row][col])
		}
	}

}
