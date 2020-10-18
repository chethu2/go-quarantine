package main

import (
	"fmt"
)

func main() {
	var runeType rune
	runeType = 'a'
	var byteType byte
	byteType = 'a'
	fmt.Printf("Type: %T Value: %v\n", runeType, runeType)
	fmt.Printf("Type: %T Value: %v\n", byteType, byteType)
}
