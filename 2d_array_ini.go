package main

import "fmt"

func main() {
	twoDArray := [3][5]int{}
	twoDArray = [3][5]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {1, 2, 3, 4, 5}}
	fmt.Println("two D array", twoDArray)

}
