package main

import (
	"fmt"
)

func main() {

	//declaration & initialization

	var defaultTemp = -5
	var typeTemp int32 = -5
	var UnsignedTypeTemp uint32 = 5 // assigning -5, Go will throw the error

	var defaultTemperature = 30.5
	var typeTemperature float32 = 30.5

	// int operation
	fmt.Printf("Int type: %T Value: %v\n", defaultTemp, defaultTemp)
	fmt.Printf("Int type: %T Value: %v\n", typeTemp, typeTemp)
	fmt.Printf("Uint type: %T Value: %v\n", UnsignedTypeTemp, UnsignedTypeTemp)

	// float operation
	fmt.Printf("Float type: %T Value: %v\n", defaultTemperature, defaultTemperature)
	fmt.Printf("Float type: %T Value: %v\n", typeTemperature, typeTemperature)
}
